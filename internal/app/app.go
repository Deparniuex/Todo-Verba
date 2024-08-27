package app

import (
	"Todo-Verba/internal/handler"
	"Todo-Verba/internal/httpserver"
	"Todo-Verba/internal/repository/pgrepo"
	"Todo-Verba/internal/service"
	"Todo-Verba/internal/storage/postgres"
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	postgresMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
)

func Run() {
	logrus.Info("starting connection to postgres")
	db, err := postgres.ConnectDB(&postgres.Config{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetInt("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
	})
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("connection successful")

	err = upMigrations(db)
	if err != nil {
		logrus.Fatal(err)
	}

	repository := pgrepo.New(db)
	services := service.New(repository)
	handler := handler.New(services)
	server := httpserver.NewServer(handler.InitRouter(), &httpserver.ServerConfig{
		Host: viper.GetString("SERVER_HOST"),
		Port: viper.GetString("SERVER_PORT"),
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

func upMigrations(db *sql.DB) error {
	driver, err := postgresMigrate.WithInstance(db, &postgresMigrate.Config{})
	if err != nil {
		return err
	}

	//running from cmd
	migrator, err := migrate.NewWithDatabaseInstance("file://migrations/", "postgres", driver)
	if err != nil {
		return err
	}

	err = migrator.Up()
	logrus.Infof("migrations status: %s", err)
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
