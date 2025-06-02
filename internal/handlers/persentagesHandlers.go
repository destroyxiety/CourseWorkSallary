package handlers

import (
	"net/http"
	"strconv"

	"github.com/destroyxiety/CourseWorkSallary/internal/services"
	"github.com/labstack/echo/v4"
)

func GetAllPercentages(svc services.PercentagesServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		percentages, err := svc.GetAllPercentages(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, percentages)
	}
}
func AddPercent(svc services.PercentagesServiceInterface) echo.HandlerFunc {
	type request struct {
		Percent float64 `json:"percent" validate:"required"`
	}

	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "неверный JSON"})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		employeeID, err := strconv.Atoi(c.Param("employeeID"))
		if err != nil || employeeID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid employeeID"})
		}
		dealID, err := strconv.Atoi(c.Param("dealID"))
		if err != nil || dealID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid dealID"})
		}

		if err := svc.AddPercent(c.Request().Context(), employeeID, dealID, req.Percent); err != nil {

			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		return c.NoContent(http.StatusCreated)
	}
}
func DeletePercent(svc services.PercentagesServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		employeeID, err := strconv.Atoi(c.Param("employeeID"))
		if err != nil || employeeID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid employeeID"})
		}
		dealID, err := strconv.Atoi(c.Param("dealID"))
		if err != nil || dealID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid dealID"})
		}
		if err := svc.DeletePercent(c.Request().Context(), employeeID, dealID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
