package models

import "time"

type Accruals struct {
	EmployeeID    int       `json:"employee_id" gorm:"primaryKey;not null"`
	Employee      Employees `json:"employee" gorm:"foreignKey:EmployeeID;references:EmployeeID"`
	PaymentID     int16     `json:"payment_id" gorm:"primaryKey;not null"`
	Payment       Payments  `json:"payment" gorm:"foreignKey:PaymentID;references:PaymentID"`
	PaymentDate   time.Time `json:"payment_date" gorm:"primaryKey;not null"`
	PaymentAmount float64   `json:"payment_amount" gorm:"not null;check:payment_amount >= -1000000000;check:payment_amount <= 1000000000"`
}
