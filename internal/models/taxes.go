package models

type CodeType string

const (
	NDFL  CodeType = "NDFL"
	MIL   CodeType = "MIL"
	PFR   CodeType = "PFR"
	FOMS  CodeType = "FOMS"
	FSS   CodeType = "FSS"
	UIT   CodeType = "UIT"
	UNEMP CodeType = "UNEMP"
	ACC   CodeType = "ACC"
	VOLH  CodeType = "VOLH"
	PREF  CodeType = "PREF"
)

func (cd CodeType) IsValid() bool {
	switch cd {
	case NDFL,
		MIL,
		PFR,
		FOMS,
		FSS,
		UIT,
		UNEMP,
		ACC,
		VOLH,
		PREF:
		return true
	default:
		return false
	}
}

type TaxType string

const (
	IncomeTax                TaxType = "Income_Tax"
	MilitaryLevy             TaxType = "Military_Levy"
	PensionContribution      TaxType = "Pension_Contribution"
	HealthInsurance          TaxType = "Health_Insurance"
	SocialInsurance          TaxType = "Social_Insurance"
	WorkInsurance            TaxType = "Work_Injury_Insurance"
	UnemploymentInsurance    TaxType = "Unemployment_Insurance"
	AccidentInsurance        TaxType = "Accident_Insurance"
	VoluntaryInsurance       TaxType = "Voluntary_Health_Insurance"
	PreferentialContribution TaxType = "Preferential_Contribution"
)

func (tx TaxType) IsValid() bool {
	switch tx {
	case IncomeTax,
		MilitaryLevy,
		PensionContribution,
		HealthInsurance,
		SocialInsurance,
		WorkInsurance,
		UnemploymentInsurance,
		AccidentInsurance,
		VoluntaryInsurance,
		PreferentialContribution:
		return true
	default:
		return false
	}
}

var codeToTax = map[CodeType]TaxType{
	NDFL:  IncomeTax,
	MIL:   MilitaryLevy,
	PFR:   PensionContribution,
	FOMS:  HealthInsurance,
	FSS:   SocialInsurance,
	UIT:   WorkInsurance,
	UNEMP: UnemploymentInsurance,
	ACC:   AccidentInsurance,
	VOLH:  VoluntaryInsurance,
	PREF:  PreferentialContribution,
}

func (c CodeType) TaxName() (TaxType, bool) {
	t, ok := codeToTax[c]
	return t, ok
}

type Taxes struct {
	TaxID    int16    `json:"tax_id" gorm:"primaryKey;autoIncrement;not null"`
	Code     CodeType `json:"code" gorm:"not null;unique;size:15"`
	TaxTitle TaxType  `json:"tax_title" gorm:"not null;unique;size:100"`
	Rate     float64  `json:"rate" gorm:"not null;check:rate > 0;check:rate < 100"`
}
