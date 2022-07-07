package repository

import (
	"github.com/jmoiron/sqlx"
)

type FinanceRepository interface {
}

type financeRepository struct {
	db *sqlx.DB
}

func NewFinanceRepository(db *sqlx.DB) financeRepository {
	return financeRepository{
		db: db,
	}
}
