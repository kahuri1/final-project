package main

import (
	todo "github.com/kahuri1/final-project"
	"github.com/kahuri1/final-project/pkg/handler"
	"github.com/kahuri1/final-project/pkg/repository"
	"github.com/kahuri1/final-project/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	db, err := repository.NewSqlLiteDB(repository.Config{
		DBName: viper.GetString("db.dbname"),
	})
	if err != nil {
		logrus.Errorf("Failed initialization db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.Newhandler(service)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath(".") //исправить на нормальный путь
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
