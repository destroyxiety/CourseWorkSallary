package models

type PaymentType string

const (
	BasicSalary   PaymentType = "Basic_Salary"
	Bonus         PaymentType = "Bonus"
	Overtime      PaymentType = "Overtime"
	Commission    PaymentType = "Commission"
	Advance       PaymentType = "Advance"
	Allowance     PaymentType = "Allowance"
	Reimbursement PaymentType = "Reimbursement"
	Severance     PaymentType = "Severance"
	RetroPay      PaymentType = "Retro_Pay"
	Deduction     PaymentType = "Deduction"
	Fine          PaymentType = "Fine"
)

func (pt PaymentType) IsValid() bool {
	switch pt {
	case BasicSalary,
		Bonus,
		Overtime,
		Commission,
		Advance,
		Allowance,
		Reimbursement,
		Severance,
		RetroPay,
		Fine,
		Deduction:
		return true
	default:
		return false
	}
}

type Payments struct {
	PaymentID     int16       `json:"payment_id" gorm:"primaryKey;autoIncrement;not null"`
	TypeOfPayment PaymentType `json:"type_of_payment" gorm:"unique;not null;size200"`
}
