package app

import (
	"GoStorageService/internal/middleware"
	"GoStorageService/internal/routes"
	"time"

	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func SetUpRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1/")
	v1.Get("", routes.Home)
	v1.Post("/adding", routes.AddData, middleware.CheckContentType())
	//v1.Get("/retrieved", routes.AddData)
}

func Run() {
	// Конифгурация сервера
	app := fiber.New(fiber.Config{
		JSONDecoder: gojson.Unmarshal,
		JSONEncoder: gojson.Marshal,
		// Prefork:      true, (Не работае в Alpine Linux который я используя из за его легковестности)
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	})

	app.Use(logger.New(), middleware.CheckJwtToken())

	SetUpRoutes(app)

	if err := app.Listen(":8080"); err != nil {
		logrus.Fatalf("Err up server - %s", err)
	}

	logrus.Info("Service is up!")
}
