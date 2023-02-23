package rest

import "github.com/gofiber/fiber/v2"

func RespStatus(apiVersion string, statusCode int, description string, content interface{}) fiber.Map {
	return fiber.Map{
		"api_version": apiVersion,
		"status_code": statusCode,
		"description": description,
		"content":     content,
	}
}

func Start(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "WORK", nil))
}
