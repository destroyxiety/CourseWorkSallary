package services

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type PositionsServiceInterface interface {
	GetAllPositions(ctx context.Context) ([]models.Positions, error)
	AddPosition(ctx context.Context, positionTitle string, montlySalary float64) error
	UpdatePositionSalary(ctx context.Context, positionID int16, montlySalary float64) error
	DeletePosition(ctx context.Context, positionID int16) error
	GetCountPositions(ctx context.Context) ([]models.CountPositions, error)
}
