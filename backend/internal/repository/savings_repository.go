package repository

import "sp-backend/internal/domain"

type SavingsRepo struct {
	gql *GQLClient
}

func NewSavingsRepo(gql *GQLClient) *SavingsRepo {
	return &SavingsRepo{gql: gql}
}

func (r *SavingsRepo) GetTypes() ([]domain.SavingsType, error) {
	var resp struct {
		SavingsTypes []domain.SavingsType `json:"savingsTypes"`
	}
	err := r.gql.Run(`query { savingsTypes { id name defaultAmount isMandatory } }`, nil, &resp)
	return resp.SavingsTypes, err
}

func (r *SavingsRepo) GetTransactions(memberID string) ([]domain.SavingsTransaction, error) {
	var resp struct {
		SavingsTransactions []domain.SavingsTransaction `json:"savingsTransactions"`
	}
	err := r.gql.Run(`query($memberId: ID) { savingsTransactions(memberId: $memberId) { id memberId savingsTypeId type amount balanceAfter note receiptNumber createdBy } }`, map[string]interface{}{"memberId": memberID}, &resp)
	return resp.SavingsTransactions, err
}

func (r *SavingsRepo) CreateTransaction(input domain.CreateSavingsInput) (*domain.SavingsTransaction, error) {
	var resp struct {
		CreateSavingsTransaction *domain.SavingsTransaction `json:"createSavingsTransaction"`
	}
	err := r.gql.Run(`mutation($memberId: ID!, $savingsTypeId: ID!, $type: String!, $amount: String!, $balanceAfter: String!, $note: String, $receiptNumber: String!, $createdBy: ID!) {
		createSavingsTransaction(memberId: $memberId, savingsTypeId: $savingsTypeId, type: $type, amount: $amount, balanceAfter: $balanceAfter, note: $note, receiptNumber: $receiptNumber, createdBy: $createdBy) {
			id memberId savingsTypeId type amount balanceAfter note receiptNumber createdBy
		}
	}`, map[string]interface{}{
		"memberId":      input.MemberID,
		"savingsTypeId": input.SavingsTypeID,
		"type":          input.Type,
		"amount":        input.Amount,
		"balanceAfter":  input.BalanceAfter,
		"note":          input.Note,
		"receiptNumber": input.ReceiptNumber,
		"createdBy":     input.CreatedBy,
	}, &resp)
	return resp.CreateSavingsTransaction, err
}
