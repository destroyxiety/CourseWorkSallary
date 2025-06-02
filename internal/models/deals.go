package models

import (
	"time"

	"gorm.io/gorm"
)

type Deals struct {
	DealID     int            `json:"deal_id" gorm:"primaryKey;autoIncrement;not null"`
	DealDate   time.Time      `json:"deal_date" gorm:"not null"`
	DealAmount float64        `json:"deal_amount" gorm:"not null check:deal_amount > 0;check:deal_amount <= 1000000000"`
	CustomerID int            `json:"customer_id" gorm:"not null;"`
	Customer   Customers      `json:"customer" gorm:"foreignKey:CustomerID;references:CustomerID"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
