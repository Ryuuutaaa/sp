package service

import (
	"errors"
	"testing"

	"sp-backend/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockSavingsRepo struct {
	mock.Mock
}

func (m *mockSavingsRepo) GetTypes() ([]domain.SavingsType, error) {
	args := m.Called()
	return args.Get(0).([]domain.SavingsType), args.Error(1)
}

func (m *mockSavingsRepo) GetTransactions(memberID string) ([]domain.SavingsTransaction, error) {
	args := m.Called(memberID)
	return args.Get(0).([]domain.SavingsTransaction), args.Error(1)
}

func (m *mockSavingsRepo) CreateTransaction(input domain.CreateSavingsInput) (*domain.SavingsTransaction, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.SavingsTransaction), args.Error(1)
}

func TestSavingsService_Deposit(t *testing.T) {
	repo := new(mockSavingsRepo)
	svc := NewSavingsService(repo)

	input := domain.CreateSavingsInput{Amount: "100000", ReceiptNumber: "R-001"}
	expected := &domain.SavingsTransaction{ID: "1", Type: "setor", Amount: "100000"}

	repo.On("CreateTransaction", mock.MatchedBy(func(in domain.CreateSavingsInput) bool {
		return in.Type == "setor"
	})).Return(expected, nil)

	got, err := svc.Deposit(input)
	assert.NoError(t, err)
	assert.Equal(t, "setor", got.Type)
	repo.AssertExpectations(t)
}

func TestSavingsService_Withdraw_InvalidAmount(t *testing.T) {
	repo := new(mockSavingsRepo)
	svc := NewSavingsService(repo)

	_, err := svc.Withdraw(domain.CreateSavingsInput{Amount: "0"})
	assert.Error(t, err)
	assert.Equal(t, "jumlah penarikan tidak valid", err.Error())
	repo.AssertNotCalled(t, "CreateTransaction", mock.Anything)
}

func TestSavingsService_Withdraw_RepoError(t *testing.T) {
	repo := new(mockSavingsRepo)
	svc := NewSavingsService(repo)

	repo.On("CreateTransaction", mock.Anything).Return(nil, errors.New("db error"))

	_, err := svc.Withdraw(domain.CreateSavingsInput{Amount: "50000"})
	assert.Error(t, err)
	repo.AssertExpectations(t)
}
