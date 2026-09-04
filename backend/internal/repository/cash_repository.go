package repository

import "sp-backend/internal/domain"

type CashRepo struct {
	gql *GQLClient
}

func NewCashRepo(gql *GQLClient) *CashRepo {
	return &CashRepo{gql: gql}
}

func (r *CashRepo) GetAll() ([]domain.CashTransaction, error) {
	var resp struct {
		CashTransactions []domain.CashTransaction `json:"cashTransactions"`
	}
	err := r.gql.Run(`query { cashTransactions { id type category amount description referenceType referenceId createdBy } }`, nil, &resp)
	return resp.CashTransactions, err
}

func (r *CashRepo) Create(input domain.CreateCashInput) (*domain.CashTransaction, error) {
	var resp struct {
		CreateCashTransaction *domain.CashTransaction `json:"createCashTransaction"`
	}
	vars := map[string]interface{}{
		"type":        input.Type,
		"category":    input.Category,
		"amount":      input.Amount,
		"description": input.Description,
		"createdBy":   input.CreatedBy,
	}
	if input.ReferenceType != nil {
		vars["referenceType"] = *input.ReferenceType
	}
	if input.ReferenceID != nil {
		vars["referenceId"] = *input.ReferenceID
	}
	err := r.gql.Run(`mutation($type: String!, $category: String!, $amount: String!, $description: String!, $referenceType: String, $referenceId: ID, $createdBy: ID!) {
		createCashTransaction(type: $type, category: $category, amount: $amount, description: $description, referenceType: $referenceType, referenceId: $referenceId, createdBy: $createdBy) {
			id type category amount description createdBy
		}
	}`, vars, &resp)
	return resp.CreateCashTransaction, err
}
