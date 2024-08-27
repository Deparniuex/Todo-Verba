package app

import (
	"Todo-Verba/internal/handler"
	"Todo-Verba/internal/httpserver"
	"Todo-Verba/internal/repository/pgrepo"
	"Todo-Verba/internal/service"
	"Todo-Verba/internal/storage/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
)

func Run() {
	logrus.Info("starting connection to postgres")
	db, err := postgres.ConnectDB(&postgres.Config{
		Host:     "",
		Port:     0,
		User:     "",
		Password: "",
		DBName:   "",
	})
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("connection successful")
	repository := pgrepo.New(db)
	services := service.New(repository)
	handler := handler.New(services)
	server := httpserver.NewServer(handler.InitRouter(), &httpserver.ServerConfig{
		Host: "",
		Port: "",
	})
	server.Start()
	logrus.Info("server started")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		logrus.Infof("signal received: %s", s.String())
	case err = <-server.Notify():
		logrus.Infof("server notify: %s", err.Error())
	}
}

func SetupConfig(path string) error {
	viper.SetConfigFile(path)
	return viper.ReadInConfig()
}

func SetupLogger() error {
	level, err := logrus.ParseLevel(viper.GetString("LOG_LEVEL"))
	if err != nil {
		return err
	}
	logrus.SetLevel(level)
	return nil
}
