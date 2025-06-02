package services

import (
	"context"
	"fmt"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"github.com/destroyxiety/CourseWorkSallary/internal/repositories"
)

type positionsService struct {
	positionsService repositories.PositionRepoInterface
}

func NewPositionsService(PS repositories.PositionRepoInterface) *positionsService {
	return &positionsService{positionsService: PS}
}

func (PS *positionsService) GetAllPositions(ctx context.Context) ([]models.Positions, error) {
	return PS.positionsService.GetAllPositions(ctx)
}
func (PS *positionsService) AddPosition(ctx context.Context, positionTitle string, montlySalary float64) error {
	if montlySalary <= 0 || montlySalary > 10000000 {
		return fmt.Errorf("the salary cannot be less or equal to 0 or cannot be more than 10000000")
	}
	return PS.positionsService.AddPosition(ctx, positionTitle, montlySalary)
}
func (PS *positionsService) UpdatePositionSalary(ctx context.Context, positionID int16, montlySalary float64) error {
	if montlySalary <= 0 || montlySalary > 10000000 {
		return fmt.Errorf("the salary cannot be less or equal to 0 or cannot be more than 10000000")
	}
	exists, err := PS.positionsService.ExistsPosition(ctx, positionID)
	if err != nil {
		return fmt.Errorf("checking position existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("position %d not found", positionID)
	}
	return PS.positionsService.UpdatePositionSalary(ctx, positionID, montlySalary)
}
func (PS *positionsService) DeletePosition(ctx context.Context, positionID int16) error {
	exists, err := PS.positionsService.ExistsPosition(ctx, positionID)
	if err != nil {
		return fmt.Errorf("checking position existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("position %d not found", positionID)
	}
	return PS.positionsService.DeletePosition(ctx, positionID)
}
func (PS *positionsService) GetCountPositions(ctx context.Context) ([]models.CountPositions, error) {
	return PS.positionsService.GetCountPositions(ctx)
}
