package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/services"
	"github.com/labstack/echo/v4"
)

func GetEmployeesBySalary(svc services.EmployeesServiceInterface) echo.HandlerFunc {
	type request struct {
		Salary float64 `json:"salary" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		employee, err := svc.GetEmployeesBySalary(c.Request().Context(), req.Salary)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, employee)
	}
}
func GetEmployeesByAmount(svc services.EmployeesServiceInterface) echo.HandlerFunc {
	type request struct {
		Salary float64 `json:"salary" validate:"required"`
		Amount float64 `json:"amount" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		employee, err := svc.GetEmployeesByAmount(c.Request().Context(), req.Salary, req.Amount)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, employee)
	}
}
func GetEmployeesByDeal(svc services.EmployeesServiceInterface) echo.HandlerFunc {
	type request struct {
		DealAmount float64 `json:"deal_amount" validate:"required"`
		Percent    float64 `json:"percent" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		employee, err := svc.GetEmployeesByDeal(c.Request().Context(), req.DealAmount, req.Percent)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, employee)
	}
}
func UpdatePositionEmployee(svc services.EmployeesServiceInterface) echo.HandlerFunc {
	type request struct {
		PositionTitle string `json:"position_title" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		employeeID, err := strconv.Atoi(c.Param("employeeID"))
		if err != nil || employeeID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid employeeID"})
		}
		if err := svc.UpdatePositionEmployee(c.Request().Context(), req.PositionTitle, employeeID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
func GetEmployeesByTotalDeal(svc services.EmployeesServiceInterface) echo.HandlerFunc {
	type request struct {
		DateStart time.Time `json:"date_start" validate:"required"`
		DateEnd   time.Time `json:"date_end" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		employee, err := svc.GetEmployeesByTotalDeal(c.Request().Context(), req.DateStart, req.DateEnd)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, employee)
	}
}
func GetEmployeesByProfit(svc services.EmployeesServiceInterface) echo.HandlerFunc {
	type request struct {
		DateStart time.Time `json:"date_start" validate:"required"`
		DateEnd   time.Time `json:"date_end" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		employee, err := svc.GetEmployeesByProfit(c.Request().Context(), req.DateStart, req.DateEnd)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, employee)
	}
}
func GetAllEmployees(svc services.EmployeesServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		employee, err := svc.GetAllEmployees(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, employee)
	}
}
func AddEmployees(svc services.EmployeesServiceInterface) echo.HandlerFunc {
	type request struct {
		Name       string `json:"name" validate:"required"`
		Surname    string `json:"surname" validate:"required"`
		SecondName string `json:"second_name"`
	}
	return func(c echo.Context) error {
		positionID, err := strconv.Atoi(c.Param("positionID"))
		if err != nil || positionID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid positionID"})
		}
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := svc.AddEmployees(c.Request().Context(), req.Name, req.Surname, req.SecondName, int16(positionID)); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusCreated)
	}
}
func DeleteEmployees(svc services.EmployeesServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		employeeID, err := strconv.Atoi(c.Param("employeeID"))
		if err != nil || employeeID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid employeeID"})
		}
		if err := svc.DeleteEmployees(c.Request().Context(), employeeID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
