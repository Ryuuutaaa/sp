package service

import (
	"fmt"
	"sp-backend/internal/domain"
	"strconv"
	"time"
)

type InstallmentService struct {
	repo domain.InstallmentRepository
}

func NewInstallmentService(repo domain.InstallmentRepository) *InstallmentService {
	return &InstallmentService{repo: repo}
}

func (s *InstallmentService) GetByLoanID(loanID string) ([]domain.Installment, error) {
	return s.repo.GetByLoanID(loanID)
}

func (s *InstallmentService) GenerateSchedule(loan domain.Loan) ([]domain.Installment, error) {
	amount, _ := strconv.ParseFloat(loan.Amount, 64)
	rate, _ := strconv.ParseFloat(loan.InterestRate, 64)
	tenor := loan.TenorMonths

	// Bunga flat
	monthlyInterest := (amount * (rate / 100)) / float64(tenor)
	monthlyPrincipal := amount / float64(tenor)
	monthlyTotal := monthlyPrincipal + monthlyInterest

	var result []domain.Installment
	now := time.Now()

	for i := 1; i <= tenor; i++ {
		dueDate := now.AddDate(0, i, 0).Format("2006-01-02")
		input := domain.CreateInstallmentInput{
			LoanID:            loan.ID,
			InstallmentNumber: i,
			DueDate:           dueDate,
			PrincipalAmount:   fmt.Sprintf("%.2f", monthlyPrincipal),
			InterestAmount:    fmt.Sprintf("%.2f", monthlyInterest),
			TotalAmount:       fmt.Sprintf("%.2f", monthlyTotal),
		}
		item, err := s.repo.Create(input)
		if err != nil {
			return nil, err
		}
		result = append(result, *item)
	}
	return result, nil
}

func (s *InstallmentService) Pay(id string, input domain.PayInstallmentInput) (*domain.Installment, error) {
	return s.repo.Pay(id, input)
}
