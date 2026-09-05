package service

import "sp-backend/internal/domain"

type CashService struct {
	repo domain.CashRepository
}

func NewCashService(repo domain.CashRepository) *CashService {
	return &CashService{repo: repo}
}

func (s *CashService) GetAll() ([]domain.CashTransaction, error) {
	return s.repo.GetAll()
}

func (s *CashService) Record(input domain.CreateCashInput) (*domain.CashTransaction, error) {
	return s.repo.Create(input)
}
