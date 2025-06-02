package models

type Percentages struct {
	EmployeeID int       `json:"employee_id" gorm:"primaryKey;not null"`
	Employee   Employees `json:"employee" gorm:"foreignKey:EmployeeID;references:EmployeeID"`
	DealID     int       `json:"deal_id" gorm:"primaryKey;not null"`
	Deal       Deals     `json:"deal" gorm:"foreignKey:DealID;references:DealID"`
	Percent    float64   `json:"percent" gorm:"check:percent > 0;check:percent < 100"`
}
