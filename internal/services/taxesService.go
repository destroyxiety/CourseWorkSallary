package services

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"github.com/destroyxiety/CourseWorkSallary/internal/repositories"
)

type taxesService struct {
	taxesService repositories.TaxesRepoInterface
}

func NewTaxesService(TS repositories.TaxesRepoInterface) *taxesService {
	return &taxesService{taxesService: TS}
}
func (TS *taxesService) GetAllTaxes(ctx context.Context) ([]models.Taxes, error) {
	return TS.taxesService.GetAllTaxes(ctx)
}
