// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"api-dating/delivery/http"
	"api-dating/infrastructure"
	"api-dating/repository"
	"api-dating/usecase"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func InitializeDB() (*sql.DB, error) {
	db, err := infrastructure.ConnectDB("postgres://postgres:Liandi99@localhost/db_dating?sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := infrastructure.SeedPremiumPackages(db); err != nil {
		return nil, err
	}

	return db, nil
}

var RepositorySet = wire.NewSet(
	repository.NewUserRepository,
	repository.NewPremiumPackageRepository,
	repository.NewSwipeRepository,
)

var UsecaseSet = wire.NewSet(
	usecase.NewUserUsecase,
	usecase.NewSwipeUsecase,
)

var HandlerSet = wire.NewSet(
	http.NewUserHandler,
	http.NewSwipeHandler,
)

func InitializeApplication() (*fiber.App, error) {
	wire.Build(
		InitializeDB,
		RepositorySet,
		UsecaseSet,
		HandlerSet,
		NewApp,
	)
	return &fiber.App{}, nil
}

func NewApp(userHandler *http.UserHandler, swipeHandler *http.SwipeHandler) *fiber.App {
	app := fiber.New()
	SetupRoutes(app, userHandler, swipeHandler)
	return app
}
