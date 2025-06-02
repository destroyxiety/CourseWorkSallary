package models

type EmployeesByDeduction struct {
	Surname       string  `gorm:"column:surname" json:"surname"`
	Name          string  `gorm:"column:name" json:"name"`
	PositionTitle string  `gorm:"column:position_title" json:"position_title"`
	Deduction     float64 `gorm:"column:deduction_amount" json:"deduction_amount"`
}
