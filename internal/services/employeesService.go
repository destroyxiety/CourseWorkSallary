package services

import (
	"context"
	"fmt"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"github.com/destroyxiety/CourseWorkSallary/internal/repositories"
)

type employeesService struct {
	employeesService repositories.EmployeesRepoInterface
	positionsService repositories.PositionRepoInterface
}

func NewEmployeesService(ER repositories.EmployeesRepoInterface, PS repositories.PositionRepoInterface) EmployeesServiceInterface {
	return &employeesService{
		employeesService: ER,
		positionsService: PS,
	}
}
func (ES *employeesService) GetEmployeesBySalary(ctx context.Context, salary float64) ([]models.EmployeesBySalary, error) {
	if salary < 0 || salary > 10000000 {
		return nil, fmt.Errorf("the salary cannot be less than or equal to 0 or cannot be more than 10000000")
	}
	return ES.employeesService.GetEmployeesBySalary(ctx, salary)
}
func (ES *employeesService) GetEmployeesByAmount(ctx context.Context, salary float64, amount float64) ([]models.EmployeesByDeduction, error) {
	if salary < 0 || salary > 10000000 {
		return nil, fmt.Errorf("the salary cannot be less or equal to 0 or cannot be more than 10000000")
	}
	if amount > 1000000000 || amount < 0 {
		return nil, fmt.Errorf("the amount cannot be less than 0 or cannot be more than 10000000")
	}
	return ES.employeesService.GetEmployeesByAmount(ctx, salary, amount)
}
func (ES *employeesService) GetEmployeesByDeal(ctx context.Context, dealAmount float64, percent float64) ([]models.EmployeesByDeal, error) {
	if percent <= 0 || percent >= 100 {
		return nil, fmt.Errorf("the percen cannot be less or equal to 0 or cannot be more or equal than 100")
	}
	if dealAmount <= 0 || dealAmount > 1000000000 {
		return nil, fmt.Errorf("the dealAmmount cannot be less or equal to 0 or cannot be more than 1000000000")
	}
	return ES.employeesService.GetEmployeesByDeal(ctx, dealAmount, percent)
}
func (ES *employeesService) UpdatePositionEmployee(ctx context.Context, positionTitle string, employeeID int) error {
	exists, err := ES.employeesService.ExistEmployees(ctx, employeeID)
	if err != nil {
		return fmt.Errorf("checking employee existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("employee %d not found", employeeID)
	}
	return ES.employeesService.UpdatePositionEmployee(ctx, positionTitle, employeeID)
}
func (ES *employeesService) GetEmployeesByTotalDeal(ctx context.Context, dateStart, dateEnd time.Time) ([]models.EmployeesByTotalDeal, error) {
	if dateStart.After(dateEnd) {
		return nil, fmt.Errorf("the start date must be before the end date")
	}
	return ES.employeesService.GetEmployeesByTotalDeal(ctx, dateStart, dateEnd)
}
func (ES *employeesService) GetEmployeesByProfit(ctx context.Context, dateStart, dateEnd time.Time) ([]models.EmployeesByProfit, error) {
	if dateStart.After(dateEnd) {
		return nil, fmt.Errorf("the start date must be before the end date")
	}
	return ES.employeesService.GetEmployeesByProfit(ctx, dateStart, dateEnd)
}
func (ES *employeesService) GetAllEmployees(ctx context.Context) ([]models.Employees, error) {
	return ES.employeesService.GetAllEmployees(ctx)
}
func (ES *employeesService) AddEmployees(ctx context.Context, name, surname, secondName string, positionID int16) error {
	exists, err := ES.positionsService.ExistsPosition(ctx, positionID)
	if err != nil {
		return fmt.Errorf("checking employee existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("employee %d not found", positionID)
	}
	return ES.employeesService.AddEmployees(ctx, name, surname, secondName, positionID)
}
func (ES *employeesService) DeleteEmployees(ctx context.Context, employeeID int) error {
	exists, err := ES.employeesService.ExistEmployees(ctx, employeeID)
	if err != nil {
		return fmt.Errorf("checking employee existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("employee %d not found", employeeID)
	}
	return ES.employeesService.DeleteEmployees(ctx, employeeID)
}
