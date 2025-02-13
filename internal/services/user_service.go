package services

import (
    "backend-picpay/internal/models"
    "backend-picpay/internal/utils"
)

func CreateUser(user *models.User) error {
    db := utils.ConnectDB()

    if err := db.Create(user).Error; err != nil {
        return err
    }

    return nil
}