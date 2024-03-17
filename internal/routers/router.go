package routers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/go-template/configs"
	"github.com/wildanfaz/go-template/internal/constants"
	"github.com/wildanfaz/go-template/internal/pkg"
	"github.com/wildanfaz/go-template/internal/repositories"
	books_router "github.com/wildanfaz/go-template/internal/routers/books-router"
	"github.com/wildanfaz/go-template/internal/services/books"
	"github.com/wildanfaz/go-template/internal/services/health"
)

func InitEchoRouter() {
	// configs
	config := configs.InitConfig()
	dbMySql := configs.InitMySql(config.DatabaseDSN)

	// pkg
	log := pkg.InitLogger()

	// repositories
	booksRepo := repositories.NewBooksRepository(dbMySql)

	// services
	_ = books.New(booksRepo, log, config)

	e := echo.New()

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/health", health.HealthCheck)

	books_router.BooksRouter(apiV1)

	e.Logger.Fatal(e.Start(config.AppPort))
}
