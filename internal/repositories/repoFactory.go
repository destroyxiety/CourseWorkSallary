package repositories

import "gorm.io/gorm"

type RepoFactory struct {
	Employees    EmployeesRepoInterface
	Positions    PositionRepoInterface
	Customers    CustomersRepoInterface
	Deals        DealsRepoInterface
	Payments     PaymentsRepoInterface
	Taxes        TaxesRepoInterface
	PaymentTaxes PaymentsTaxesRepoInterface
	Persentages  PercentagesRepoInterface
	Accruals     AccrualsRepoInterface
}

func NewRepoFactory(db *gorm.DB) *RepoFactory {
	return &RepoFactory{
		Employees:    NewEmployeesRepo(db),
		Positions:    NewPositionsRepo(db),
		Customers:    NewCustomersRepo(db),
		Deals:        NewDealsRepo(db),
		Payments:     NewPaymentsRepo(db),
		Taxes:        NewTaxesRepo(db),
		PaymentTaxes: NewPaymentTaxesRepo(db),
		Persentages:  NewPersentagesRepo(db),
		Accruals:     NewAccrualsRepo(db),
	}
}
