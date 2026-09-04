package domain

import "time"

type ShuDistribution struct {
	ID                string    `json:"id"`
	PeriodYear        int       `json:"periodYear"`
	MemberID          string    `json:"memberId"`
	SavingsProportion string    `json:"savingsProportion"`
	LoanProportion    string    `json:"loanProportion"`
	TotalShu          string    `json:"totalShu"`
	CreatedAt         time.Time `json:"createdAt"`
}

type ShuRepository interface {
	GetByYear(year int) ([]ShuDistribution, error)
	Create(input CreateShuInput) (*ShuDistribution, error)
}

type ShuService interface {
	GetByYear(year int) ([]ShuDistribution, error)
	Calculate(year int) ([]ShuDistribution, error)
}

type CreateShuInput struct {
	PeriodYear        int    `json:"periodYear"`
	MemberID          string `json:"memberId"`
	SavingsProportion string `json:"savingsProportion"`
	LoanProportion    string `json:"loanProportion"`
	TotalShu          string `json:"totalShu"`
}
