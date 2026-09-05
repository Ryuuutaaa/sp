package service

import (
	"errors"
	"sp-backend/internal/domain"
	"strconv"
)

type SavingsService struct {
	repo domain.SavingsRepository
}

func NewSavingRepository(repo domain.SavingsRepository) *SavingsService {
	return &SavingsService{repo: repo}
}

func (s *SavingsService) GetTypes() ([]domain.SavingsType, error) {
	return s.repo.GetTypes()
}

func (s *SavingsService) GetTransaction(memberID string) ([]domain.SavingsTransaction, error) {
	return s.repo.GetTransactions(memberID)
}

func (s *SavingsService) Deposit(input domain.CreateSavingsInput) (*domain.SavingsTransaction, error) {
	// TODO: ambil saldo terakhir, tambahkan amount,
	input.Type = "setor"
	return s.repo.CreateTransaction(input)
}

func (s *SavingsService) Withdraw(input domain.CreateSavingsInput) (*domain.SavingsTransaction, error) {
	// TODO: Validasi hanya simpanan sukarela, validasi saldo cukup
	amt, _ := strconv.ParseFloat(input.Amount, 64)
	if amt <= 0 {
		return nil, errors.New("jumlah penarikan tidak valid")
	}
	input.Type = "tarik"
	return s.repo.CreateTransaction(input)
}
