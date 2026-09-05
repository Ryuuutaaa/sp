package service

import "sp-backend/internal/domain"

type ShuService struct {
	repo domain.ShuRepository
}

func NewShuService(repo domain.ShuRepository) *ShuService {
	return &ShuService{repo: repo}
}

func (s *ShuService) GetByYear(year int) ([]domain.ShuDistribution, error) {
	return s.repo.GetByYear(year)
}

func (s *ShuService) Calculate(year int) ([]domain.ShuDistribution, error) {
	// TODO: Algoritma kalkulasi pembagian SHU anggota per periode tahunan
	return s.repo.GetByYear(year)
}
