package main

import (
	"log"

	"github.com/Xlussov/rest-api-v2"
	"github.com/Xlussov/rest-api-v2/pkg/handler"
	"github.com/Xlussov/rest-api-v2/pkg/repository"
	"github.com/Xlussov/rest-api-v2/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	
	srv := new(restApi.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
 
