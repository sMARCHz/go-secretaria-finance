package grpc

import (
	"context"

	"github.com/sMARCHz/go-secretaria-finance/internal/adapters/driving/grpc/pb"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/services"
	"google.golang.org/protobuf/types/known/emptypb"
)

type financeServiceServer struct {
	service services.FinanceService
	pb.UnimplementedFinanceServiceServer
}

func (f financeServiceServer) Withdraw(ctx context.Context, r *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	return nil, nil
}

func (f financeServiceServer) Deposit(ctx context.Context, r *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	return nil, nil
}

func (f financeServiceServer) Transfer(ctx context.Context, r *pb.TransferRequest) (*pb.TransferResponse, error) {
	return nil, nil
}

func (f financeServiceServer) GetBalance(ctx context.Context, r *emptypb.Empty) (*pb.Balance, error) {
	return nil, nil
}

func (f financeServiceServer) GetOverviewMonthlyStatement(ctx context.Context, r *emptypb.Empty) (*pb.OverviewStatement, error) {
	return nil, nil
}

func (f financeServiceServer) GetOverviewAnnualStatement(ctx context.Context, r *emptypb.Empty) (*pb.OverviewStatement, error) {
	return nil, nil
}
