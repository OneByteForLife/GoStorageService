package middleware

import (
	"GoStorageService/internal/controllers/rest"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func CheckJwtToken() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey:   []byte(os.Getenv("JWT_SEED")),
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(rest.RespStatus("1.0", fiber.StatusUnauthorized, "Incorrect token auth", nil))
}
