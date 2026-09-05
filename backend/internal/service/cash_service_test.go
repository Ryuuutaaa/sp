package service

import (
	"errors"
	"testing"

	"sp-backend/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCashRepo struct {
	mock.Mock
}

func (m *mockCashRepo) GetAll() ([]domain.CashTransaction, error) {
	args := m.Called()
	return args.Get(0).([]domain.CashTransaction), args.Error(1)
}

func (m *mockCashRepo) Create(input domain.CreateCashInput) (*domain.CashTransaction, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.CashTransaction), args.Error(1)
}

func TestCashService_GetAll(t *testing.T) {
	repo := new(mockCashRepo)
	svc := NewCashService(repo)

	expected := []domain.CashTransaction{
		{ID: "1", Type: "masuk", Amount: "100000"},
	}
	repo.On("GetAll").Return(expected, nil)

	got, err := svc.GetAll()
	assert.NoError(t, err)
	assert.Len(t, got, 1)
	repo.AssertExpectations(t)
}

func TestCashService_Record_Success(t *testing.T) {
	repo := new(mockCashRepo)
	svc := NewCashService(repo)

	input := domain.CreateCashInput{Type: "masuk", Category: "simpanan", Amount: "100000"}
	expected := &domain.CashTransaction{ID: "1", Type: "masuk", Amount: "100000"}
	repo.On("Create", input).Return(expected, nil)

	got, err := svc.Record(input)
	assert.NoError(t, err)
	assert.Equal(t, "masuk", got.Type)
	repo.AssertExpectations(t)
}

func TestCashService_Record_Error(t *testing.T) {
	repo := new(mockCashRepo)
	svc := NewCashService(repo)

	repo.On("Create", mock.Anything).Return(nil, errors.New("db error"))

	_, err := svc.Record(domain.CreateCashInput{})
	assert.Error(t, err)
}
