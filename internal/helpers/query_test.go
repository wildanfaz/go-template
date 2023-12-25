package helpers

import (
	"fmt"
	"testing"
)

type user struct {
	Id    int    `column:"where:id"`
	Name  string `column:"name"`
	Email string `column:"where:email"`
}

func TestMySQLUpdateQueryValues(t *testing.T) {
	user := &user{
		Id:    1,
		Name:  "Foo",
		Email: "foo@bar.com",
	}

	query, values, err := MySQLUpdateQueryValues(user,  "users", "column")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%s\n%v\n", query, values)
}
