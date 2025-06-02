package models

type CountPositions struct {
	PositionTitle string  `gorm:"column:position_title" json:"position_title"`
	EMPCount      int64   `gorm:"column:emp_count" json:"emp_count"`
	Salary        float64 `gorm:"column:salary" json:"salary"`
}
