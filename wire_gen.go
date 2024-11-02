// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func InitializeApplication() (*fiber.App, error) {
	db, err := InitializeDB()
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(db)
	premiumPackageRepository := repository.NewPremiumPackageRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, premiumPackageRepository)
	userHandler := http.NewUserHandler(userUsecase)
	swipeRepository := repository.NewSwipeRepository(db)
	swipeUsecase := usecase.NewSwipeUsecase(userRepository, swipeRepository)
	swipeHandler := http.NewSwipeHandler(swipeUsecase)
	app := NewApp(userHandler, swipeHandler)
	return app, nil
}

// wire.go:

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

var RepositorySet = wire.NewSet(repository.NewUserRepository, repository.NewPremiumPackageRepository, repository.NewSwipeRepository)

var UsecaseSet = wire.NewSet(usecase.NewUserUsecase, usecase.NewSwipeUsecase)

var HandlerSet = wire.NewSet(http.NewUserHandler, http.NewSwipeHandler)

func NewApp(userHandler *http.UserHandler, swipeHandler *http.SwipeHandler) *fiber.App {
	app := fiber.New()
	SetupRoutes(app, userHandler, swipeHandler)
	return app
}