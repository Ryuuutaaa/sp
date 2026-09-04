package domain

import "time"

type CashTransaction struct {
	ID            string    `json:"id"`
	Type          string    `json:"type"`
	Category      string    `json:"category"`
	Amount        string    `json:"amount"`
	Description   string    `json:"description"`
	ReferenceType *string   `json:"referenceType"`
	ReferenceID   *string   `json:"referenceId"`
	CreatedBy     string    `json:"createdBy"`
	CreatedAt     time.Time `json:"createdAt"`
}

type CashRepository interface {
	GetAll() ([]CashTransaction, error)
	Create(input CreateCashInput) (*CashTransaction, error)
}

type CashService interface {
	GetAll() ([]CashTransaction, error)
	Record(input CreateCashInput) (*CashTransaction, error)
}

type CreateCashInput struct {
	Type          string  `json:"type"`
	Category      string  `json:"category"`
	Amount        string  `json:"amount"`
	Description   string  `json:"description"`
	ReferenceType *string `json:"referenceType"`
	ReferenceID   *string `json:"referenceId"`
	CreatedBy     string  `json:"createdBy"`
}
