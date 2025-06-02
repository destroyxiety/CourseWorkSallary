package services

import (
	"context"
	"fmt"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"github.com/destroyxiety/CourseWorkSallary/internal/repositories"
)

type accrualsService struct {
	accrualsService  repositories.AccrualsRepoInterface
	paymentsService  repositories.PaymentsRepoInterface
	employeesService repositories.EmployeesRepoInterface
	taxesService     repositories.TaxesRepoInterface
}

func NewAccrualsService(AR repositories.AccrualsRepoInterface,
	PR repositories.PaymentsRepoInterface, ER repositories.EmployeesRepoInterface,
	TR repositories.TaxesRepoInterface) AccrualsServiceInterface {
	return &accrualsService{
		accrualsService:  AR,
		paymentsService:  PR,
		employeesService: ER,
		taxesService:     TR,
	}
}
func (AS *accrualsService) GetAllAccruals(ctx context.Context) ([]models.Accruals, error) {
	return AS.accrualsService.GetAllAccruals(ctx)
}
func (AS *accrualsService) AddAccrual(ctx context.Context, employeeID int, paymentID int16, paymentDate time.Time, paymentAmount float64) error {
	if paymentAmount < -1000000000 || paymentAmount > 1000000000 {
		return fmt.Errorf("the amount cannot be less than -1000000000 or cannot be more than 1000000000")
	}
	existEmployees, err := AS.employeesService.ExistEmployees(ctx, employeeID)
	if err != nil {
		return fmt.Errorf("checking employee existence: %w", err)
	}
	if !existEmployees {
		return fmt.Errorf("employee %d not found", employeeID)
	}
	existsPayment, err := AS.paymentsService.ExistPayment(ctx, paymentID)
	if err != nil {
		return fmt.Errorf("checking payment existence: %w", err)
	}
	if !existsPayment {
		return fmt.Errorf("payment %d not found", paymentID)
	}
	sumTaxRate, err := AS.taxesService.SumRatesByPaymentID(ctx, paymentID)
	if err != nil {
		return fmt.Errorf("couldn't get the amount of tax rates: %w", err)
	}
	procent := sumTaxRate / 100.0
	paymentAmount = paymentAmount * (1 - procent)
	return AS.accrualsService.AddAccrual(ctx, employeeID, paymentID, paymentDate, paymentAmount)
}
func (AS *accrualsService) DeleteAccrual(ctx context.Context, employeeID int, paymentID int16, paymentDate time.Time) error {
	existEmployees, err := AS.employeesService.ExistEmployees(ctx, employeeID)
	if err != nil {
		return fmt.Errorf("checking employee existence: %w", err)
	}
	if !existEmployees {
		return fmt.Errorf("employee %d not found", employeeID)
	}
	existsPayment, err := AS.paymentsService.ExistPayment(ctx, paymentID)
	if err != nil {
		return fmt.Errorf("checking payment existence: %w", err)
	}
	if !existsPayment {
		return fmt.Errorf("payment %d not found", paymentID)
	}
	return AS.accrualsService.DeleteAccrual(ctx, employeeID, paymentID, paymentDate)
}
