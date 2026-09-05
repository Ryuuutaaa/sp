package service

import (
	"testing"

	"sp-backend/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockInstallmentRepo struct {
	mock.Mock
}

func (m *mockInstallmentRepo) GetByLoanID(loanID string) ([]domain.Installment, error) {
	args := m.Called(loanID)
	return args.Get(0).([]domain.Installment), args.Error(1)
}

func (m *mockInstallmentRepo) Create(input domain.CreateInstallmentInput) (*domain.Installment, error) {
	args := m.Called(input)
	return args.Get(0).(*domain.Installment), args.Error(1)
}

func (m *mockInstallmentRepo) Pay(id string, input domain.PayInstallmentInput) (*domain.Installment, error) {
	args := m.Called(id, input)
	return args.Get(0).(*domain.Installment), args.Error(1)
}

func TestInstallmentService_GenerateSchedule(t *testing.T) {
	repo := new(mockInstallmentRepo)
	svc := NewInstallmentService(repo)

	// Pinjaman 12jt, bunga 12%, tenor 12 -> pokok 1jt/bln, bunga 120rb/bln
	loan := domain.Loan{ID: "loan-1", Amount: "12000000", InterestRate: "12", TenorMonths: 12}

	repo.On("Create", mock.Anything).Return(
		&domain.Installment{ID: "inst-1", LoanID: "loan-1", Status: "unpaid"},
		nil,
	)

	got, err := svc.GenerateSchedule(loan)
	assert.NoError(t, err)
	assert.Len(t, got, 12)
	repo.AssertNumberOfCalls(t, "Create", 12)
}
