package service

import "sp-backend/internal/domain"

type SettingService struct {
	repo domain.SettingRepository
}

func NewSettingService(repo domain.SettingRepository) *SettingService {
	return &SettingService{repo: repo}
}

func (s *SettingService) GetAll() ([]domain.Setting, error) {
	return s.repo.GetAll()
}

func (s *SettingService) GetByKey(key string) (*domain.Setting, error) {
	return s.repo.GetByKey(key)
}

func (s *SettingService) Upsert(input domain.UpsertSettingInput) (*domain.Setting, error) {
	return s.repo.Upsert(input)
}
