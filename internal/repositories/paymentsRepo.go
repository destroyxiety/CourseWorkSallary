package repositories

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"gorm.io/gorm"
)

type paymentsRepo struct {
	DB *gorm.DB
}

func NewPaymentsRepo(db *gorm.DB) *paymentsRepo {
	return &paymentsRepo{DB: db}
}
func (PR *paymentsRepo) GetAllPayments(ctx context.Context) ([]models.Payments, error) {
	var payments []models.Payments
	if err := PR.DB.WithContext(ctx).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}
func (PR *paymentsRepo) ExistPayment(ctx context.Context, paymentID int16) (bool, error) {
	var count int64
	err := PR.DB.WithContext(ctx).Model(&models.Payments{}).Where("payment_id = ?", paymentID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
