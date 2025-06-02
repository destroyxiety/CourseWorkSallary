package models

import "time"

type EmployeesByDeal struct {
	DealID     int       `gorm:"column:deal_id" json:"deal_id"`
	DealDate   time.Time `gorm:"column:deal_date" json:"deal_date"`
	DealAmount float64   `gorm:"column:deal_amount" json:"deal_amount"`
	Surname    string    `gorm:"column:surname" json:"surname"`
	Name       string    `gorm:"column:name" json:"name"`
	Percent    float64   `gorm:"column:percent" json:"percent"`
}
