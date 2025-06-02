package services

import "github.com/destroyxiety/CourseWorkSallary/internal/repositories"

type ServicesFactory struct {
	EmploeesService      EmployeesServiceInterface
	PositionsService     PositionsServiceInterface
	CustomersService     CustomersServiceInterface
	DealsService         DealsServiceInterface
	PaymentsService      PaymentsServiceInterface
	TaxesService         TaxesServiceInterface
	PaymentsTaxesService PaymentsTaxesServiceInterface
	PercentagesService   PercentagesServiceInterface
	AccrualsService      AccrualsServiceInterface
}

func NewServicesFactory(repoFactory *repositories.RepoFactory) *ServicesFactory {
	return &ServicesFactory{
		EmploeesService:      NewEmployeesService(repoFactory.Employees, repoFactory.Positions),
		PositionsService:     NewPositionsService(repoFactory.Positions),
		CustomersService:     NewCustomersService(repoFactory.Customers),
		DealsService:         NewDealsService(repoFactory.Deals, repoFactory.Customers),
		PaymentsService:      NewPaymentsService(repoFactory.Payments),
		TaxesService:         NewTaxesService(repoFactory.Taxes),
		PaymentsTaxesService: NewPaymentsTaxesService(repoFactory.PaymentTaxes, repoFactory.Payments, repoFactory.Taxes),
		PercentagesService: NewPercentagesService(repoFactory.Persentages, repoFactory.Employees, repoFactory.Deals,
			repoFactory.Accruals, repoFactory.Taxes),
		AccrualsService: NewAccrualsService(repoFactory.Accruals, repoFactory.Payments, repoFactory.Employees, repoFactory.Taxes),
	}
}
