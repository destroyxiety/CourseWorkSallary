package repositories

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"gorm.io/gorm"
)

type taxesRepo struct {
	DB *gorm.DB
}

func NewTaxesRepo(db *gorm.DB) *taxesRepo {
	return &taxesRepo{DB: db}
}
func (TR *taxesRepo) GetAllTaxes(ctx context.Context) ([]models.Taxes, error) {
	var taxes []models.Taxes
	if err := TR.DB.WithContext(ctx).Find(&taxes).Error; err != nil {
		return nil, err
	}
	return taxes, nil
}
func (TR *taxesRepo) ExistTax(ctx context.Context, taxID int16) (bool, error) {
	var count int64
	err := TR.DB.WithContext(ctx).Model(&models.Taxes{}).Where("tax_id = ?", taxID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (TR *taxesRepo) SumRatesByPaymentID(ctx context.Context, paymentID int16) (float64, error) {
	var sumRates float64
	row := TR.DB.WithContext(ctx).
		Table("salary.payments_taxes AS pt").
		Select("COALESCE(SUM(t.rate), 0)").
		Joins("JOIN salary.taxes t ON pt.tax_id = t.tax_id").
		Where("pt.payment_id = ?", paymentID).
		Row()
	if err := row.Scan(&sumRates); err != nil {
		return 0, err
	}
	return sumRates, nil
}
