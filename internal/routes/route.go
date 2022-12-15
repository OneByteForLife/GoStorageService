package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RespStatus(apiVersion string, statusCode int, description string, content interface{}) fiber.Map {
	return fiber.Map{
		"api_version": apiVersion,
		"status_code": statusCode,
		"description": description,
		"content":     content,
	}
}

func Home(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	// Тут будет документацияg
	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "Hello, it is starting page", nil))
}
