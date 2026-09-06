package service

import (
	"errors"
	"sp-backend/internal/domain"
)

type LoanService struct {
	repo domain.LoanRepository
}

func NewLoanService(repo domain.LoanRepository) *LoanService {
	return &LoanService{repo: repo}
}

func (s *LoanService) GetAll() ([]domain.Loan, error) {
	return s.repo.GetAll()
}

func (s *LoanService) GetByID(id string) (*domain.Loan, error) {
	return s.repo.GetByID(id)
}

func (s *LoanService) GetByMemberID(memberID string) ([]domain.Loan, error) {
	return s.repo.GetByMemberID(memberID)
}

func (s *LoanService) Apply(input domain.CreateLoanInput) (*domain.Loan, error) {
	// TODO: Validasi plafon, hitung cicilan bulanan (flat/anuitas)
	return s.repo.Create(input)
}

func (s *LoanService) Approve(id string, approvedBy string) (*domain.Loan, error) {
	return s.repo.UpdateStatus(id, "approved", approvedBy)
}

func (s *LoanService) Reject(id string, rejectedBy string) (*domain.Loan, error) {
	return s.repo.UpdateStatus(id, "rejected", rejectedBy)
}

func (s *LoanService) Disburse(id string, disbursedBy string) (*domain.Loan, error) {
	loan, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if loan.Status != "approved" {
		return nil, errors.New("hanya pinjaman yang disetujui yang bisa dicairkan")
	}
	return s.repo.UpdateStatus(id, "disbursed", disbursedBy)
}
