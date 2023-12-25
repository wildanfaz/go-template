package helpers

import (
	"fmt"
	"reflect"
	"strings"
)

func MySQLUpdateQueryValues(payload any, table, tag string) (string, []any, error) {
	vo := reflect.ValueOf(payload)
	to := reflect.TypeOf(payload)

	if vo.Kind() == reflect.Ptr {
		vo = vo.Elem()
		to = to.Elem()
	}

	if vo.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("Expected %v, got %v", reflect.Struct, vo.Kind())
	}

	query := fmt.Sprintf("UPDATE %s \nSET ", table)

	values := []any{}
	whereKeys := []string{}
	whereValues := []any{}

	for i := 0; i < vo.NumField(); i++ {
		v := vo.Field(i)

		if v.IsZero() {
			continue
		}

		column := to.Field(i).Tag.Get("column")

		if strings.Contains(column, "where") {
			whereValues = append(whereValues, v.Interface())
			whereKeys = append(whereKeys, strings.Split(column, ":")[1])
			continue
		}
		query = fmt.Sprintf(`%s%s = ?, `, query, column)
		values = append(values, v.Interface())
	}
	query = strings.TrimSuffix(query, ", ")

	for i := 0; i < len(whereKeys); i++ {
		if i == 0 {
			query += "\nWHERE "
		}
		query = fmt.Sprintf("%s%s = ? AND ", query, whereKeys[i])
		values = append(values, whereValues[i])
	}
	query = strings.TrimSuffix(query, "AND ")

	return query, values, nil
}
