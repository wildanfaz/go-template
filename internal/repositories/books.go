package repositories

import "database/sql"

type ImplementBooks struct {
	dbMySql *sql.DB
}

type Books interface {
}

func NewBooksRepository(dbMySql *sql.DB) Books {
	return &ImplementBooks{
		dbMySql: dbMySql,
	}
}
