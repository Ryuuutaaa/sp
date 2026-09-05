package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Validasi header Authorization atau session token dari Nuxt/Better-Auth
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized access",
			})
		}
		return c.Next()
	}
}
