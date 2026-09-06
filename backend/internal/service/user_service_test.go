package service

import (
	"testing"

	"sp-backend/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepo struct {
	mock.Mock
}

func (m *mockUserRepo) GetAll() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *mockUserRepo) UpdateRole(id string, role string) (*domain.User, error) {
	args := m.Called(id, role)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestUserService_UpdateRole_Valid(t *testing.T) {
	repo := new(mockUserRepo)
	svc := NewUserService(repo)

	expected := &domain.User{ID: "1", Role: "teller"}
	repo.On("UpdateRole", "1", "teller").Return(expected, nil)

	got, err := svc.UpdateRole("1", "teller")
	assert.NoError(t, err)
	assert.Equal(t, "teller", got.Role)
	repo.AssertExpectations(t)
}

func TestUserService_UpdateRole_Invalid(t *testing.T) {
	repo := new(mockUserRepo)
	svc := NewUserService(repo)

	_, err := svc.UpdateRole("1", "presiden")
	assert.Error(t, err)
	assert.Equal(t, "invalid role", err.Error())
	repo.AssertNotCalled(t, "UpdateRole", mock.Anything, mock.Anything)
}
