package repositories

import (
	"context"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"gorm.io/gorm"
)

type dealsRepo struct {
	DB *gorm.DB
}

func NewDealsRepo(db *gorm.DB) *dealsRepo {
	return &dealsRepo{DB: db}
}
func (DR *dealsRepo) GetDealsByDate(ctx context.Context, dealsDate time.Time) ([]models.DealsByDate, error) {
	var deals []models.DealsByDate
	sql := `
	SELECT deal_id,
    deal_date,
    deal_amount
	FROM salary.deals
	WHERE deal_date < ?
	ORDER BY deal_date;`
	if err := DR.DB.WithContext(ctx).Raw(sql, dealsDate).Scan(&deals).Error; err != nil {
		return nil, err
	}
	return deals, nil
}
func (DR *dealsRepo) GetAllDeals(ctx context.Context) ([]models.Deals, error) {
	var deals []models.Deals
	if err := DR.DB.WithContext(ctx).Preload("Customer").Find(&deals).Error; err != nil {
		return nil, err
	}
	return deals, nil
}
func (DR *dealsRepo) AddDeals(ctx context.Context, dealDate time.Time, dealAmount float64, customerID int) error {
	deal := models.Deals{
		DealDate:   dealDate,
		DealAmount: dealAmount,
		CustomerID: customerID,
	}
	return DR.DB.WithContext(ctx).Create(&deal).Error
}
func (DR *dealsRepo) DeleteDeal(ctx context.Context, dealID int) error {
	sql := `
	DELETE FROM salary.deals
 	WHERE deal_id = ?;`
	return DR.DB.WithContext(ctx).Exec(sql, dealID).Error
}
func (DR *dealsRepo) ExistsDeal(ctx context.Context, dealID int) (bool, error) {
	var count int64
	err := DR.DB.WithContext(ctx).Model(&models.Deals{}).Where("deal_id = ?", dealID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (r *dealsRepo) GetDealAmount(ctx context.Context, dealID int) (float64, error) {
	var deal models.Deals
	if err := r.DB.WithContext(ctx).
		Select("deal_amount").
		Where("deal_id = ?", dealID).
		First(&deal).Error; err != nil {
		return 0, err
	}
	return deal.DealAmount, nil
}
