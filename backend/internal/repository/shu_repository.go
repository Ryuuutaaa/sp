package repository

import "sp-backend/internal/domain"

type ShuRepo struct {
	gql *GQLClient
}

func NewShuRepo(gql *GQLClient) *ShuRepo {
	return &ShuRepo{gql: gql}
}

func (r *ShuRepo) GetByYear(year int) ([]domain.ShuDistribution, error) {
	var resp struct {
		ShuDistributions []domain.ShuDistribution `json:"shuDistributions"`
	}
	err := r.gql.Run(`query($year: Int!) { shuDistributions(periodYear: $year) { id periodYear memberId savingsProportion loanProportion totalShu } }`, map[string]interface{}{"year": year}, &resp)
	return resp.ShuDistributions, err
}

func (r *ShuRepo) Create(input domain.CreateShuInput) (*domain.ShuDistribution, error) {
	var resp struct {
		CreateShuDistribution *domain.ShuDistribution `json:"createShuDistribution"`
	}
	err := r.gql.Run(`mutation($periodYear: Int!, $memberId: ID!, $savingsProportion: String!, $loanProportion: String!, $totalShu: String!) {
		createShuDistribution(periodYear: $periodYear, memberId: $memberId, savingsProportion: $savingsProportion, loanProportion: $loanProportion, totalShu: $totalShu) {
			id periodYear memberId totalShu
		}
	}`, map[string]interface{}{
		"periodYear":        input.PeriodYear,
		"memberId":          input.MemberID,
		"savingsProportion": input.SavingsProportion,
		"loanProportion":    input.LoanProportion,
		"totalShu":          input.TotalShu,
	}, &resp)
	return resp.CreateShuDistribution, err
}
