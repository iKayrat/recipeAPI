package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iKairat/RecipeAPI/internal/utils"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := utils.ParseJWT(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthneticated",
		})
	}

	return c.Next()
}
