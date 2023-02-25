package rest

import (
	usecase "GoStorageService/internal/usecase/ads"
	"GoStorageService/utils/validate"

	"clevergo.tech/jsend"
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

	if err, status := validate.CheckQueryMarket(name, category); !status {
		return c.Status(fiber.StatusBadRequest).JSON(jsend.NewError(err.Error(), fiber.StatusBadRequest, nil))
	}

	err := h.service.Add(c.Body(), name, category)
	if err != nil {
		logrus.Errorf("error create new data: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(jsend.NewError(err.Error(), fiber.StatusBadRequest, nil))
	}

	return c.Status(fiber.StatusOK).JSON(Response("success"))
}
