package main

import (
	"Todo-Verba/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	err := app.SetupConfig("config/.env")
	if err != nil {
		logrus.Fatal(err)
	}
	err = app.SetupLogger()
	if err != nil {
		logrus.Fatal(err)
	}
	app.Run()
}
