package middleware

import (
	"os"

	"clevergo.tech/jsend"
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
	return c.Status(fiber.StatusUnauthorized).JSON(jsend.NewError("Incorrect token auth", fiber.StatusUnauthorized, nil))
}
