package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"university/internal/config"
	"university/internal/database"
	"university/internal/handlers/v1"
	"university/internal/repository"
	"university/internal/service"
)

func main() {
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s\n", err.Error())
	}

	cfg := config.New()

	db, err := database.NewPostgresDB(cfg.Post)
	if err != nil {
		logrus.Fatalf("db opening error: %s\n", err.Error())
	}

	e := echo.New()
	repo := repository.New(db)
	serv := service.New(repo)

	v1.New(serv).Register(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
