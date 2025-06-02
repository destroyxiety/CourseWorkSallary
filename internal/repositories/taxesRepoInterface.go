package repositories

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type TaxesRepoInterface interface {
	GetAllTaxes(ctx context.Context) ([]models.Taxes, error)
	ExistTax(ctx context.Context, taxID int16) (bool, error)
	SumRatesByPaymentID(ctx context.Context, paymentID int16) (float64, error)
}
