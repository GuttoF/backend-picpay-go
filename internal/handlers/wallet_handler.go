package handlers

import (
    "backend-picpay/internal/models"
    "backend-picpay/internal/services"
    "github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
    "github.com/google/uuid"
)

// CreateWalletHandler cria uma nova carteira para um usuário.
// @Summary Create a new wallet
// @Description Create a new wallet for a user
// @Tags wallets
// @Accept json
// @Produce json
// @Param user_id body string true "User ID (UUID format)"
// @Success 201 {object} models.Wallet
// @Failure 400 {object} models.ErrorResponse
// @Router /wallets [post]
func CreateWalletHandler(c *fiber.Ctx) error {
    var request struct {
        UserID string `json:"user_id"`
    }

    if err := c.BodyParser(&request); err != nil {
        return c.Status(400).JSON(models.ErrorResponse{Error: "Cannot parse JSON"})
    }

    userUUID, err := uuid.Parse(request.UserID)
    if err != nil {
        return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid UUID format"})
    }

    wallet, err := services.CreateWallet(userUUID)
    if err != nil {
        return c.Status(400).JSON(models.ErrorResponse{Error: err.Error()})
    }

    return c.Status(201).JSON(wallet)
}

// GetWalletByUserIDHandler retorna a carteira de um usuário pelo ID do usuário.
// @Summary Get wallet by user ID
// @Description Get a wallet by user ID
// @Tags wallets
// @Accept json
// @Produce json
// @Param user_id path string true "User ID (UUID format)"
// @Success 200 {object} models.Wallet
// @Failure 400 {object} models.ErrorResponse
// @Router /wallets/{user_id} [get]
func GetWalletByUserIDHandler(c *fiber.Ctx) error {
    userID := c.Params("user_id")

    userUUID, err := uuid.Parse(userID)
    if err != nil {
        return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid UUID format"})
    }

    wallet, err := services.GetWalletByUserID(userUUID)
    if err != nil {
        return c.Status(400).JSON(models.ErrorResponse{Error: err.Error()})
    }

    return c.Status(200).JSON(wallet)
}

var validate = validator.New()

// DepositHandler handles the request to add balance to a wallet.
// @Summary Deposit some amount to wallet
// @Description Adds a specified amount of balance to the wallet identified by WalletID.
// @Tags wallets
// @Accept json
// @Produce json
// @Param request body models.BalanceRequest true "Balance Request"
// @Success 200 {object} models.Wallet "Updated wallet information"
// @Failure 400 {object} models.ErrorResponse "Invalid input data or request body"
// @Router /wallets/deposit [post]
func DepositHandler(c *fiber.Ctx) error {
    var request models.BalanceRequest

    if err := c.BodyParser(&request); err != nil {
        return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid JSON request body"})
    }

    if err := validate.Struct(request); err != nil {
        return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid input data"})
    }

    walletUUID, _ := uuid.Parse(request.WalletID)
    wallet, err := services.Deposit(walletUUID, request.Amount)
    if err != nil {
        return c.Status(400).JSON(models.ErrorResponse{Error: err.Error()})
    }

    return c.Status(200).JSON(wallet)
}

// WithdrawHandler handles the withdrawal of balance from a wallet.
// @Summary Withdraw balance from wallet
// @Description Withdraws a specified amount from the wallet identified by WalletID.
// @Tags wallets 
// @Accept json
// @Produce json
// @Param request body models.BalanceRequest true "Balance Request"
// @Success 200 {object} models.Wallet "Wallet with updated balance"
// @Failure 400 {object} models.ErrorResponse "Invalid input data or request body"
// @Router /wallets/withdraw [post]
func WithdrawHandler(c *fiber.Ctx) error {
    var request models.BalanceRequest

    if err := c.BodyParser(&request); err != nil {
        return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid JSON request body"})
    }

    if err := validate.Struct(request); err != nil {
        return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid input data"})
    }

    walletUUID, _ := uuid.Parse(request.WalletID)
    wallet, err := services.Withdraw(walletUUID, request.Amount)
    if err != nil {
        return c.Status(400).JSON(models.ErrorResponse{Error: err.Error()})
    }

    return c.Status(200).JSON(wallet)
}
