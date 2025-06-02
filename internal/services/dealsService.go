package services

import (
	"context"
	"fmt"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"github.com/destroyxiety/CourseWorkSallary/internal/repositories"
)

type dealsService struct {
	dealsService     repositories.DealsRepoInterface
	customersService repositories.CustomersRepoInterface
}

func NewDealsService(DS repositories.DealsRepoInterface, CS repositories.CustomersRepoInterface) DealsServiceInterface {
	return &dealsService{
		dealsService:     DS,
		customersService: CS,
	}
}
func (DS *dealsService) GetDealsByDate(ctx context.Context, dealsDate time.Time) ([]models.DealsByDate, error) {
	return DS.dealsService.GetDealsByDate(ctx, dealsDate)
}
func (DS *dealsService) GetAllDeals(ctx context.Context) ([]models.Deals, error) {
	return DS.dealsService.GetAllDeals(ctx)
}
func (DS *dealsService) AddDeals(ctx context.Context, dealDate time.Time, dealAmount float64, customerID int) error {
	if dealAmount > 1000000000 || dealAmount < 0 {
		return fmt.Errorf("the amount cannot be less than 0 or cannot be more than 10000000")
	}
	exists, err := DS.customersService.ExistsCustomer(ctx, customerID)
	if err != nil {
		return fmt.Errorf("checking customer existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("customer %d not found", customerID)
	}
	return DS.dealsService.AddDeals(ctx, dealDate, dealAmount, customerID)
}
func (DS *dealsService) DeleteDeal(ctx context.Context, dealID int) error {
	exists, err := DS.dealsService.ExistsDeal(ctx, dealID)
	if err != nil {
		return fmt.Errorf("checking deal existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("deal %d not found", dealID)
	}
	return DS.dealsService.DeleteDeal(ctx, dealID)
}
