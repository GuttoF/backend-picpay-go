package handlers

import (
	"backend-picpay/internal/models"
	"backend-picpay/internal/services"
	"github.com/gofiber/fiber/v2"
)

func Transfer(c *fiber.Ctx) error {
	var transaction models.Transaction

	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Cannot parse JSON"})
	}

	if transaction.PayerID == [16]byte{} || transaction.PayeeID == [16]byte{} || transaction.Amount <= 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "PayerID, PayeeID, and valid Amount are required"})
	}

	if err := services.Transfer(&transaction); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: err.Error()})
	}

	return c.Status(201).JSON(transaction)
}