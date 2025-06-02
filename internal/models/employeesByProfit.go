package models

type EmployeesByProfit struct {
	EmployeeID      int     `gorm:"column:employee_id" json:"employee_id"`
	Name            string  `gorm:"column:name" json:"name"`
	Surname         string  `gorm:"column:surname" json:"surname"`
	TotalDealAmount float64 `gorm:"column:total_deal_amount" json:"total_deal_amount"`
}
