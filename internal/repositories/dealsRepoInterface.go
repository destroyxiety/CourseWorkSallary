package repositories

import (
	"context"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
)

type DealsRepoInterface interface {
	GetDealsByDate(ctx context.Context, dealsDate time.Time) ([]models.DealsByDate, error)
	GetAllDeals(ctx context.Context) ([]models.Deals, error)
	AddDeals(ctx context.Context, dealDate time.Time, dealAmount float64, customerID int) error
	DeleteDeal(ctx context.Context, dealID int) error
	ExistsDeal(ctx context.Context, dealID int) (bool, error)
	GetDealAmount(ctx context.Context, dealID int) (float64, error)
}
