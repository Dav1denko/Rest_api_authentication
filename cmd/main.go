package main

import (
	restapiauthentication "Rest_api_authentication"
	"Rest_api_authentication/pkg/handler"
	"Rest_api_authentication/pkg/repository"
	"Rest_api_authentication/pkg/service"
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init configs %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error init env var %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	client, err := repository.NewMongoDB(ctx, repository.Config{
		Username: os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		Host:     viper.GetString("client.host"),
		Port:     viper.GetString("client.port"),
	})
	if err != nil {
		log.Fatalf("failed to initialize mongo %s", err.Error())
	}

	repos := repository.NewRepository(client)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restapiauthentication.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
