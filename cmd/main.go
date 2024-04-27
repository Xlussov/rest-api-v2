package main

import (
	"log"

	"github.com/Xlussov/rest-api-v2"
	"github.com/Xlussov/rest-api-v2/pkg/handler"
	"github.com/Xlussov/rest-api-v2/pkg/repository"
	"github.com/Xlussov/rest-api-v2/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("error initialization configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	
	srv := new(restApi.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
 
func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}