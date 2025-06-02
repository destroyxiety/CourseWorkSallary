package services

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type TaxesServiceInterface interface {
	GetAllTaxes(ctx context.Context) ([]models.Taxes, error)
}
