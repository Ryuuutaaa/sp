package domain

import "time"

type SavingsType struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	DefaultAmount string    `json:"defaultAmount"`
	IsMandatory   bool      `json:"isMandatory"`
	CreatedAt     time.Time `json:"createdAt"`
}

type SavingsTransaction struct {
	ID            string    `json:"id"`
	MemberID      string    `json:"memberId"`
	SavingsTypeID string    `json:"savingsTypeId"`
	Type          string    `json:"type"`
	Amount        string    `json:"amount"`
	BalanceAfter  string    `json:"balanceAfter"`
	Note          *string   `json:"note"`
	ReceiptNumber string    `json:"receiptNumber"`
	CreatedBy     string    `json:"createdBy"`
	CreatedAt     time.Time `json:"createdAt"`
}

type SavingsRepository interface {
	GetTypes() ([]SavingsType, error)
	GetTransactions(memberID string) ([]SavingsTransaction, error)
	CreateTransaction(input CreateSavingsInput) (*SavingsTransaction, error)
}

type SavingsService interface {
	GetTypes() ([]SavingsType, error)
	GetTransactions(memberID string) ([]SavingsTransaction, error)
	Deposit(input CreateSavingsInput) (*SavingsTransaction, error)
	Withdraw(input CreateSavingsInput) (*SavingsTransaction, error)
}

type CreateSavingsInput struct {
	MemberID      string  `json:"memberId"`
	SavingsTypeID string  `json:"savingsTypeId"`
	Type          string  `json:"type"`
	Amount        string  `json:"amount"`
	BalanceAfter  string  `json:"balanceAfter"`
	Note          *string `json:"note"`
	ReceiptNumber string  `json:"receiptNumber"`
	CreatedBy     string  `json:"createdBy"`
}
