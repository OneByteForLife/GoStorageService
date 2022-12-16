package routes

import (
	"GoStorageService/internal/models"

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

var data []models.Data

func Home(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	// Тут будет документация

	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "Hello, it is starting page", nil))
}

func AddData(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	if status, massage := models.AddingData(c.Body(), c.Query("market"), c.Query("category")); !status {
		return c.Status(fiber.StatusBadRequest).JSON(RespStatus("1.0", fiber.StatusBadRequest, massage, nil))
	}

	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "Success Adding", nil))
}

func FindData(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	// Логика по получению данных из базы

	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "Success", nil))
}
