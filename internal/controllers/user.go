package controllers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) getUsers(c *fiber.Ctx) error {

	users, err := s.store.GetAllUsers(c.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(users)
}
