package models

type CustomersByAmount struct {
	CustomerName string  `gorm:"column:customer_name" json:"customer_name"`
	DealsCount   int64   `gorm:"column:deals_count" json:"deals_count"`
	TotalAmount  float64 `gorm:"column:total_amount" json:"total_amount"`
}
