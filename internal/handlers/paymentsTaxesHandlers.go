package handlers

import (
	"net/http"
	"strconv"

	"github.com/destroyxiety/CourseWorkSallary/internal/services"
	"github.com/labstack/echo/v4"
)

func GetAllPaymentsTaxes(svc services.PaymentsTaxesServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		paymentsTaxes, err := svc.GetAllPaymentsTaxes(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, paymentsTaxes)
	}
}
func AddPaymentTax(svc services.PaymentsTaxesServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		paymentID, err := strconv.Atoi(c.Param("paymentID"))
		if err != nil || paymentID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid paymentID"})
		}
		taxID, err := strconv.Atoi(c.Param("taxID"))
		if err != nil || taxID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid taxID"})
		}
		if err := svc.AddPaymentTax(c.Request().Context(), int16(taxID), int16(paymentID)); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusCreated)
	}
}
func DeletePaymentTax(svc services.PaymentsTaxesServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		paymentID, err := strconv.Atoi(c.Param("paymentID"))
		if err != nil || paymentID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid paymentID"})
		}
		taxID, err := strconv.Atoi(c.Param("taxID"))
		if err != nil || taxID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid taxID"})
		}
		if err := svc.DeletePaymentTax(c.Request().Context(), int16(taxID), int16(paymentID)); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
