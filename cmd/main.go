package main

import (
	"github.com/sirupsen/logrus"
	todo "todo-app"
	"todo-app/pkg/config"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/repository/postgres"
	"todo-app/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg, err := config.InitConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := postgres.NewPostgresDB(&postgres.DBConfig{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err = srv.Run(cfg.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}

}
