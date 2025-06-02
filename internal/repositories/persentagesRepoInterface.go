package repositories

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type PercentagesRepoInterface interface {
	GetAllPercentages(ctx context.Context) ([]models.Percentages, error)
	AddPercent(ctx context.Context, employeeID, dealID int, percent float64) error
	DeletePercent(ctx context.Context, employeeID, dealID int) error
}
