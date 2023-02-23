package rest

import (
	"github.com/gofiber/fiber/v2"
)

func InitApi(app *fiber.App, api Handler) {
	app.Get("/api/v1", Start)
	app.Post("/api/v1/ads/add", api.Add)
}
