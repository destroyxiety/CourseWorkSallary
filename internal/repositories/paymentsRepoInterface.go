package repositories

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type PaymentsRepoInterface interface {
	GetAllPayments(ctx context.Context) ([]models.Payments, error)
	ExistPayment(ctx context.Context, paymentID int16) (bool, error)
}
