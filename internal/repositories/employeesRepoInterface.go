package repositories

import (
	"context"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type EmployeesRepoInterface interface {
	GetEmployeesBySalary(ctx context.Context, salary float64) ([]models.EmployeesBySalary, error)
	GetEmployeesByAmount(ctx context.Context, salary float64, amount float64) ([]models.EmployeesByDeduction, error)
	GetEmployeesByDeal(ctx context.Context, dealAmount float64, percent float64) ([]models.EmployeesByDeal, error)
	UpdatePositionEmployee(ctx context.Context, positionTitle string, employeeID int) error
	GetEmployeesByTotalDeal(ctx context.Context, dateStart, dateEnd time.Time) ([]models.EmployeesByTotalDeal, error)
	GetEmployeesByProfit(ctx context.Context, dateStart, dateEnd time.Time) ([]models.EmployeesByProfit, error)
	GetAllEmployees(ctx context.Context) ([]models.Employees, error)
	AddEmployees(ctx context.Context, name, surname, secondName string, positionID int16) error
	DeleteEmployees(ctx context.Context, employeeID int) error
	ExistEmployees(ctx context.Context, employeeID int) (bool, error)
}
