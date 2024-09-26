package main

import (
	"github.com/biyoba1/redisProject"
	"github.com/biyoba1/redisProject/initializer"
	"github.com/biyoba1/redisProject/internal/handler"
	"github.com/biyoba1/redisProject/internal/repository"
	services2 "github.com/biyoba1/redisProject/internal/services"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
	initializer.SyncDatabase()
	initializer.ConnectToRedis()
}

func main() {

	repos := repository.NewRepository(initializer.DB)
	services := services2.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(redisProject.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Failed to start a server: %s", err.Error())
	}
}
