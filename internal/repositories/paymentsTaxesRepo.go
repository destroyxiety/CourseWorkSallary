package repositories

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"gorm.io/gorm"
)

type paymentTaxesRepo struct {
	DB *gorm.DB
}

func NewPaymentTaxesRepo(db *gorm.DB) *paymentTaxesRepo {
	return &paymentTaxesRepo{DB: db}
}

func (PTR *paymentTaxesRepo) GetAllPaymentsTaxes(ctx context.Context) ([]models.PaymentsTaxes, error) {
	var paymentsTaxes []models.PaymentsTaxes
	if err := PTR.DB.WithContext(ctx).Preload("Payment").Preload("Tax").Find(&paymentsTaxes).Error; err != nil {
		return nil, err
	}
	return paymentsTaxes, nil
}
func (PTR *paymentTaxesRepo) AddPaymentTax(ctx context.Context, taxID, paymentID int16) error {
	paymentsTaxes := models.PaymentsTaxes{
		TaxID:     taxID,
		PaymentID: paymentID,
	}
	return PTR.DB.Create(&paymentsTaxes).Error
}
func (PTR *paymentTaxesRepo) DeletePaymentTax(ctx context.Context, taxID, paymentID int16) error {
	return PTR.DB.WithContext(ctx).Delete(&models.PaymentsTaxes{TaxID: taxID, PaymentID: paymentID}).Error
}
func (PTR *paymentTaxesRepo) SumRatesByPaymentID(ctx context.Context, paymentID int) (float64, error) {
	var sumRates float64
	row := PTR.DB.WithContext(ctx).
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
