package repositories

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type PaymentsTaxesRepoInterface interface {
	GetAllPaymentsTaxes(ctx context.Context) ([]models.PaymentsTaxes, error)
	AddPaymentTax(ctx context.Context, taxID, paymentID int16) error
	DeletePaymentTax(ctx context.Context, taxID, paymentID int16) error
	SumRatesByPaymentID(ctx context.Context, paymentID int) (float64, error)
}
