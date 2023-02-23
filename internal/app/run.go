package app

import (
	"GoStorageService/config"
	rest "GoStorageService/internal/controllers/rest"
	"GoStorageService/internal/middleware"
	usecase "GoStorageService/internal/usecase/ads"
	"GoStorageService/utils/database"
	"time"

	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Run(conf *config.ConfigDb) error {
	// Конифгурация сервера
	app := fiber.New(fiber.Config{
		JSONDecoder: gojson.Unmarshal,
		JSONEncoder: gojson.Marshal,
		// Prefork:      true, (Не работае в Alpine Linux который я используя из за его легковестности) при сборке в Docker
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	})
	app.Use(logger.New(), middleware.CheckJwtToken())

	db, err := database.ConnectDataBase()
	if err != nil {
		return err
	}

	api := rest.NewHandler(usecase.NewService(usecase.NewStorage(db)))
	rest.InitApi(app, api)

	if err := app.Listen(":8080"); err != nil {
		return err
	}
	return nil
}
