package routers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/go-template/configs"
	"github.com/wildanfaz/go-template/internal/constants"
	"github.com/wildanfaz/go-template/internal/pkg"
	"github.com/wildanfaz/go-template/internal/repositories"
	"github.com/wildanfaz/go-template/internal/services/books"
	"github.com/wildanfaz/go-template/internal/services/health"
)

func InitEchoRouter() {
	fmt.Println(constants.Blue, "---Init Echo Router---")

	// configs
	config := configs.InitConfig()
	dbMySql := configs.InitMySql()

	// pkg
	log := pkg.InitLogger()

	// repositories
	booksRepo := repositories.NewBooksRepository(dbMySql)

	// services
	_ = books.NewService(booksRepo, log, config)

	e := echo.New()

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/health", health.HealthCheck)

	e.Logger.Fatal(e.Start(config.EchoPort))
}
