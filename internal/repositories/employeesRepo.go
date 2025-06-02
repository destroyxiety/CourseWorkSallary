package repositories

import (
	"context"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"gorm.io/gorm"
)

type employeesRepo struct {
	DB *gorm.DB
}

func NewEmployeesRepo(db *gorm.DB) *employeesRepo {
	return &employeesRepo{DB: db}
}

func (ER *employeesRepo) GetEmployeesBySalary(ctx context.Context, salary float64) ([]models.EmployeesBySalary, error) {
	var employees []models.EmployeesBySalary
	sql := `
	SELECT e.surname,
    e.name,
    p.position_title,
    p.monthly_salary
	FROM salary.employees e,
    salary.positions p
	WHERE e.position_id = p.position_id
 	AND p.monthly_salary > ?
	ORDER BY e.surname;`
	if err := ER.DB.WithContext(ctx).Raw(sql, salary).Scan(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}
func (ER *employeesRepo) GetEmployeesByAmount(ctx context.Context, salary float64, amount float64) ([]models.EmployeesByDeduction, error) {
	var employees []models.EmployeesByDeduction
	sql := `
	SELECT e.surname,
    e.name,
    pos.position_title,
    a.payment_amount AS deduction_amount
	FROM salary.employees   e,
    salary.positions   pos,
    salary.accruals    a,
    salary.payments    pay
	WHERE e.position_id      = pos.position_id
  	AND e.employee_id      = a.employee_id
  	AND a.payment_id       = pay.payment_id
 	AND pos.monthly_salary >= ?
 	AND a.payment_amount < ?
	ORDER BY deduction_amount;`
	if err := ER.DB.WithContext(ctx).Raw(sql, salary, -amount).Scan(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}
func (ER *employeesRepo) GetEmployeesByDeal(ctx context.Context, dealAmount float64, percent float64) ([]models.EmployeesByDeal, error) {
	var employees []models.EmployeesByDeal
	sql := `
	SELECT d.deal_id,
    d.deal_date,
    d.deal_amount,
    e.surname,
	e.name,
    perc.percent
	FROM salary.deals       d
	JOIN salary.percentages perc ON d.deal_id   = perc.deal_id
	JOIN salary.employees   e    ON perc.employee_id = e.employee_id
	WHERE d.deal_amount > ?
 	AND perc.percent   > ?
	ORDER BY perc.percent DESC;`
	if err := ER.DB.WithContext(ctx).Raw(sql, dealAmount, percent).Scan(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}
func (ER *employeesRepo) UpdatePositionEmployee(ctx context.Context, positionTitle string, employeeID int) error {
	sql := `
	UPDATE salary.employees e
	SET position_id = (
  	SELECT p.position_id
  	FROM salary.positions p
  	WHERE p.position_title = ?)
	WHERE e.employee_id = ?;`
	return ER.DB.WithContext(ctx).Exec(sql, positionTitle, employeeID).Error
}
func (ER *employeesRepo) GetEmployeesByProfit(ctx context.Context, dateStart, dateEnd time.Time) ([]models.EmployeesByProfit, error) {
	var employees []models.EmployeesByProfit
	sql := `SELECT
    e.employee_id,
    e.name,
    e.surname,
    SUM(d.deal_amount) AS total_deal_amount
	FROM
    salary.employees AS e
    JOIN salary.percentages AS p
      ON e.employee_id = p.employee_id
    JOIN salary.deals AS d
      ON p.deal_id = d.deal_id
	WHERE
    d.deal_date >= ?
    AND d.deal_date < ?
    AND d.deleted_at IS NULL
    AND e.deleted_at IS NULL
	GROUP BY
    e.employee_id,
    e.name,
    e.surname
	ORDER BY
    total_deal_amount DESC;`
	if err := ER.DB.WithContext(ctx).Raw(sql, dateStart, dateEnd).Scan(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}
func (ER *employeesRepo) GetEmployeesByTotalDeal(ctx context.Context, dateStart, dateEnd time.Time) ([]models.EmployeesByTotalDeal, error) {
	var employees []models.EmployeesByTotalDeal
	sql := `
	SELECT
    e.name,
    e.surname,
    SUM(a.payment_amount) AS total_amount
	FROM
    salary.employees AS e
    JOIN salary.accruals AS a
      ON e.employee_id = a.employee_id
	WHERE
    a.payment_date >= ?
    AND a.payment_date <  ?
    AND e.deleted_at IS NULL
	GROUP BY
    e.employee_id,
    e.name,
    e.surname
	ORDER BY
    total_amount DESC;`
	if err := ER.DB.WithContext(ctx).Raw(sql, dateStart, dateEnd).Scan(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}
func (ER *employeesRepo) GetAllEmployees(ctx context.Context) ([]models.Employees, error) {
	var employees []models.Employees
	if err := ER.DB.WithContext(ctx).Preload("Position").Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}
func (ER *employeesRepo) AddEmployees(ctx context.Context, name, surname, secondName string, positionID int16) error {
	employee := models.Employees{
		Name:       name,
		Surname:    surname,
		SecondName: secondName,
		PositionID: positionID,
	}
	return ER.DB.Create(&employee).Error
}
func (ER *employeesRepo) DeleteEmployees(ctx context.Context, employeeID int) error {
	return ER.DB.WithContext(ctx).Delete(&models.Employees{}, employeeID).Error
}
func (ER *employeesRepo) ExistEmployees(ctx context.Context, employeeID int) (bool, error) {
	var count int64
	err := ER.DB.WithContext(ctx).Model(&models.Employees{}).Where("employee_id = ?", employeeID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
