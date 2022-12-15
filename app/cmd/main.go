package main

import (
	"GoStorageService/internal/app"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})
}

func main() {
	app.Run()
}
