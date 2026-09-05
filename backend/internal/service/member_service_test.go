package service

import (
	"testing"

	"sp-backend/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockMemberRepo struct {
	mock.Mock
}

func (m *mockMemberRepo) GetAll() ([]domain.Member, error) {
	args := m.Called()
	return args.Get(0).([]domain.Member), args.Error(1)
}

func (m *mockMemberRepo) GetByID(id string) (*domain.Member, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Member), args.Error(1)
}

func (m *mockMemberRepo) Create(input domain.CreateMemberInput) (*domain.Member, error) {
	args := m.Called(input)
	return args.Get(0).(*domain.Member), args.Error(1)
}

func (m *mockMemberRepo) Update(id string, input domain.UpdateMemberInput) (*domain.Member, error) {
	args := m.Called(id, input)
	return args.Get(0).(*domain.Member), args.Error(1)
}

func (m *mockMemberRepo) UpdateStatus(id, status string) (*domain.Member, error) {
	args := m.Called(id, status)
	return args.Get(0).(*domain.Member), args.Error(1)
}

func (m *mockMemberRepo) Verify(id, status, verifiedBy string) (*domain.Member, error) {
	args := m.Called(id, status, verifiedBy)
	return args.Get(0).(*domain.Member), args.Error(1)
}

func TestMemberService_Deactivate(t *testing.T) {
	repo := new(mockMemberRepo)
	svc := NewMemberService(repo)

	expected := &domain.Member{ID: "1", Status: "inactive"}
	repo.On("UpdateStatus", "1", "inactive").Return(expected, nil)

	got, err := svc.Deactivate("1")
	assert.NoError(t, err)
	assert.Equal(t, "inactive", got.Status)
	repo.AssertExpectations(t)
}

func TestMemberService_Verify(t *testing.T) {
	repo := new(mockMemberRepo)
	svc := NewMemberService(repo)

	expected := &domain.Member{ID: "1", VerificationStatus: "verified"}
	repo.On("Verify", "1", "verified", "admin-1").Return(expected, nil)

	got, err := svc.Verify("1", "verified", "admin-1")
	assert.NoError(t, err)
	assert.Equal(t, "verified", got.VerificationStatus)
	repo.AssertExpectations(t)
}
