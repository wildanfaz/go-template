package books

import (
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/go-template/internal/repositories"
)

type ImplementServices struct {
	booksRepo repositories.Books
	log       *logrus.Logger
}

type Services interface {
}

func NewService(
	booksRepo repositories.Books,
	log *logrus.Logger,
) Services {
	return &ImplementServices{
		booksRepo: booksRepo,
		log:       log,
	}
}
