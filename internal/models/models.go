package models

import (
	"time"
)

type User struct {
	ID			uint		`gorm:"primaryKey"`
	Name		string		`gorm:"not null"`
	Document	string		`gorm:"unique;not null"`
	Email		string		`gorm:"unique;not null"`
	Password	string		`gorm:"not null"`
	UserType	string		`gorm:"not null"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP"`
}

type Wallet struct {
	ID			uint		`gorm:"primaryKey"`
	UserID		uint		`gorm:"not null"`
	Balance		float64		`gorm:"not null"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP"`
}

type Transaction struct {
	ID			uint		`gorm:"primaryKey"`
	Value		float64		`gorm:"not null"`
	PayeerID	uint		`gorm:"not null"`
	PayeeID		uint		`gorm:"not null"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP"`
}

type Notification struct {
	ID				uint		`gorm:"primaryKey"`
	UserID			uint		`gorm:"not null"`
	TransactionID	uint		`gorm:"not null"`
	Message			string		`gorm:"not null"`
	Status			string		`gorm:"not null"`
	CreatedAt		time.Time	`gorm:"default:CURRENT_TIMESTAMP"`
}