package services

import (
    "backend-picpay/internal/models"
    "backend-picpay/internal/utils"
    "errors"
    "net/http"
)

func Transfer(transaction *models.Transaction) error {
    db := utils.ConnectDB()

    // Check if payer has enough balance
    var payerWallet models.Wallet
    if err := db.Where("user_id = ?", transaction.PayerID).First(&payerWallet).Error; err != nil {
        return err
    }

    if payerWallet.Balance < transaction.Value {
        return errors.New("insufficient balance")
    }

    // Check if payee exists
    var payee models.User
    if err := db.Where("id = ?", transaction.PayeeID).First(&payee).Error; err != nil {
        return err
    }

    var payer models.User
    if err := db.Where("id = ?", transaction.PayerID).First(&payer).Error; err != nil {
        return err
    }

    if payer.UserType != "common" {
        return errors.New("only common users can transfer money")
    }

    resp, err := http.Get("https://util.devi.tools/api/v2/authorize")
    if err != nil || resp.StatusCode != http.StatusOK {
        return errors.New("transaction not authorized")
    }

    payerWallet.Balance -= transaction.Value
    if err := db.Save(&payerWallet).Error; err != nil {
        return err
    }

    var payeeWallet models.Wallet
    if err := db.Where("user_id = ?", transaction.PayeeID).First(&payeeWallet).Error; err != nil {
        return err
    }

    payeeWallet.Balance += transaction.Value
    if err := db.Save(&payeeWallet).Error; err != nil {
        return err
    }

    if err := db.Create(transaction).Error; err != nil {
        return err
    }

    _, err = http.Post("https://util.devi.tools/api/v1/notify", "application/json", nil)
    if err != nil {
        return errors.New("failed to notify payee")
    }

    return nil
}