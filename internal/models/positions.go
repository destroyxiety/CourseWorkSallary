package models

import "gorm.io/gorm"

type Positions struct {
	PositionID    int16          `json:"position_id" gorm:"primaryKey;autoIncrement;not null"`
	PositionTitle string         `json:"position_title" gorm:"not null;unique;size:100"`
	MonthlySalary float64        `json:"monthly_salary" gorm:"not null;check:monthly_salary > 0;check:monthly_salary <= 10000000" `
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
