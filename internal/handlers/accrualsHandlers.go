package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/services"
	"github.com/labstack/echo/v4"
)

func GetAllAccruals(svc services.AccrualsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		accruals, err := svc.GetAllAccruals(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, accruals)
	}
}
func AddAccrual(svc services.AccrualsServiceInterface) echo.HandlerFunc {
	type request struct {
		PaymentDate   time.Time `json:"payment_date" validate:"required"`
		PaymentAmount float64   `json:"payment_amount" validate:"required"`
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
		paymentID, err := strconv.Atoi(c.Param("paymentID"))
		if err != nil || paymentID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid paymentID"})
		}
		if err := svc.AddAccrual(c.Request().Context(), employeeID, int16(paymentID), req.PaymentDate, req.PaymentAmount); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusCreated)
	}
}
func DeleteAccrual(svc services.AccrualsServiceInterface) echo.HandlerFunc {
	type request struct {
		PaymentDate time.Time `json:"payment_date" validate:"required"`
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
		paymentID, err := strconv.Atoi(c.Param("paymentID"))
		if err != nil || paymentID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid paymentID"})
		}
		if err := svc.DeleteAccrual(c.Request().Context(), employeeID, int16(paymentID), req.PaymentDate); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
