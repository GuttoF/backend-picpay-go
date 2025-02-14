package services

import (
    "backend-picpay/internal/models"
    "backend-picpay/internal/utils"
    "errors"
    "github.com/google/uuid"
)

func CreateWallet(userID uuid.UUID) (*models.Wallet, error) {
    db := utils.ConnectDB()

    var existingWallet models.Wallet
    if err := db.Where("user_id = ?", userID).First(&existingWallet).Error; err == nil {
        return nil, errors.New("user already has a wallet")
    }

    wallet := models.Wallet{
        UserID:  userID,
        Balance: 0,
    }

    if err := db.Create(&wallet).Error; err != nil {
        return nil, err
    }

    return &wallet, nil
}

func GetWalletByUserID(userID uuid.UUID) (*models.Wallet, error) {
    db := utils.ConnectDB()

    var wallet models.Wallet
    if err := db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
        return nil, errors.New("wallet not found")
    }

    return &wallet, nil
}

func Deposit(walletID uuid.UUID, amount float64) (*models.Wallet, error) {
    db := utils.ConnectDB()

    if amount <= 0 {
        return nil, errors.New("amount must be greater than zero")
    }

    var wallet models.Wallet
    if err := db.Where("id = ?", walletID).First(&wallet).Error; err != nil {
        return nil, errors.New("wallet not found")
    }

    wallet.Balance += amount
    if err := db.Save(&wallet).Error; err != nil {
        return nil, err
    }

    return &wallet, nil
}

func Withdraw(walletID uuid.UUID, amount float64) (*models.Wallet, error) {
    db := utils.ConnectDB()

    if amount <= 0 {
        return nil, errors.New("withdrawal amount must be greater than zero")
    }

    var wallet models.Wallet
    if err := db.Where("id = ?", walletID).First(&wallet).Error; err != nil {
        return nil, errors.New("wallet not found")
    }

    if wallet.Balance < amount {
        return nil, errors.New("insufficient balance")
    }

    wallet.Balance -= amount
    if err := db.Save(&wallet).Error; err != nil {
        return nil, err
    }

    return &wallet, nil
}
