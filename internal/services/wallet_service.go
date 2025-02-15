package services

import (
    "backend-picpay/internal/models"
    "errors"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

func CreateWallet(db *gorm.DB, userID uuid.UUID) (*models.Wallet, error) {
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

func GetWalletByUserID(db *gorm.DB, userID uuid.UUID) (*models.Wallet, error) {
    var wallet models.Wallet
    if err := db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
        return nil, errors.New("wallet not found")
    }

    return &wallet, nil
}

func Deposit(db *gorm.DB, walletID uuid.UUID, amount float64) (*models.Wallet, error) {
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

func Withdraw(db *gorm.DB, walletID uuid.UUID, amount float64) (*models.Wallet, error) {
    if amount <= 0 {
        return nil, errors.New("withdrawal amount must be greater than zero")
    }

    tx := db.Begin()
    if tx.Error != nil {
        return nil, tx.Error
    }

    var wallet models.Wallet
    if err := tx.Where("id = ?", walletID).First(&wallet).Error; err != nil {
        tx.Rollback()
        return nil, errors.New("wallet not found")
    }

    if wallet.Balance < amount {
        tx.Rollback()
        return nil, errors.New("insufficient balance")
    }

    wallet.Balance -= amount

    if err := tx.Save(&wallet).Error; err != nil {
        tx.Rollback()
        return nil, err
    }

    if err := tx.Commit().Error; err != nil {
        return nil, err
    }

    return &wallet, nil
}
