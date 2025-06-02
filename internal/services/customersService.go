package services

import (
	"context"
	"fmt"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"github.com/destroyxiety/CourseWorkSallary/internal/repositories"
)

type customersService struct {
	customersService repositories.CustomersRepoInterface
}

func NewCustomersService(CR repositories.CustomersRepoInterface) CustomersServiceInterface {
	return &customersService{
		customersService: CR,
	}
}
func (CS *customersService) GetAllCustomers(ctx context.Context) ([]models.Customers, error) {
	return CS.customersService.GetAllCustomers(ctx)
}
func (CS *customersService) GetCustomersByAmount(ctx context.Context, dealAmount float64) ([]models.CustomersByAmount, error) {
	if dealAmount > 1000000000 || dealAmount < 0 {
		return nil, fmt.Errorf("the amount cannot be less than 0 or cannot be more than 10000000")
	}
	return CS.customersService.GetCustomersByAmount(ctx, dealAmount)
}
func (CS *customersService) AddCustomer(ctx context.Context, customerName string, phoneNumber string, email string) error {
	return CS.customersService.AddCustomer(ctx, customerName, phoneNumber, email)
}
func (CS *customersService) DeleteCustomer(ctx context.Context, customerID int) error {
	exists, err := CS.customersService.ExistsCustomer(ctx, customerID)
	if err != nil {
		return fmt.Errorf("checking customer existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("customer %d not found", customerID)
	}
	return CS.customersService.DeleteCustomer(ctx, customerID)
}
func (CS *customersService) UpdateCutomer(ctx context.Context, customerID int, customerName, phoneNumber, email *string) error {
	exists, err := CS.customersService.ExistsCustomer(ctx, customerID)
	if err != nil {
		return fmt.Errorf("checking customer existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("customer %d not found", customerID)
	}
	return CS.customersService.UpdateCutomer(ctx, customerID, customerName, phoneNumber, email)
}
