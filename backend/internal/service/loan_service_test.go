package service

import (
	"errors"
	"testing"

	"sp-backend/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockLoanRepo struct {
	mock.Mock
}

func (m *mockLoanRepo) GetAll() ([]domain.Loan, error) {
	args := m.Called()
	return args.Get(0).([]domain.Loan), args.Error(1)
}

func (m *mockLoanRepo) GetByID(id string) (*domain.Loan, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Loan), args.Error(1)
}

func (m *mockLoanRepo) GetByMemberID(memberID string) ([]domain.Loan, error) {
	args := m.Called(memberID)
	return args.Get(0).([]domain.Loan), args.Error(1)
}

func (m *mockLoanRepo) Create(input domain.CreateLoanInput) (*domain.Loan, error) {
	args := m.Called(input)
	return args.Get(0).(*domain.Loan), args.Error(1)
}

func (m *mockLoanRepo) UpdateStatus(id, status, userID string) (*domain.Loan, error) {
	args := m.Called(id, status, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Loan), args.Error(1)
}

func TestLoanService_Approve(t *testing.T) {
	repo := new(mockLoanRepo)
	svc := NewLoanService(repo)

	expected := &domain.Loan{ID: "1", Status: "approved"}
	repo.On("UpdateStatus", "1", "approved", "admin-1").Return(expected, nil)

	got, err := svc.Approve("1", "admin-1")
	assert.NoError(t, err)
	assert.Equal(t, "approved", got.Status)
	repo.AssertExpectations(t)
}

func TestLoanService_Disburse_Success(t *testing.T) {
	repo := new(mockLoanRepo)
	svc := NewLoanService(repo)

	repo.On("GetByID", "1").Return(&domain.Loan{ID: "1", Status: "approved"}, nil)
	repo.On("UpdateStatus", "1", "disbursed", "teller-1").Return(&domain.Loan{ID: "1", Status: "disbursed"}, nil)

	got, err := svc.Disburse("1", "teller-1")
	assert.NoError(t, err)
	assert.Equal(t, "disbursed", got.Status)
	repo.AssertExpectations(t)
}

func TestLoanService_Disburse_NotApproved(t *testing.T) {
	repo := new(mockLoanRepo)
	svc := NewLoanService(repo)

	repo.On("GetByID", "1").Return(&domain.Loan{ID: "1", Status: "pending"}, nil)

	_, err := svc.Disburse("1", "teller-1")
	assert.Error(t, err)
	assert.Equal(t, "hanya pinjaman yang disetujui yang bisa dicairkan", err.Error())
	repo.AssertNotCalled(t, "UpdateStatus", mock.Anything, mock.Anything, mock.Anything)
}

func TestLoanService_Disburse_GetError(t *testing.T) {
	repo := new(mockLoanRepo)
	svc := NewLoanService(repo)

	repo.On("GetByID", "1").Return(nil, errors.New("not found"))

	_, err := svc.Disburse("1", "teller-1")
	assert.Error(t, err)
}
