package handlers

import (
	"backend-picpay/internal/models"
	"backend-picpay/internal/services"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Cannot parse JSON"})
	}

	if user.Email == "" || user.Password == "" {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Email and password are required"})
	}

	if err := services.CreateUser(&user); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: err.Error()})
	}

	return c.Status(201).JSON(user)
}