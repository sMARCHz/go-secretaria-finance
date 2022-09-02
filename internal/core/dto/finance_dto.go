package dto

import (
	"time"

	"github.com/sMARCHz/go-secretaria-finance/internal/adapters/driving/grpc/pb"
)

type TransactionRequest struct {
	AccountName string    `json:"account_name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"-"`
}

type TransactionResponse struct {
	AccountName string  `json:"account_name"`
	Balance     float64 `json:"balance"`
}

func (t TransactionResponse) ToProto() *pb.TransactionResponse {
	response := pb.TransactionResponse{
		AccountName: t.AccountName,
		Balance:     t.Balance,
	}
	return &response
}

type TransferRequest struct {
	FromAccountName string  `json:"from_account"`
	ToAccountName   string  `json:"to_account"`
	Description     string  `json:"description"`
	Amount          float64 `json:"amount"`
}

type TransferResponse struct {
	FromAccountName    string  `json:"from_account"`
	FromAccountBalance float64 `json:"balance"`
}

func (t TransferResponse) ToProto() *pb.TransferResponse {
	response := pb.TransferResponse{
		FromAccountName: t.FromAccountName,
		Balance:         t.FromAccountBalance,
	}
	return &response
}

type BalanceResponse struct {
	AccountName string  `json:"account_name"`
	Balance     float64 `json:"balance"`
}

func (b BalanceResponse) ToProto() *pb.AccountBalance {
	response := pb.AccountBalance{
		AccountName: b.AccountName,
		Balance:     b.Balance,
	}
	return &response
}

type GetOverviewStatementRequest struct {
	From time.Time
	To   time.Time
}

type GetOverviewStatementResponse struct {
	Revenue OverviewStatementSection
	Expense OverviewStatementSection
	Profit  float64
}

func (g GetOverviewStatementResponse) ToProto() *pb.OverviewStatementResponse {
	return &pb.OverviewStatementResponse{
		Revenue: g.Revenue.ToProto(),
		Expense: g.Expense.ToProto(),
		Profit:  g.Profit,
	}
}

type OverviewStatementSection struct {
	Total   float64
	Entries []CategorizedEntry
}

func (o OverviewStatementSection) ToProto() *pb.OverviewStatmentSection {
	entries := make([]*pb.CategorizedEntry, len(o.Entries))
	for i, v := range o.Entries {
		entries[i] = v.ToProto()
	}
	return &pb.OverviewStatmentSection{
		Total:   o.Total,
		Entries: entries,
	}
}

type CategorizedEntry struct {
	Category string
	Amount   float64
}

func (c CategorizedEntry) ToProto() *pb.CategorizedEntry {
	return &pb.CategorizedEntry{
		Category: c.Category,
		Amount:   c.Amount,
	}
}
