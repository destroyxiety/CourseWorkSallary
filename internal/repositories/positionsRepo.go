package repositories

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"gorm.io/gorm"
)

type positionRepo struct {
	DB *gorm.DB
}

func NewPositionsRepo(db *gorm.DB) *positionRepo {
	return &positionRepo{DB: db}
}
func (PR *positionRepo) GetAllPositions(ctx context.Context) ([]models.Positions, error) {
	var positions []models.Positions
	if err := PR.DB.WithContext(ctx).Find(&positions).Error; err != nil {
		return nil, err
	}
	return positions, nil
}
func (PR *positionRepo) AddPosition(ctx context.Context, positionTitle string, montlySalary float64) error {
	position := models.Positions{
		PositionTitle: positionTitle,
		MonthlySalary: montlySalary,
	}
	return PR.DB.WithContext(ctx).Create(&position).Error
}
func (PR *positionRepo) UpdatePositionSalary(ctx context.Context, positionID int16, montlySalary float64) error {
	return PR.DB.WithContext(ctx).Where("position_id = ?", positionID).
		Model(&models.Positions{}).Update("monthly_salary", montlySalary).Error
}
func (PR *positionRepo) DeletePosition(ctx context.Context, positionID int16) error {
	return PR.DB.WithContext(ctx).Delete(&models.Positions{}, positionID).Error
}
func (PR *positionRepo) ExistsPosition(ctx context.Context, positionID int16) (bool, error) {
	var count int64
	err := PR.DB.WithContext(ctx).Model(&models.Positions{}).Where("position_id = ?", positionID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (PR *positionRepo) GetCountPositions(ctx context.Context) ([]models.CountPositions, error) {
	var customers []models.CountPositions
	sql := `
	SELECT
  	p.position_title,
  	COUNT(e.employee_id)    AS emp_count,
 	p.monthly_salary        AS salary
	FROM salary.positions AS p
	LEFT JOIN salary.employees  AS e
 	ON p.position_id = e.position_id
	GROUP BY
 	p.position_title,
 	monthly_salary
	ORDER BY   emp_count DESC;`
	if err := PR.DB.WithContext(ctx).Raw(sql).Scan(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
