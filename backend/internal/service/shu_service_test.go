package service

import (
	"errors"
	"testing"

	"sp-backend/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockShuRepo struct {
	mock.Mock
}

func (m *mockShuRepo) GetByYear(year int) ([]domain.ShuDistribution, error) {
	args := m.Called(year)
	return args.Get(0).([]domain.ShuDistribution), args.Error(1)
}

func (m *mockShuRepo) Create(input domain.CreateShuInput) (*domain.ShuDistribution, error) {
	args := m.Called(input)
	return args.Get(0).(*domain.ShuDistribution), args.Error(1)
}

func TestShuService_GetByYear(t *testing.T) {
	repo := new(mockShuRepo)
	svc := NewShuService(repo)

	expected := []domain.ShuDistribution{
		{ID: "1", PeriodYear: 2026, TotalShu: "500000"},
	}
	repo.On("GetByYear", 2026).Return(expected, nil)

	got, err := svc.GetByYear(2026)
	assert.NoError(t, err)
	assert.Len(t, got, 1)
	assert.Equal(t, 2026, got[0].PeriodYear)
	repo.AssertExpectations(t)
}

func TestShuService_Calculate(t *testing.T) {
	repo := new(mockShuRepo)
	svc := NewShuService(repo)

	expected := []domain.ShuDistribution{
		{ID: "1", PeriodYear: 2026, TotalShu: "500000"},
	}
	repo.On("GetByYear", 2026).Return(expected, nil)

	got, err := svc.Calculate(2026)
	assert.NoError(t, err)
	assert.Len(t, got, 1)
	repo.AssertExpectations(t)
}

func TestShuService_GetByYear_Error(t *testing.T) {
	repo := new(mockShuRepo)
	svc := NewShuService(repo)

	repo.On("GetByYear", 2026).Return([]domain.ShuDistribution(nil), errors.New("db error"))

	_, err := svc.GetByYear(2026)
	assert.Error(t, err)
}
