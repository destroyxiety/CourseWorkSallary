package repositories

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"gorm.io/gorm"
)

type persentagesRepo struct {
	DB *gorm.DB
}

func NewPersentagesRepo(db *gorm.DB) *persentagesRepo {
	return &persentagesRepo{DB: db}
}
func (PeR *persentagesRepo) GetAllPercentages(ctx context.Context) ([]models.Percentages, error) {
	var percentages []models.Percentages
	if err := PeR.DB.WithContext(ctx).Preload("Employee").Preload("Deal").Find(&percentages).Error; err != nil {
		return nil, err
	}
	return percentages, nil
}
func (PeR *persentagesRepo) AddPercent(ctx context.Context, employeeID, dealID int, percent float64) error {
	percents := models.Percentages{
		DealID:     dealID,
		EmployeeID: employeeID,
		Percent:    percent,
	}
	return PeR.DB.Create(&percents).Error
}
func (PeR *persentagesRepo) DeletePercent(ctx context.Context, employeeID, dealID int) error {
	return PeR.DB.WithContext(ctx).Delete(&models.Percentages{EmployeeID: employeeID, DealID: dealID}).Error
}
