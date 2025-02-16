package main

import (
	blogging "blogging_app"
	"blogging_app/pkg/handler"
	"blogging_app/pkg/repository"
	"blogging_app/pkg/service"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("error occured reading environment variables: %s", err.Error())
	}

	if err := InitConfigs(); err != nil {
		logrus.Errorf("error occured setting configuration variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Errorf("error occured connecting to db: %s", err.Error())
	}

	defer db.Close()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(blogging.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Errorf("error occured server didn't start working: %s", err.Error())
		}
	}()

	logrus.Print("BloggingApp has started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logrus.Print("BloggingApp shutting down")
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func InitConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
