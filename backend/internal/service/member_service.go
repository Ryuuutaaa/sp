package service

import "sp-backend/internal/domain"

type MemberService struct {
	repo domain.MemberRepository
}

func NewMemberService(repo domain.MemberRepository) *MemberService {
	return &MemberService{repo: repo}
}

func (s *MemberService) GetAll() ([]domain.Member, error) {
	return s.repo.GetAll()
}

func (s *MemberService) GetByID(id string) (*domain.Member, error) {
	return s.repo.GetByID(id)
}

func (s *MemberService) Register(input domain.CreateMemberInput) (*domain.Member, error) {
	// TODO: vlidasi NIK, KTP, dsb
	return s.repo.Create(input)
}

func (s *MemberService) Update(id string, input domain.UpdateMemberInput) (*domain.Member, error) {
	return s.repo.Update(id, input)
}

func (s *MemberService) Deactivate(id string) (*domain.Member, error) {
	return s.repo.UpdateStatus(id, "inactive")
}

func (s *MemberService) UpdateStatus(id string, status string) (*domain.Member, error) {
	return s.repo.UpdateStatus(id, status)
}

func (s *MemberService) Verify(id string, status string, verifiedBy string) (*domain.Member, error) {
	return s.repo.Verify(id, status, verifiedBy)
}
