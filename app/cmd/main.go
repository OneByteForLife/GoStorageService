package main

import (
	"GoStorageService/internal/app"
	"GoStorageService/pkg"
)

func init() {
	pkg.ConfigLog()
}

func main() {
	app.Run()
}
