package grpc

import (
	"context"

	"github.com/sMARCHz/go-secretaria-finance/internal/adapters/driving/grpc/pb"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/services"
	"github.com/sMARCHz/go-secretaria-finance/internal/logger"
	"github.com/sMARCHz/go-secretaria-finance/internal/utils"
	"google.golang.org/protobuf/types/known/emptypb"
)

type financeServiceServer struct {
	service services.FinanceService
	logger  logger.Logger
	pb.UnimplementedFinanceServiceServer
}

func newFinanceServiceServer(service services.FinanceService, logger logger.Logger) financeServiceServer {
	return financeServiceServer{
		service: service,
		logger:  logger,
	}
}

func (f financeServiceServer) Withdraw(ctx context.Context, r *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	req := dto.TransactionRequest{
		AccountName: r.AccountName,
		Category:    r.Category,
		Description: r.Description,
		Amount:      r.Amount,
		CreatedAt:   r.Timestamp.AsTime(),
	}
	response, err := f.service.Withdraw(req)
	pbResponse := response.ToProto()
	if err != nil {
		pbResponse.Status = int32(err.StatusCode)
		pbResponse.Error = err.Message
		return pbResponse, utils.ConvertHttpErrToGRPC(err, f.logger)
	}
	return pbResponse, nil
}

func (f financeServiceServer) Deposit(ctx context.Context, r *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	req := dto.TransactionRequest{
		AccountName: r.AccountName,
		Category:    r.Category,
		Description: r.Description,
		Amount:      r.Amount,
		CreatedAt:   r.Timestamp.AsTime(),
	}
	response, err := f.service.Deposit(req)
	pbResponse := response.ToProto()
	if err != nil {
		pbResponse.Status = int32(err.StatusCode)
		pbResponse.Error = err.Message
		return pbResponse, utils.ConvertHttpErrToGRPC(err, f.logger)
	}
	return pbResponse, nil
}

func (f financeServiceServer) Transfer(ctx context.Context, r *pb.TransferRequest) (*pb.TransferResponse, error) {
	req := dto.TransferRequest{
		FromAccountName: r.FromAccountName,
		ToAccountName:   r.ToAccountName,
		Description:     r.Description,
		Amount:          r.Amount,
	}
	response, err := f.service.Transfer(req)
	pbResponse := response.ToProto()
	if err != nil {
		pbResponse.Status = int32(err.StatusCode)
		pbResponse.Error = err.Message
		return pbResponse, utils.ConvertHttpErrToGRPC(err, f.logger)
	}
	return pbResponse, nil
}

func (f financeServiceServer) GetBalance(ctx context.Context, r *emptypb.Empty) (*pb.GetBalanceResponse, error) {
	response, err := f.service.GetBalance()

	accountsBalance := make([]*pb.AccountBalance, len(response))
	for i, v := range response {
		accountsBalance[i] = v.ToProto()
	}

	pbResponse := &pb.GetBalanceResponse{Accounts: accountsBalance}
	if err != nil {
		pbResponse.Status = int32(err.StatusCode)
		pbResponse.Error = err.Message
		return pbResponse, utils.ConvertHttpErrToGRPC(err, f.logger)
	}
	return pbResponse, nil
}

func (f financeServiceServer) GetOverviewStatement(ctx context.Context, r *pb.OverviewStatmentRequest) (*pb.OverviewStatementResponse, error) {
	req := dto.GetOverviewStatementRequest{
		From: r.From.AsTime(),
		To:   r.To.AsTime(),
	}
	response, err := f.service.GetOverviewStatement(req)
	pbResponse := response.ToProto()
	if err != nil {
		pbResponse.Status = int32(err.StatusCode)
		pbResponse.Error = err.Message
		return pbResponse, utils.ConvertHttpErrToGRPC(err, f.logger)
	}
	return pbResponse, nil
}

func (f financeServiceServer) GetOverviewMonthlyStatement(ctx context.Context, r *emptypb.Empty) (*pb.OverviewStatementResponse, error) {
	response, err := f.service.GetOverviewMonthlyStatement()
	pbResponse := response.ToProto()
	if err != nil {
		pbResponse.Status = int32(err.StatusCode)
		pbResponse.Error = err.Message
		return pbResponse, utils.ConvertHttpErrToGRPC(err, f.logger)
	}
	return pbResponse, nil
}

func (f financeServiceServer) GetOverviewAnnualStatement(ctx context.Context, r *emptypb.Empty) (*pb.OverviewStatementResponse, error) {
	response, err := f.service.GetOverviewAnnualStatement()
	pbResponse := response.ToProto()
	if err != nil {
		pbResponse.Status = int32(err.StatusCode)
		pbResponse.Error = err.Message
		return pbResponse, utils.ConvertHttpErrToGRPC(err, f.logger)
	}
	return pbResponse, nil
}
