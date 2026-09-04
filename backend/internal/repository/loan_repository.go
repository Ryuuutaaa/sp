package repository

import "sp-backend/internal/domain"

type LoanRepo struct {
	gql *GQLClient
}

func NewLoanRepo(gql *GQLClient) *LoanRepo {
	return &LoanRepo{gql: gql}
}

func (r *LoanRepo) GetAll() ([]domain.Loan, error) {
	var resp struct {
		Loans []domain.Loan `json:"loans"`
	}
	err := r.gql.Run(`query { loans { id memberId loanNumber amount interestRate interestType tenorMonths monthlyInstallment status approvedBy disbursedBy } }`, nil, &resp)
	return resp.Loans, err
}

func (r *LoanRepo) GetByID(id string) (*domain.Loan, error) {
	var resp struct {
		Loan *domain.Loan `json:"loan"`
	}
	err := r.gql.Run(`query($id: ID!) { loan(id: $id) { id memberId loanNumber amount interestRate interestType tenorMonths monthlyInstallment status approvedBy approvedAt disbursedBy disbursedAt } }`, map[string]interface{}{"id": id}, &resp)
	return resp.Loan, err
}

func (r *LoanRepo) Create(input domain.CreateLoanInput) (*domain.Loan, error) {
	var resp struct {
		CreateLoan *domain.Loan `json:"createLoan"`
	}
	err := r.gql.Run(`mutation($memberId: ID!, $loanNumber: String!, $amount: String!, $interestRate: String!, $interestType: String!, $tenorMonths: Int!, $monthlyInstallment: String!) {
		createLoan(memberId: $memberId, loanNumber: $loanNumber, amount: $amount, interestRate: $interestRate, interestType: $interestType, tenorMonths: $tenorMonths, monthlyInstallment: $monthlyInstallment) {
			id memberId loanNumber amount interestRate interestType tenorMonths monthlyInstallment status
		}
	}`, map[string]interface{}{
		"memberId":           input.MemberID,
		"loanNumber":         input.LoanNumber,
		"amount":             input.Amount,
		"interestRate":       input.InterestRate,
		"interestType":       input.InterestType,
		"tenorMonths":        input.TenorMonths,
		"monthlyInstallment": input.MonthlyInstallment,
	}, &resp)
	return resp.CreateLoan, err
}

func (r *LoanRepo) UpdateStatus(id string, status string, userID string) (*domain.Loan, error) {
	var resp struct {
		UpdateLoanStatus *domain.Loan `json:"updateLoanStatus"`
	}
	err := r.gql.Run(`mutation($id: ID!, $status: String!, $userId: ID!) {
		updateLoanStatus(id: $id, status: $status, userId: $userId) { id status approvedBy disbursedBy }
	}`, map[string]interface{}{"id": id, "status": status, "userId": userID}, &resp)
	return resp.UpdateLoanStatus, err
}
