package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
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

	r := gin.Default()
	repo := repository.New(db)
	serv := service.New(repo)

	v1.New(serv).Register(r)

	// Start server
	log.Fatal(r.Run(":8080"))
}
