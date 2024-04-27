package main

import (
	"log"

	"github.com/Xlussov/rest-api-v2"
	"github.com/Xlussov/rest-api-v2/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	
	srv := new(restApi.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}

