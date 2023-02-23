package main

import (
	"GoStorageService/config"
	"GoStorageService/internal/app"
	logs "GoStorageService/utils/log"

	"github.com/sirupsen/logrus"
)

func main() {
	logs.ConfigLog()
	conf := config.ReadConfig()

	if err := app.Run(conf); err != nil {
		logrus.Fatalf("fatal run server: %v", err)
	}
}
