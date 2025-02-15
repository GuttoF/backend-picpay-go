package handlers

import (
	"backend-picpay/internal/models"
	"backend-picpay/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Transfer handles the transfer of funds between two users.
// @Summary Transfer funds between users
// @Description Transfers a specified amount from the payer to the payee
// @Tags transfer
// @Accept json
// @Produce json
// @Param transfer body models.Transaction true "Transaction details"
// @Success 201 {object} models.Transaction
// @Failure 400 {object} models.ErrorResponse "Invalid input or transfer error"
// @Router /transfer [post]
func TransferHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var transaction models.Transaction

		if err := c.BodyParser(&transaction); err != nil {
			return c.Status(400).JSON(models.ErrorResponse{Error: "Cannot parse JSON"})
		}

		if transaction.PayerID == [16]byte{} || transaction.PayeeID == [16]byte{} || transaction.Value <= 0 {
			return c.Status(400).JSON(models.ErrorResponse{Error: "PayerID, PayeeID, and valid Amount are required"})
		}

		if err := services.Transfer(db, &transaction); err != nil {
			return c.Status(400).JSON(models.ErrorResponse{Error: err.Error()})
		}

		return c.Status(201).JSON(transaction)
	}
}