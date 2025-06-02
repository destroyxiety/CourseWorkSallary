package handlers

import (
	"net/http"

	"github.com/destroyxiety/CourseWorkSallary/internal/services"
	"github.com/labstack/echo/v4"
)

func GetAllTaxes(svc services.TaxesServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		taxes, err := svc.GetAllTaxes(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, taxes)
	}
}
