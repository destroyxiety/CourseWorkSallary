package services

import (
	"context"
	"fmt"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"github.com/destroyxiety/CourseWorkSallary/internal/repositories"
)

type paymentsTaxesService struct {
	paymentsTaxesService repositories.PaymentsTaxesRepoInterface
	paymentsService      repositories.PaymentsRepoInterface
	taxesService         repositories.TaxesRepoInterface
}

func NewPaymentsTaxesService(PTR repositories.PaymentsTaxesRepoInterface,
	PS repositories.PaymentsRepoInterface, TS repositories.TaxesRepoInterface) PaymentsTaxesServiceInterface {
	return &paymentsTaxesService{
		paymentsTaxesService: PTR,
		paymentsService:      PS,
		taxesService:         TS,
	}
}
func (PTS *paymentsTaxesService) GetAllPaymentsTaxes(ctx context.Context) ([]models.PaymentsTaxes, error) {
	return PTS.paymentsTaxesService.GetAllPaymentsTaxes(ctx)
}
func (PTS *paymentsTaxesService) AddPaymentTax(ctx context.Context, taxID, paymentID int16) error {
	existsPayment, err := PTS.paymentsService.ExistPayment(ctx, paymentID)
	if err != nil {
		return fmt.Errorf("checking payment existence: %w", err)
	}
	if !existsPayment {
		return fmt.Errorf("payment %d not found", paymentID)
	}
	existsTax, err := PTS.taxesService.ExistTax(ctx, taxID)
	if err != nil {
		return fmt.Errorf("checking tax existence: %w", err)
	}
	if !existsTax {
		return fmt.Errorf("tax %d not found", taxID)
	}
	return PTS.paymentsTaxesService.AddPaymentTax(ctx, taxID, paymentID)
}
func (PTS *paymentsTaxesService) DeletePaymentTax(ctx context.Context, taxID, paymentID int16) error {
	existsPayment, err := PTS.paymentsService.ExistPayment(ctx, paymentID)
	if err != nil {
		return fmt.Errorf("checking payment existence: %w", err)
	}
	if !existsPayment {
		return fmt.Errorf("payment %d not found", paymentID)
	}
	existsTax, err := PTS.taxesService.ExistTax(ctx, taxID)
	if err != nil {
		return fmt.Errorf("checking tax existence: %w", err)
	}
	if !existsTax {
		return fmt.Errorf("tax %d not found", taxID)
	}
	return PTS.paymentsTaxesService.DeletePaymentTax(ctx, taxID, paymentID)
}
