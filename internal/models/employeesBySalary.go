package models

type EmployeesBySalary struct {
	Surname       string  `gorm:"column:surname" json:"surname"`
	Name          string  `gorm:"column:name" json:"name"`
	PositionTitle string  `gorm:"column:position_title" json:"position_title"`
	Salary        float64 `gorm:"column:monthly_salary" json:"monthly_salary"`
}
