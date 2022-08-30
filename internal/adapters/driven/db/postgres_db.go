package db

import (
	"database/sql"

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
	return financeRepository{
		db:     db,
		logger: logger,
	}
}

func (f financeRepository) GetAccountIDByName(name string) (int, *errors.AppError) {
	var accountID int
	err := f.db.Get(&accountID, "SELECT account_id FROM accounts WHERE name = $1 LIMIT 1", name)
	if err != nil {
		if err == sql.ErrNoRows {
			f.logger.Errorf("account not found where name='%v'", name)
			return -1, errors.NotFoundError("account not found")
		}
		f.logger.Error("failed to get accountID: ", err)
		return -1, errors.InternalServerError("failed to get accountID")
	}
	return accountID, nil
}

func (f financeRepository) GetCategoryIDByAbbrName(categoryAbbrName string) (int, *errors.AppError) {
	var categoryID int
	err := f.db.Get(&categoryID, "SELECT category_id FROM categories WHERE name_abbr = $1 LIMIT 1", categoryAbbrName)
	if err != nil {
		if err == sql.ErrNoRows {
			f.logger.Errorf("category not found where abbreviation='%v'", categoryAbbrName)
			return -1, errors.NotFoundError("category not found")
		}
		f.logger.Error("failed to get categoryID: ", err)
		return -1, errors.InternalServerError("failed to get categoryID")
	}
	return categoryID, nil
}

func (f financeRepository) Withdraw(t domain.Transaction) (domain.Account, *errors.AppError) {
	tx, err := f.db.Begin()
	if err != nil {
		f.logger.Error("failed to begin transaction: ", err)
		return domain.Account{}, errors.InternalServerError("failed to begin transaction")
	}

	// Check if account's balance isn't less than withdrawal amount
	var balance float64
	if err := tx.QueryRow("SELECT balance FROM accounts WHERE account_id = $1", t.AccountID).Scan(&balance); err != nil {
		tx.Rollback()
		f.logger.Error("failed to get current balance: ", err)
		return domain.Account{}, errors.InternalServerError("failed to get current balance")
	}
	if balance < -t.Amount {
		tx.Rollback()
		f.logger.Error("account's balance is less than withdrawal amount")
		return domain.Account{}, errors.UnprocessableEntityServerError("balance can't be less than the withdrawal amount")
	}

	// Insert entry
	var entryID int
	query := "INSERT INTO entries(account_id, category_id, amount, description, created_at) VALUES($1, $2, $3, $4, $5) RETURNING entry_id"
	if err := tx.QueryRow(query, t.AccountID, t.CategoryID, t.Amount, t.Description, t.CreatedAt).Scan(&entryID); err != nil {
		tx.Rollback()
		f.logger.Error("failed to insert entry: ", err)
		return domain.Account{}, errors.InternalServerError("failed to insert entry")
	}
	// Update the account's balance
	var account domain.Account
	err = tx.QueryRow("UPDATE accounts SET balance = balance + $1 WHERE account_id = $2 RETURNING name, balance, currency, created_at", t.Amount, t.AccountID).Scan(&account.Name, &account.Balance, &account.Currency, &account.CreatedAt)
	if err != nil {
		tx.Rollback()
		f.logger.Error("failed to update balance of account: ", err)
		return domain.Account{}, errors.InternalServerError("failed to update balance of account")
	}

	// Commit
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		f.logger.Error("failed to commit transaction: ", err)
		return domain.Account{}, errors.InternalServerError("failed to commit transaction")
	}
	return account, nil
}

func (f financeRepository) Deposit(t domain.Transaction) (domain.Account, *errors.AppError) {
	tx, err := f.db.Begin()
	if err != nil {
		f.logger.Error("failed to begin transaction: ", err)
		return domain.Account{}, errors.InternalServerError("failed to begin transaction")
	}

	// Insert entry
	var entryID int
	query := "INSERT INTO entries(account_id, category_id, amount, description, created_at) VALUES($1, $2, $3, $4, $5) RETURNING entry_id"
	if err := tx.QueryRow(query, t.AccountID, t.CategoryID, t.Amount, t.Description, t.CreatedAt).Scan(&entryID); err != nil {
		tx.Rollback()
		f.logger.Error("failed to insert entry: ", err)
		return domain.Account{}, errors.InternalServerError("failed to insert entry")
	}
	// Update the account's balance
	var account domain.Account
	err = tx.QueryRow("UPDATE accounts SET balance = balance + $1 WHERE account_id = $2 RETURNING name, balance, currency, created_at", t.Amount, t.AccountID).Scan(&account.Name, &account.Balance, &account.Currency, &account.CreatedAt)
	if err != nil {
		tx.Rollback()
		f.logger.Error("failed to update balance of account: ", err)
		return domain.Account{}, errors.InternalServerError("failed to update balance of account")
	}

	// Commit
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		f.logger.Error("failed to commit transaction: ", err)
		return domain.Account{}, errors.InternalServerError("failed to commit transaction")
	}
	return account, nil
}

func (f financeRepository) GetAllAccountBalance() ([]domain.Account, *errors.AppError) {
	var accounts []domain.Account
	if err := f.db.Select(&accounts, "SELECT name, balance FROM accounts"); err != nil {
		f.logger.Error("failed to query accounts: ", err)
		return nil, errors.InternalServerError("failed to get all balance of accounts")
	}
	return accounts, nil
}
