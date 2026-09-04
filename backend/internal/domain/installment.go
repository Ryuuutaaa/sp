package domain

import "time"

type Installment struct {
	ID                string     `json:"id"`
	LoanID            string     `json:"loanId"`
	InstallmentNumber int        `json:"installmentNumber"`
	DueDate           string     `json:"dueDate"`
	PrincipalAmount   string     `json:"principalAmount"`
	InterestAmount    string     `json:"interestAmount"`
	Penalty           string     `json:"penalty"`
	TotalAmount       string     `json:"totalAmount"`
	PaidAmount        string     `json:"paidAmount"`
	PaidAt            *time.Time `json:"paidAt"`
	PaidReceivedBy    *string    `json:"paidReceivedBy"`
	ReceiptNumber     *string    `json:"receiptNumber"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"createdAt"`
}

type InstallmentRepository interface {
	GetByLoanID(loanID string) ([]Installment, error)
	Create(input CreateInstallmentInput) (*Installment, error)
	Pay(id string, input PayInstallmentInput) (*Installment, error)
}

type InstallmentService interface {
	GetByLoanID(loanID string) ([]Installment, error)
	GenerateSchedule(loan Loan) ([]Installment, error)
	Pay(id string, input PayInstallmentInput) (*Installment, error)
}

type CreateInstallmentInput struct {
	LoanID            string `json:"loanId"`
	InstallmentNumber int    `json:"installmentNumber"`
	DueDate           string `json:"dueDate"`
	PrincipalAmount   string `json:"principalAmount"`
	InterestAmount    string `json:"interestAmount"`
	TotalAmount       string `json:"totalAmount"`
}

type PayInstallmentInput struct {
	PaidAmount     string `json:"paidAmount"`
	Penalty        string `json:"penalty"`
	ReceiptNumber  string `json:"receiptNumber"`
	PaidReceivedBy string `json:"paidReceivedBy"`
}
