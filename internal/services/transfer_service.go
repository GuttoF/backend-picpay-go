package services

import (
    "backend-picpay/internal/models"
    "errors"
    "gorm.io/gorm"
    "net/http"
)

func Transfer(db *gorm.DB, transaction *models.Transfer) error {
    tx := db.Begin()
    if tx.Error != nil {
        return tx.Error
    }

    var payerWallet models.Wallet
    if err := tx.Where("user_id = ?", transaction.PayerID).First(&payerWallet).Error; err != nil {
        tx.Rollback()
        return errors.New("payer wallet not found")
    }

    if payerWallet.Balance < transaction.Value {
        tx.Rollback()
        return errors.New("insufficient balance")
    }

    var payer models.User
    if err := tx.Where("id = ?", transaction.PayerID).First(&payer).Error; err != nil {
        tx.Rollback()
        return errors.New("payer not found")
    }
    if payer.UserType != "common" {
        tx.Rollback()
        return errors.New("only common users can transfer money")
    }

    var payee models.User
    if err := tx.Where("id = ?", transaction.PayeeID).First(&payee).Error; err != nil {
        tx.Rollback()
        return errors.New("payee not found")
    }

    resp, err := http.Get("https://util.devi.tools/api/v2/authorize")
    if err != nil || resp.StatusCode != http.StatusOK {
        tx.Rollback()
        return errors.New("transaction not authorized")
    }

    payerWallet.Balance -= transaction.Value
    if err := tx.Save(&payerWallet).Error; err != nil {
        tx.Rollback()
        return err
    }

    var payeeWallet models.Wallet
    if err := tx.Where("user_id = ?", transaction.PayeeID).First(&payeeWallet).Error; err != nil {
        tx.Rollback()
        return errors.New("payee wallet not found")
    }

    payeeWallet.Balance += transaction.Value
    if err := tx.Save(&payeeWallet).Error; err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Create(transaction).Error; err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Commit().Error; err != nil {
        return err
    }

    _, err = http.Post("https://util.devi.tools/api/v1/notify", "application/json", nil)
    if err != nil {
        return errors.New("failed to notify payee")
    }

    return nil
}
