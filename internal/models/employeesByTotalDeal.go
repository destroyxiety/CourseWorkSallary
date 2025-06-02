package models

type EmployeesByTotalDeal struct {
	Surname     string  `gorm:"column:surname" json:"surname"`
	Name        string  `gorm:"column:name" json:"name"`
	TotalAmount float64 `gorm:"column:total_amount" json:"total_amount"`
}
