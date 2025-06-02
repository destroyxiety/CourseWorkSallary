package services

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type PaymentsTaxesServiceInterface interface {
	GetAllPaymentsTaxes(ctx context.Context) ([]models.PaymentsTaxes, error)
	AddPaymentTax(ctx context.Context, taxID, paymentID int16) error
	DeletePaymentTax(ctx context.Context, taxID, paymentID int16) error
}
