package service

import (
	"errors"
	"sp-backend/internal/domain"
	"strconv"
)

type SavingsService struct {
	repo domain.SavingsRepository
}

func NewSavingsService(repo domain.SavingsRepository) *SavingsService {
	return &SavingsService{repo: repo}
}

func (s *SavingsService) GetTypes() ([]domain.SavingsType, error) {
	return s.repo.GetTypes()
}

func (s *SavingsService) GetTransactions(memberID string) ([]domain.SavingsTransaction, error) {
	return s.repo.GetTransactions(memberID)
}

func (s *SavingsService) Deposit(input domain.CreateSavingsInput) (*domain.SavingsTransaction, error) {
	input.Type = "setor"
	return s.repo.CreateTransaction(input)
}

func (s *SavingsService) Withdraw(input domain.CreateSavingsInput) (*domain.SavingsTransaction, error) {
	amt, _ := strconv.ParseFloat(input.Amount, 64)
	if amt <= 0 {
		return nil, errors.New("jumlah penarikan tidak valid")
	}
	input.Type = "tarik"
	return s.repo.CreateTransaction(input)
}
