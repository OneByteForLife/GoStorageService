package main

import (
	"GoStorageService/internal/app"
	"GoStorageService/internal/database"
	"GoStorageService/pkg"

	"github.com/sirupsen/logrus"
)

func init() {
	pkg.ConfigLog()
	// Проверка подключения к базе
}

func main() {
	if db, err := database.ConnectDataBase(); db == nil || err != nil {
		logrus.Warnf("No connection to the base was established -%s", err)
	} else {
		db.Close()
	}

	app.Run()
}
