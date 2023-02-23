package rest

import (
	usecase "GoStorageService/internal/usecase/ads"
	"GoStorageService/utils/validate"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type MarketHandler struct {
	service usecase.Service
}

func NewHandler(s usecase.Service) Handler {
	return &MarketHandler{s}
}

func (h *MarketHandler) Add(c *fiber.Ctx) error {
	name, category := c.Query("name"), c.Query("category")

	if status := validate.CheckUrlQuery(name, category); !status {
		return c.Status(fiber.StatusBadRequest).JSON(RespStatus("1.0", fiber.StatusBadRequest, "invalid url query", nil))
	}

	err := h.service.Add(c.Body(), name, category)
	if err != nil {
		logrus.Errorf("error create new data: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(RespStatus("1.0", fiber.StatusBadRequest, err.Error(), nil))
	}

	return nil
}
