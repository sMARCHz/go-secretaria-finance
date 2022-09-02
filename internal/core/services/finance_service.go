package services

import (
	"database/sql"
	"time"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/repository"
	"github.com/sMARCHz/go-secretaria-finance/internal/logger"
)

type FinanceService interface {
	Withdraw(dto.TransactionRequest) (dto.TransactionResponse, *errors.AppError)
	Deposit(dto.TransactionRequest) (dto.TransactionResponse, *errors.AppError)
	Transfer(dto.TransferRequest) (dto.TransferResponse, *errors.AppError)
	GetBalance() ([]dto.BalanceResponse, *errors.AppError)
	GetOverviewStatement(dto.GetOverviewStatementRequest) (dto.GetOverviewStatementResponse, *errors.AppError)
	GetOverviewMonthlyStatement() (dto.GetOverviewStatementResponse, *errors.AppError)
	GetOverviewAnnualStatement() (dto.GetOverviewStatementResponse, *errors.AppError)
}

type financeService struct {
	repository repository.FinanceRepository
	logger     logger.Logger
}

func NewFinanceService(repo repository.FinanceRepository, logger logger.Logger) FinanceService {
	return financeService{
		repository: repo,
		logger:     logger,
	}
}

func (f financeService) Withdraw(req dto.TransactionRequest) (dto.TransactionResponse, *errors.AppError) {
	categoryID, err := f.repository.GetCategoryIDByAbbrNameAndTransactionType(req.Category, "WITHDRAW")
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	accountID, err := f.repository.GetAccountIDByName(req.AccountName)
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
	categoryID, err := f.repository.GetCategoryIDByAbbrNameAndTransactionType(req.Category, "DEPOSIT")
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	accountID, err := f.repository.GetAccountIDByName(req.AccountName)
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

func (f financeService) Transfer(req dto.TransferRequest) (dto.TransferResponse, *errors.AppError) {
	fromAccountID, err := f.repository.GetAccountIDByName(req.FromAccountName)
	if err != nil {
		return dto.TransferResponse{}, err
	}
	toAccountID, err := f.repository.GetAccountIDByName(req.ToAccountName)
	if err != nil {
		return dto.TransferResponse{}, err
	}

	transfer := domain.Transfer{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Description:   sql.NullString{String: req.Description},
		Amount:        req.Amount,
	}
	fromAccount, err := f.repository.Transfer(transfer)
	if err != nil {
		return dto.TransferResponse{}, err
	}
	return fromAccount.ToTransferResponseDto(), nil
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

func (f financeService) GetOverviewStatement(req dto.GetOverviewStatementRequest) (dto.GetOverviewStatementResponse, *errors.AppError) {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		f.logger.Error("failed to load time location")
		return dto.GetOverviewStatementResponse{}, errors.InternalServerError("failed to load time location")
	}
	from := time.Date(req.From.Year(), req.From.Month(), req.From.Day(), 0, 0, 0, 0, loc)
	to := time.Date(req.To.Year(), req.To.Month(), req.To.Day(), 23, 59, 59, 0, loc)

	entries, appErr := f.repository.GetEntryByDaterange(from, to)
	if appErr != nil {
		return dto.GetOverviewStatementResponse{}, appErr
	}

	// calculate profit and split the entries 2 groups(revenue,expense)
	profit := 0.0
	totalRevenue := 0.0
	totalExpense := 0.0
	statement := map[string][]domain.Entry{
		"revenue": {},
		"expense": {},
	}
	for _, v := range entries {
		profit += v.Amount
		if v.Amount > 0 {
			totalRevenue += v.Amount
			statement["revenue"] = append(statement["revenue"], v)
		} else {
			totalExpense += v.Amount
			statement["expense"] = append(statement["expense"], v)
		}
	}

	// group entries by category
	revenue := dto.OverviewStatementSection{
		Total: totalRevenue,
	}
	expense := dto.OverviewStatementSection{
		Total: totalExpense,
	}
	for k, entries := range statement {
		categorizedEntry := f.groupEntriesByCategory(entries)
		if k == "revenue" {
			revenue.Entries = categorizedEntry
		} else if k == "expense" {
			expense.Entries = categorizedEntry
		}
	}
	response := dto.GetOverviewStatementResponse{
		Profit:  profit,
		Revenue: revenue,
		Expense: expense,
	}
	return response, nil
}

func (financeService) groupEntriesByCategory(entries []domain.Entry) []dto.CategorizedEntry {
	m := make(map[string]dto.CategorizedEntry)
	for _, e := range entries {
		categorizedEntry, present := m[e.CategoryName]
		if present {
			categorizedEntry.Amount += e.Amount
			m[e.CategoryName] = categorizedEntry
		} else {
			m[e.CategoryName] = dto.CategorizedEntry{Category: e.CategoryName, Amount: e.Amount}
		}
	}
	categorizedEntries := make([]dto.CategorizedEntry, 0)
	for _, v := range m {
		categorizedEntries = append(categorizedEntries, v)
	}
	return categorizedEntries
}

func (f financeService) GetOverviewMonthlyStatement() (dto.GetOverviewStatementResponse, *errors.AppError) {
	today := time.Now()
	from := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, time.UTC)
	to := from.AddDate(0, 1, -1)
	req := dto.GetOverviewStatementRequest{
		From: from,
		To:   to,
	}
	return f.GetOverviewStatement(req)
}

func (f financeService) GetOverviewAnnualStatement() (dto.GetOverviewStatementResponse, *errors.AppError) {
	today := time.Now()
	from := time.Date(today.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(today.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	req := dto.GetOverviewStatementRequest{
		From: from,
		To:   to,
	}
	return f.GetOverviewStatement(req)
}
