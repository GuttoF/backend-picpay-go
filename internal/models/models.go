package models

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID			uuid.UUID		`gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name		string			`gorm:"not null"`
	Document	string			`gorm:"unique;not null"`
	Email		string			`gorm:"unique;not null"`
	Password	string			`gorm:"not null"`
	UserType	string			`gorm:"not null"`
	CreatedAt	time.Time		`gorm:"default:CURRENT_TIMESTAMP"`
}

type Wallet struct {
	ID			uuid.UUID		`gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID		uuid.UUID		`gorm:"type:uuid;not null"`
	Balance		float64			`gorm:"not null"`
	CreatedAt	time.Time		`gorm:"default:CURRENT_TIMESTAMP"`
}

type Transfer struct {
	ID			uuid.UUID		`gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Value		float64			`gorm:"not null"`
	PayerID		uuid.UUID		`gorm:"type:uuid;not null"`
	PayeeID		uuid.UUID		`gorm:"type:uuid;not null"`
	CreatedAt	time.Time		`gorm:"default:CURRENT_TIMESTAMP"`
}

type Notification struct {
	ID				uuid.UUID	`gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID			uuid.UUID	`gorm:"type:uuid;not null"`
	TransactionID	uuid.UUID	`gorm:"type:uuid;not null"`
	Message			string		`gorm:"not null"`
	Status			string		`gorm:"not null"`
	CreatedAt		time.Time	`gorm:"default:CURRENT_TIMESTAMP"`
}