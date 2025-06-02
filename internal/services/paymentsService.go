package services

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"github.com/destroyxiety/CourseWorkSallary/internal/repositories"
)

type paymentsService struct {
	paymentsService repositories.PaymentsRepoInterface
}

func NewPaymentsService(PaR repositories.PaymentsRepoInterface) PaymentsServiceInterface {
	return &paymentsService{
		paymentsService: PaR,
	}
}
func (PaR *paymentsService) GetAllPayments(ctx context.Context) ([]models.Payments, error) {
	return PaR.paymentsService.GetAllPayments(ctx)
}
