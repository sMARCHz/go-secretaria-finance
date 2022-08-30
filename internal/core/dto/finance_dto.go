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

type BalanceResponse struct {
	AccountName string  `json:"account_name"`
	Balance     float64 `json:"balance"`
}

func (t BalanceResponse) ToProto() *pb.AccountBalance {
	response := pb.AccountBalance{
		AccountName: t.AccountName,
		Balance:     t.Balance,
	}
	return &response
}
