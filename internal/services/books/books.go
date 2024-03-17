package books

import (
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/go-template/configs"
	"github.com/wildanfaz/go-template/internal/repositories"
)

type ImplementServices struct {
	books  repositories.Books
	log    *logrus.Logger
	config *configs.Config
}

type Services interface {
}

func New(
	books repositories.Books,
	log *logrus.Logger,
	config *configs.Config,
) Services {
	return &ImplementServices{
		books:  books,
		log:    log,
		config: config,
	}
}
