package rest

import "github.com/gofiber/fiber/v2"

func Response(status string) fiber.Map {
	return fiber.Map{
		"status": status,
	}
}

func Start(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Response("success"))
}
