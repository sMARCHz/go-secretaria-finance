package services

import "github.com/sMARCHz/go-secretaria-finance/internal/core/repository"

type FinanceService interface {
	Withdraw()
	Deposit()
	Transfer()
	GetBalance()
	GetOverviewMonthlyStatement()
	GetOverviewAnnualStatement()
}

type financeService struct {
	repository repository.FinanceRepository
}

func NewFinanceService(repo repository.FinanceRepository) financeService {
	return financeService{
		repository: repo,
	}
}

func (f financeService) Withdraw() {

}

func (f financeService) Deposit() {

}

func (f financeService) Transfer() {

}

func (f financeService) GetBalance() {

}

func (f financeService) GetOverviewMonthlyStatement() {

}

func (f financeService) GetOverviewAnnualStatement() {

}
