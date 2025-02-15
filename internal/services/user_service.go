package services

import (
    "backend-picpay/internal/models"
    "gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
    if err := db.Create(user).Error; err != nil {
        return err
    }
    return nil
}
