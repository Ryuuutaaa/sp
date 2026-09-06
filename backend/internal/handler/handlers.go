package handler

type Handlers struct {
	Member      *MemberHandler
	Savings     *SavingsHandler
	Loan        *LoanHandler
	Installment *InstallmentHandler
	Cash        *CashHandler
	Shu         *ShuHandler
	Setting     *SettingHandler
	User        *UserHandler
}
