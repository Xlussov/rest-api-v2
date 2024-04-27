package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

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

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}


	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})


	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
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