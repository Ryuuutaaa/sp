package repository

import "sp-backend/internal/domain"

type InstallmentRepo struct {
	gql *GQLClient
}

func NewInstallmentRepo(gql *GQLClient) *InstallmentRepo {
	return &InstallmentRepo{gql: gql}
}

func (r *InstallmentRepo) GetByLoanID(loanID string) ([]domain.Installment, error) {
	var resp struct {
		Installments []domain.Installment `json:"installments"`
	}
	err := r.gql.Run(`query($loanId: ID!) { installments(loanId: $loanId) { id loanId installmentNumber dueDate totalAmount paidAmount status } }`, map[string]interface{}{"loanId": loanID}, &resp)
	return resp.Installments, err
}

func (r *InstallmentRepo) Create(input domain.CreateInstallmentInput) (*domain.Installment, error) {
	var resp struct {
		CreateInstallment *domain.Installment `json:"createInstallment"`
	}
	err := r.gql.Run(`mutation($loanId: ID!, $installmentNumber: Int!, $dueDate: String!, $principalAmount: String!, $interestAmount: String!, $totalAmount: String!) {
		createInstallment(loanId: $loanId, installmentNumber: $installmentNumber, dueDate: $dueDate, principalAmount: $principalAmount, interestAmount: $interestAmount, totalAmount: $totalAmount) {
			id loanId installmentNumber status
		}
	}`, map[string]interface{}{
		"loanId":            input.LoanID,
		"installmentNumber": input.InstallmentNumber,
		"dueDate":           input.DueDate,
		"principalAmount":   input.PrincipalAmount,
		"interestAmount":    input.InterestAmount,
		"totalAmount":       input.TotalAmount,
	}, &resp)
	return resp.CreateInstallment, err
}

func (r *InstallmentRepo) Pay(id string, input domain.PayInstallmentInput) (*domain.Installment, error) {
	var resp struct {
		PayInstallment *domain.Installment `json:"payInstallment"`
	}
	err := r.gql.Run(`mutation($id: ID!, $paidAmount: String!, $penalty: String!, $receiptNumber: String!, $paidReceivedBy: ID!) {
		payInstallment(id: $id, paidAmount: $paidAmount, penalty: $penalty, receiptNumber: $receiptNumber, paidReceivedBy: $paidReceivedBy) {
			id status paidAt receiptNumber
		}
	}`, map[string]interface{}{
		"id":             id,
		"paidAmount":     input.PaidAmount,
		"penalty":        input.Penalty,
		"receiptNumber":  input.ReceiptNumber,
		"paidReceivedBy": input.PaidReceivedBy,
	}, &resp)
	return resp.PayInstallment, err
}
