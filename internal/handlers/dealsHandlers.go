package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/services"
	"github.com/labstack/echo/v4"
)

func GetDealsByDate(svc services.DealsServiceInterface) echo.HandlerFunc {
	type request struct {
		DealDate time.Time `json:"deal_date" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		deals, err := svc.GetDealsByDate(c.Request().Context(), req.DealDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, deals)
	}
}
func GetAllDeals(svc services.DealsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		deals, err := svc.GetAllDeals(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, deals)
	}
}
func AddDeals(svc services.DealsServiceInterface) echo.HandlerFunc {
	type request struct {
		DealDate    time.Time `json:"deal_date" validate:"required"`
		DealsAmount float64   `json:"deal_amount" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		customerID, err := strconv.Atoi(c.Param("customerID"))
		if err != nil || customerID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid customerID"})
		}
		if err := svc.AddDeals(c.Request().Context(), req.DealDate, req.DealsAmount, customerID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusCreated)
	}
}
func DeleteDeal(svc services.DealsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		dealID, err := strconv.Atoi(c.Param("dealID"))
		if err != nil || dealID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid dealID"})
		}
		if err := svc.DeleteDeal(c.Request().Context(), dealID); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
