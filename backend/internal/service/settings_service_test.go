package service

import (
	"errors"
	"testing"

	"sp-backend/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockSettingRepo struct {
	mock.Mock
}

func (m *mockSettingRepo) GetAll() ([]domain.Setting, error) {
	args := m.Called()
	return args.Get(0).([]domain.Setting), args.Error(1)
}

func (m *mockSettingRepo) GetByKey(key string) (*domain.Setting, error) {
	args := m.Called(key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Setting), args.Error(1)
}

func (m *mockSettingRepo) Upsert(input domain.UpsertSettingInput) (*domain.Setting, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Setting), args.Error(1)
}

func TestSettingService_Upsert(t *testing.T) {
	repo := new(mockSettingRepo)
	svc := NewSettingService(repo)

	input := domain.UpsertSettingInput{Key: "bunga", Value: "12"}
	expected := &domain.Setting{ID: "1", Key: "bunga", Value: "12"}
	repo.On("Upsert", input).Return(expected, nil)

	got, err := svc.Upsert(input)
	assert.NoError(t, err)
	assert.Equal(t, "12", got.Value)
	repo.AssertExpectations(t)
}

func TestSettingService_GetAll_Error(t *testing.T) {
	repo := new(mockSettingRepo)
	svc := NewSettingService(repo)

	repo.On("GetAll").Return([]domain.Setting(nil), errors.New("db error"))

	_, err := svc.GetAll()
	assert.Error(t, err)
}
