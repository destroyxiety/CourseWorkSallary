package services

import (
	"context"
	"fmt"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"github.com/destroyxiety/CourseWorkSallary/internal/repositories"
)

type percentagesService struct {
	percentagesService repositories.PercentagesRepoInterface
	emploeesService    repositories.EmployeesRepoInterface
	dealsService       repositories.DealsRepoInterface
	accrualsService    repositories.AccrualsRepoInterface
	taxesService       repositories.TaxesRepoInterface
}

func NewPercentagesService(PeR repositories.PercentagesRepoInterface,
	ER repositories.EmployeesRepoInterface, DR repositories.DealsRepoInterface,
	AR repositories.AccrualsRepoInterface, TR repositories.TaxesRepoInterface) PercentagesServiceInterface {
	return &percentagesService{
		percentagesService: PeR,
		emploeesService:    ER,
		dealsService:       DR,
		accrualsService:    AR,
		taxesService:       TR,
	}
}
func (PeS *percentagesService) GetAllPercentages(ctx context.Context) ([]models.Percentages, error) {
	return PeS.percentagesService.GetAllPercentages(ctx)
}
func (PeS *percentagesService) AddPercent(ctx context.Context, employeeID, dealID int, percent float64) error {
	if percent <= 0 || percent >= 100 {
		return fmt.Errorf("the percentage should be strictly between 0 and 100")
	}

	existsEmp, err := PeS.emploeesService.ExistEmployees(ctx, employeeID)
	if err != nil {
		return fmt.Errorf("employee verification error: %w", err)
	}
	if !existsEmp {
		return fmt.Errorf("employee with ID=%d not found", employeeID)
	}

	existsDeal, err := PeS.dealsService.ExistsDeal(ctx, dealID)
	if err != nil {
		return fmt.Errorf("transaction verification error: %w", err)
	}
	if !existsDeal {
		return fmt.Errorf("deal with ID=%d not found", dealID)
	}

	if err := PeS.percentagesService.AddPercent(ctx, employeeID, dealID, percent); err != nil {
		return fmt.Errorf("couldn't add a percentage record: %w", err)
	}

	dealAmount, err := PeS.dealsService.GetDealAmount(ctx, dealID)
	if err != nil {
		return fmt.Errorf("couldn't get the transaction amount: %w", err)
	}

	const paymentID = 2
	sumTaxRate, err := PeS.taxesService.SumRatesByPaymentID(ctx, paymentID)
	if err != nil {
		return fmt.Errorf("couldn't get the amount of tax rates: %w", err)
	}

	p := percent / 100.0
	t := sumTaxRate / 100.0
	accrualAmount := dealAmount * p * (1 - t)

	if err := PeS.accrualsService.AddAccrual(ctx, employeeID, paymentID, time.Now(), accrualAmount); err != nil {
		return fmt.Errorf("failed to create an accrual record: %w", err)
	}

	return nil
}

func (PeS *percentagesService) DeletePercent(ctx context.Context, employeeID, dealID int) error {
	existEmployees, err := PeS.emploeesService.ExistEmployees(ctx, employeeID)
	if err != nil {
		return fmt.Errorf("checking employee existence: %w", err)
	}
	if !existEmployees {
		return fmt.Errorf("employee %d not found", employeeID)
	}
	existsDeal, err := PeS.dealsService.ExistsDeal(ctx, dealID)
	if err != nil {
		return fmt.Errorf("checking deal existence: %w", err)
	}
	if !existsDeal {
		return fmt.Errorf("deal %d not found", dealID)
	}
	return PeS.percentagesService.DeletePercent(ctx, employeeID, dealID)
}
