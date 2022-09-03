package db

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/repository"
	"github.com/sMARCHz/go-secretaria-finance/internal/logger"
)

type financeRepository struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewFinanceRepository(db *sqlx.DB, logger logger.Logger) repository.FinanceRepository {
	return &financeRepository{
		db:     db,
		logger: logger,
	}
}

func (f *financeRepository) GetAccountIDByName(name string) (*int, *errors.AppError) {
	var accountID int
	err := f.db.Get(&accountID, "SELECT account_id FROM accounts WHERE name = $1 LIMIT 1", name)
	if err != nil {
		if err == sql.ErrNoRows {
			f.logger.Errorf("account not found where name='%v'", name)
			return nil, errors.NotFoundError("account not found")
		}
		f.logger.Error("failed to get accountID: ", err)
		return nil, errors.InternalServerError("failed to get accountID")
	}
	return &accountID, nil
}

func (f *financeRepository) GetCategoryIDByAbbrNameAndTransactionType(categoryAbbrName string, transactionType string) (*int, *errors.AppError) {
	var categoryID int
	err := f.db.Get(&categoryID, "SELECT category_id FROM categories WHERE name_abbr = $1 AND transaction_type = $2 LIMIT 1", categoryAbbrName, transactionType)
	if err != nil {
		if err == sql.ErrNoRows {
			f.logger.Errorf("category not found where abbreviation='%v', transactionType='%v'", categoryAbbrName, transactionType)
			return nil, errors.NotFoundError("category not found")
		}
		f.logger.Error("failed to get categoryID: ", err)
		return nil, errors.InternalServerError("failed to get categoryID")
	}
	return &categoryID, nil
}

func (f *financeRepository) Withdraw(t domain.Transaction) (*domain.Account, *errors.AppError) {
	tx, err := f.db.Begin()
	if err != nil {
		f.logger.Error("failed to begin transaction: ", err)
		return nil, errors.InternalServerError("failed to begin transaction")
	}

	// Check if account's balance isn't less than withdrawal amount
	var balance float64
	if err := tx.QueryRow("SELECT balance FROM accounts WHERE account_id = $1", t.AccountID).Scan(&balance); err != nil {
		tx.Rollback()
		f.logger.Error("failed to get current balance: ", err)
		return nil, errors.InternalServerError("failed to get current balance")
	}
	if balance < -t.Amount {
		tx.Rollback()
		f.logger.Error("account's balance is less than withdrawal amount")
		return nil, errors.UnprocessableEntityServerError("balance can't be less than the withdrawal amount")
	}

	// Insert entry
	var entryID int
	query := "INSERT INTO entries(account_id, category_id, amount, description, created_at) VALUES($1, $2, $3, $4, $5) RETURNING entry_id"
	if err := tx.QueryRow(query, t.AccountID, t.CategoryID, t.Amount, t.Description, t.CreatedAt).Scan(&entryID); err != nil {
		tx.Rollback()
		f.logger.Error("failed to insert entries: ", err)
		return nil, errors.InternalServerError("failed to insert entries")
	}
	// Update the account's balance
	var account domain.Account
	err = tx.QueryRow("UPDATE accounts SET balance = balance + $1 WHERE account_id = $2 RETURNING name, balance, currency, created_at", t.Amount, t.AccountID).Scan(&account.Name, &account.Balance, &account.Currency, &account.CreatedAt)
	if err != nil {
		tx.Rollback()
		f.logger.Error("failed to update balance of account: ", err)
		return nil, errors.InternalServerError("failed to update balance of account")
	}

	// Commit
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		f.logger.Error("failed to commit transaction: ", err)
		return nil, errors.InternalServerError("failed to commit transaction")
	}
	return &account, nil
}

func (f *financeRepository) Deposit(t domain.Transaction) (*domain.Account, *errors.AppError) {
	tx, err := f.db.Begin()
	if err != nil {
		f.logger.Error("failed to begin transaction: ", err)
		return nil, errors.InternalServerError("failed to begin transaction")
	}

	// Insert entry
	var entryID int
	query := "INSERT INTO entries(account_id, category_id, amount, description, created_at) VALUES($1, $2, $3, $4, $5) RETURNING entry_id"
	if err := tx.QueryRow(query, t.AccountID, t.CategoryID, t.Amount, t.Description, t.CreatedAt).Scan(&entryID); err != nil {
		tx.Rollback()
		f.logger.Error("failed to insert entries: ", err)
		return nil, errors.InternalServerError("failed to insert entries")
	}
	// Update the account's balance
	var account domain.Account
	err = tx.QueryRow("UPDATE accounts SET balance = balance + $1 WHERE account_id = $2 RETURNING name, balance, currency, created_at", t.Amount, t.AccountID).Scan(&account.Name, &account.Balance, &account.Currency, &account.CreatedAt)
	if err != nil {
		tx.Rollback()
		f.logger.Error("failed to update balance of account: ", err)
		return nil, errors.InternalServerError("failed to update balance of account")
	}

	// Commit
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		f.logger.Error("failed to commit transaction: ", err)
		return nil, errors.InternalServerError("failed to commit transaction")
	}
	return &account, nil
}

func (f *financeRepository) Transfer(t domain.Transfer) (*domain.Account, *errors.AppError) {
	tx, err := f.db.Begin()
	if err != nil {
		f.logger.Error("failed to begin transaction: ", err)
		return nil, errors.InternalServerError("failed to begin transaction")
	}

	// Insert transfer
	if _, err := tx.Exec("INSERT INTO transfers(from_account_id, to_account_id, amount) VALUES($1, $2, $3)", t.FromAccountID, t.ToAccountID, t.Amount); err != nil {
		tx.Rollback()
		f.logger.Error("failed to insert transfers: ", err)
		return nil, errors.InternalServerError("failed to insert transfers")
	}

	// Get categoryID of TRANSFER
	var categoryID int
	if err := tx.QueryRow("SELECT category_id FROM categories WHERE name = 'transfer' AND transaction_type = 'TRANSFER' LIMIT 1").Scan(&categoryID); err != nil {
		tx.Rollback()
		f.logger.Error("failed to get categoryID: ", err)
		return nil, errors.InternalServerError("failed to get categoryID")
	}

	// Insert entries of fromAccount and toAccount
	if _, err := tx.Exec("INSERT INTO entries(account_id, category_id, amount, description) VALUES($1, $2, $3, $4)", t.FromAccountID, categoryID, -t.Amount, t.Description); err != nil {
		tx.Rollback()
		f.logger.Error("failed to insert entries: ", err)
		return nil, errors.InternalServerError("failed to insert entries")
	}
	if _, err := tx.Exec("INSERT INTO entries(account_id, category_id, amount, description) VALUES($1, $2, $3, $4)", t.ToAccountID, categoryID, t.Amount, t.Description); err != nil {
		tx.Rollback()
		f.logger.Error("failed to insert entries: ", err)
		return nil, errors.InternalServerError("failed to insert entries")
	}

	// Update balance of fromAccount and to Account
	var account domain.Account
	if err := tx.QueryRow("UPDATE accounts SET balance = balance + $1 WHERE account_id = $2 RETURNING name, balance", -t.Amount, t.FromAccountID).Scan(&account.Name, &account.Balance); err != nil {
		tx.Rollback()
		f.logger.Error("failed to update from_account balance: ", err)
		return nil, errors.InternalServerError("failed to update from_account balance")
	}
	if _, err := tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE account_id = $2", t.Amount, t.ToAccountID); err != nil {
		tx.Rollback()
		f.logger.Error("failed to update to_account balance: ", err)
		return nil, errors.InternalServerError("failed to update to_account balance")
	}

	// Commit
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		f.logger.Error("failed to commit transaction: ", err)
		return nil, errors.InternalServerError("failed to commit transaction")
	}
	return &account, nil
}

func (f *financeRepository) GetAllAccountBalance() ([]domain.Account, *errors.AppError) {
	var accounts []domain.Account
	if err := f.db.Select(&accounts, "SELECT name, balance FROM accounts"); err != nil {
		f.logger.Error("failed to query accounts: ", err)
		return nil, errors.InternalServerError("failed to get all balance of accounts")
	}
	return accounts, nil
}

func (f *financeRepository) GetEntryByDaterange(from time.Time, to time.Time) ([]domain.Entry, *errors.AppError) {
	var entries []domain.Entry
	query := `SELECT e.entry_id, e.account_id, e.category_id, c.name as category_name, e.amount, e.description, e.created_at 
	FROM entries e
	INNER JOIN categories c
	ON e.category_id = c.category_id 
	WHERE (e.created_at BETWEEN $1 AND $2)
	AND c."transaction_type" <> 'TRANSFER'
	`
	err := f.db.Select(&entries, query, from, to)
	if err != nil {
		f.logger.Error("failed to get entry by time range: ", err)
		return nil, errors.InternalServerError("failed to get entry by time range")
	}
	return entries, nil
}
