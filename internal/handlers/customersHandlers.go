package handlers

import (
	"net/http"
	"strconv"

	"github.com/destroyxiety/CourseWorkSallary/internal/services"
	"github.com/labstack/echo/v4"
)

func GetAllCustomers(svc services.CustomersServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		customers, err := svc.GetAllCustomers(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, customers)
	}
}
func GetCustomersByAmount(svc services.CustomersServiceInterface) echo.HandlerFunc {
	type request struct {
		DealAmount float64 `json:"deal_amount" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		customers, err := svc.GetCustomersByAmount(c.Request().Context(), req.DealAmount)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, customers)
	}
}
func AddCustomer(svc services.CustomersServiceInterface) echo.HandlerFunc {
	type request struct {
		CustomerName string `json:"customer_name" validate:"required"`
		PhoneNumber  string `json:"phone_number" validate:"required"`
		Email        string `json:"email"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := svc.AddCustomer(c.Request().Context(), req.CustomerName, req.PhoneNumber, req.Email); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusCreated)
	}
}
func DeleteCustomer(svc services.CustomersServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		customerID, err := strconv.Atoi(c.Param("customerID"))
		if err != nil || customerID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid customerID"})
		}
		if err := svc.DeleteCustomer(c.Request().Context(), customerID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
func UpdateCutomer(svc services.CustomersServiceInterface) echo.HandlerFunc {
	type request struct {
		CustomerName *string `json:"customer_name"`
		PhoneNumber  *string `json:"phone_number"`
		Email        *string `json:"email"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		customerID, err := strconv.Atoi(c.Param("customerID"))
		if err != nil || customerID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid customerID"})
		}
		if err := svc.UpdateCutomer(c.Request().Context(), customerID, req.CustomerName, req.PhoneNumber, req.Email); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		return c.NoContent(http.StatusOK)
	}
}
