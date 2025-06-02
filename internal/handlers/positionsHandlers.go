package handlers

import (
	"net/http"
	"strconv"

	"github.com/destroyxiety/CourseWorkSallary/internal/services"
	"github.com/labstack/echo/v4"
)

func GetAllPositions(svc services.PositionsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		positions, err := svc.GetAllPositions(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, positions)
	}
}
func AddPosition(svc services.PositionsServiceInterface) echo.HandlerFunc {
	type request struct {
		PositionTitle string  `json:"position_title" validate:"required"`
		MontlySalary  float64 `json:"monthly_salary" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := svc.AddPosition(c.Request().Context(), req.PositionTitle, req.MontlySalary); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusCreated)
	}
}
func UpdatePositionSalary(svc services.PositionsServiceInterface) echo.HandlerFunc {
	type request struct {
		MontlySalary float64 `json:"monthly_salary" validate:"required"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		positionID, err := strconv.Atoi(c.Param("positionID"))
		if err != nil || positionID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid positionID"})
		}
		if err := svc.UpdatePositionSalary(c.Request().Context(), int16(positionID), req.MontlySalary); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
func DeletePosition(svc services.PositionsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		positionID, err := strconv.Atoi(c.Param("positionID"))
		if err != nil || positionID <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid positionID"})
		}
		if err := svc.DeletePosition(c.Request().Context(), int16(positionID)); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.NoContent(http.StatusOK)
	}
}
func GetCountPositions(svc services.PositionsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		positions, err := svc.GetCountPositions(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, positions)
	}
}
