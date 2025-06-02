package models

import "gorm.io/gorm"

type Customers struct {
	CustomerID   int            `json:"customer_id" gorm:"primaryKey;autoIncrement;not null"`
	CustomerName string         `json:"customer_name"  gorm:"unique;not null;size:200"`
	PhoneNumber  string         `json:"phone_number" gorm:"unique;not null;size:30"`
	Email        string         `json:"email" gorm:"unique;size:200"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
