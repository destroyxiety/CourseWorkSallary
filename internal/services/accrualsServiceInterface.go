package services

import (
	"context"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type AccrualsServiceInterface interface {
	GetAllAccruals(ctx context.Context) ([]models.Accruals, error)
	AddAccrual(ctx context.Context, employeeID int, paymentID int16, paymentDate time.Time, paymentAmount float64) error
	DeleteAccrual(ctx context.Context, employeeID int, paymentID int16, paymentDate time.Time) error
}
