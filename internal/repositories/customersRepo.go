package repositories

import (
	"context"
	"fmt"

	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"gorm.io/gorm"
)

type customersRepo struct {
	DB *gorm.DB
}

func NewCustomersRepo(db *gorm.DB) *customersRepo {
	return &customersRepo{DB: db}
}
func (CR *customersRepo) GetAllCustomers(ctx context.Context) ([]models.Customers, error) {
	var customers []models.Customers
	if err := CR.DB.WithContext(ctx).Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
func (CR *customersRepo) GetCustomersByAmount(ctx context.Context, dealAmount float64) ([]models.CustomersByAmount, error) {
	var customers []models.CustomersByAmount
	sql := `
	SELECT c.customer_name,
    COUNT(d.deal_id)   AS deals_count,
    SUM(d.deal_amount) AS total_amount
	FROM salary.customers c
	JOIN salary.deals     d ON c.customer_id = d.customer_id
	GROUP BY c.customer_name
	HAVING SUM(d.deal_amount) > ?
	ORDER BY total_amount DESC;`
	if err := CR.DB.WithContext(ctx).Raw(sql, dealAmount).Scan(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
func (CR *customersRepo) AddCustomer(ctx context.Context, customerName string, phoneNumber string, email string) error {
	customers := models.Customers{
		CustomerName: customerName,
		PhoneNumber:  phoneNumber,
		Email:        email,
	}
	return CR.DB.WithContext(ctx).Create(&customers).Error
}
func (CR *customersRepo) DeleteCustomer(ctx context.Context, customerID int) error {
	return CR.DB.WithContext(ctx).Delete(&models.Customers{}, customerID).Error
}
func (CR *customersRepo) UpdateCutomer(ctx context.Context, customerID int, customerName, phoneNumber, email *string) error {
	updates := make(map[string]interface{})
	if customerName != nil {
		updates["customer_name"] = *customerName
	}
	if phoneNumber != nil {
		updates["phone_number"] = *phoneNumber
	}
	if email != nil {
		updates["email"] = *email
	}
	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}
	return CR.DB.WithContext(ctx).Where("customer_id", customerID).Model(&models.Customers{}).Updates(updates).Error
}
func (CR *customersRepo) ExistsCustomer(ctx context.Context, customerID int) (bool, error) {
	var count int64
	err := CR.DB.WithContext(ctx).Model(&models.Customers{}).Where("customer_id = ?", customerID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
