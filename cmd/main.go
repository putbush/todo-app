package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	go func() {
		if err = srv.Run(cfg.Port, handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatal(err)
		}
	}()

	logrus.Println("TodoApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		logrus.Println("timeout of 3 seconds.")
	}
	logrus.Println("Server exiting")
	_ = db.Close()
}
