package repositories

import "database/sql"

type Repository struct {
	dbMySql *sql.DB
}

type BooksRepository interface {
}

func NewBooksRepository(dbMySql *sql.DB) BooksRepository {
	return &Repository{
		dbMySql: dbMySql,
	}
}
