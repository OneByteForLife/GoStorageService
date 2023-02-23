package rest

import "github.com/gofiber/fiber/v2"

type (
	Handler interface {
		Add(c *fiber.Ctx) error
	}
)
