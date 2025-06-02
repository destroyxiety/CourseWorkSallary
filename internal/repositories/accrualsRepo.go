package repositories

import (
	"context"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"gorm.io/gorm"
)

type accrualsRepo struct {
	DB *gorm.DB
}

func NewAccrualsRepo(db *gorm.DB) *accrualsRepo {
	return &accrualsRepo{DB: db}
}
func (AR *accrualsRepo) GetAllAccruals(ctx context.Context) ([]models.Accruals, error) {
	var accruals []models.Accruals
	if err := AR.DB.WithContext(ctx).Preload("Employee").Preload("Payment").Find(&accruals).Error; err != nil {
		return nil, err
	}
	return accruals, nil
}
func (AR *accrualsRepo) AddAccrual(ctx context.Context, employeeID int, paymentID int16, paymentDate time.Time, paymentAmount float64) error {
	accrual := models.Accruals{
		EmployeeID:    employeeID,
		PaymentID:     paymentID,
		PaymentDate:   paymentDate,
		PaymentAmount: paymentAmount,
	}
	return AR.DB.WithContext(ctx).Create(&accrual).Error

}
func (AR *accrualsRepo) DeleteAccrual(ctx context.Context, employeeID int, paymentID int16, paymentDate time.Time) error {
	return AR.DB.WithContext(ctx).Delete(&models.Accruals{EmployeeID: employeeID, PaymentID: paymentID, PaymentDate: paymentDate}).Error
}
