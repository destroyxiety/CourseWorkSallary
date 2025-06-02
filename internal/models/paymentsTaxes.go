package models

type PaymentsTaxes struct {
	PaymentID int16    `json:"payment_id" gorm:"primaryKey;not null"`
	Payment   Payments `json:"payment" gorm:"foreignKey:PaymentID;references:PaymentID"`
	TaxID     int16    `json:"tax_id" gorm:"primaryKey;not null"`
	Tax       Taxes    `json:"tax" gorm:"foreignKey:TaxID;references:TaxID"`
}
