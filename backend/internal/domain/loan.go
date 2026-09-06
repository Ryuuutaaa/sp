package domain

import "time"

type Loan struct {
	ID                 string     `json:"id"`
	MemberID           string     `json:"memberId"`
	LoanNumber         string     `json:"loanNumber"`
	Amount             string     `json:"amount"`
	InterestRate       string     `json:"interestRate"`
	InterestType       string     `json:"interestType"`
	TenorMonths        int        `json:"tenorMonths"`
	MonthlyInstallment string     `json:"monthlyInstallment"`
	Status             string     `json:"status"`
	ApprovedBy         *string    `json:"approvedBy"`
	ApprovedAt         *time.Time `json:"approvedAt"`
	DisbursedBy        *string    `json:"disbursedBy"`
	DisbursedAt        *time.Time `json:"disbursedAt"`
	CreatedAt          time.Time  `json:"createdAt"`
	UpdatedAt          time.Time  `json:"updatedAt"`
}

type LoanRepository interface {
	GetAll() ([]Loan, error)
	GetByMemberID(memberID string) ([]Loan, error)
	GetByID(id string) (*Loan, error)
	Create(input CreateLoanInput) (*Loan, error)
	UpdateStatus(id string, status string, userID string) (*Loan, error)
}

type LoanService interface {
	GetAll() ([]Loan, error)
	GetByMemberID(memberID string) ([]Loan, error)
	GetByID(id string) (*Loan, error)
	Apply(input CreateLoanInput) (*Loan, error)
	Approve(id string, approvedBy string) (*Loan, error)
	Reject(id string, rejectedBy string) (*Loan, error)
	Disburse(id string, disbursedBy string) (*Loan, error)
}

type CreateLoanInput struct {
	MemberID           string `json:"memberId"`
	LoanNumber         string `json:"loanNumber"`
	Amount             string `json:"amount"`
	InterestRate       string `json:"interestRate"`
	InterestType       string `json:"interestType"`
	TenorMonths        int    `json:"tenorMonths"`
	MonthlyInstallment string `json:"monthlyInstallment"`
}
