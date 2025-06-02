package repositories

import (
	"context"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type PositionRepoInterface interface {
	GetCountPositions(ctx context.Context) ([]models.CountPositions, error)
	GetAllPositions(ctx context.Context) ([]models.Positions, error)
	AddPosition(ctx context.Context, position string, montlySalary float64) error
	UpdatePositionSalary(ctx context.Context, positionID int16, montlySalary float64) error
	DeletePosition(ctx context.Context, positionID int16) error
	ExistsPosition(ctx context.Context, positionID int16) (bool, error)
}
