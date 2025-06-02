package repositories

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type CustomersRepoInterface interface {
	GetAllCustomers(ctx context.Context) ([]models.Customers, error)
	GetCustomersByAmount(ctx context.Context, dealAmount float64) ([]models.CustomersByAmount, error)
	AddCustomer(ctx context.Context, customerName string, phoneNumber string, email string) error
	DeleteCustomer(ctx context.Context, customerID int) error
	UpdateCutomer(ctx context.Context, customerID int, customerName, phone_number, email *string) error
	ExistsCustomer(ctx context.Context, customerID int) (bool, error)
}
