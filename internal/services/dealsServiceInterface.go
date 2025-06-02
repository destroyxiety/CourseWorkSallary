package services

import (
	"context"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type DealsServiceInterface interface {
	GetDealsByDate(ctx context.Context, dealsDate time.Time) ([]models.DealsByDate, error)
	GetAllDeals(ctx context.Context) ([]models.Deals, error)
	AddDeals(ctx context.Context, dealDate time.Time, dealAmount float64, customerID int) error
	DeleteDeal(ctx context.Context, dealID int) error
}
