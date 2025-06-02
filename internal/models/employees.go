package models

import "gorm.io/gorm"

type Employees struct {
	EmployeeID int            `json:"employee_id" gorm:"primaryKey;autoIncrement;not null"`
	Name       string         `json:"name" gorm:"not null;size:200"`
	Surname    string         `json:"surname" gorm:"not null;size:200"`
	SecondName string         `json:"second_name" gorm:"size:200"`
	PositionID int16          `json:"position_id" gorm:"not null"`
	Position   Positions      `json:"position" gorm:"foreignKey:PositionID;references:PositionID"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
