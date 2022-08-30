package services

import (
	"database/sql"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/repository"
)

type FinanceService interface {
	Withdraw(dto.TransactionRequest) (dto.TransactionResponse, *errors.AppError)
	Deposit(dto.TransactionRequest) (dto.TransactionResponse, *errors.AppError)
	Transfer()
	GetBalance() ([]dto.BalanceResponse, *errors.AppError)
	GetOverviewMonthlyStatement()
	GetOverviewAnnualStatement()
}

type financeService struct {
	repository repository.FinanceRepository
}

func NewFinanceService(repo repository.FinanceRepository) FinanceService {
	return financeService{
		repository: repo,
	}
}

func (f financeService) Withdraw(req dto.TransactionRequest) (dto.TransactionResponse, *errors.AppError) {
	accountID, err := f.repository.GetAccountIDByName(req.AccountName)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	categoryID, err := f.repository.GetCategoryIDByAbbrName(req.Category)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	transaction := domain.Transaction{
		AccountID:   accountID,
		CategoryID:  categoryID,
		Description: sql.NullString{String: req.Description},
		Amount:      -req.Amount,
		CreatedAt:   req.CreatedAt,
	}
	account, err := f.repository.Withdraw(transaction)
	if err != nil {
		return dto.TransactionResponse{}, err
	}
	return account.ToTransactionResponseDto(), nil
}

func (f financeService) Deposit(req dto.TransactionRequest) (dto.TransactionResponse, *errors.AppError) {
	accountID, err := f.repository.GetAccountIDByName(req.AccountName)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	categoryID, err := f.repository.GetCategoryIDByAbbrName(req.Category)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	transaction := domain.Transaction{
		AccountID:   accountID,
		CategoryID:  categoryID,
		Description: sql.NullString{String: req.Description},
		Amount:      req.Amount,
		CreatedAt:   req.CreatedAt,
	}
	account, err := f.repository.Deposit(transaction)
	if err != nil {
		return dto.TransactionResponse{}, err
	}
	return account.ToTransactionResponseDto(), nil
}

func (f financeService) Transfer() {

}

func (f financeService) GetBalance() ([]dto.BalanceResponse, *errors.AppError) {
	accounts, err := f.repository.GetAllAccountBalance()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.BalanceResponse, len(accounts))
	for i, v := range accounts {
		responses[i] = v.ToBalanceResponseDto()
	}
	return responses, nil
}

func (f financeService) GetOverviewMonthlyStatement() {

}

func (f financeService) GetOverviewAnnualStatement() {

}
