package services

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type PaymentsServiceInterface interface {
	GetAllPayments(ctx context.Context) ([]models.Payments, error)
}
