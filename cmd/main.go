package main

import (
	todo "awesomeProject"
	"awesomeProject/pkg/handler"
	"awesomeProject/pkg/repo"
	"awesomeProject/pkg/service"
	//"github.com/golang-migrate/migrate/v4"
	//"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init configs: %s", err.Error())
	}
	///Внедрение зависимостей
	repos := repo.NewRepo()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
