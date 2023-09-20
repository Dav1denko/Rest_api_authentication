package main

import (
	restapiauthentication "Rest_api_authentication"
	"Rest_api_authentication/pkg/handler"
	"Rest_api_authentication/pkg/repository"
	"Rest_api_authentication/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restapiauthentication.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
