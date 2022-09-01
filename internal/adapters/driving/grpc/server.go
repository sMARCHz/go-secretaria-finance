package grpc

import (
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/sMARCHz/go-secretaria-finance/internal/adapters/driven/db"
	"github.com/sMARCHz/go-secretaria-finance/internal/adapters/driving/grpc/pb"
	"github.com/sMARCHz/go-secretaria-finance/internal/config"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/services"
	"github.com/sMARCHz/go-secretaria-finance/internal/logger"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	server   *grpc.Server
	config   config.AppConfiguration
	logger   logger.Logger
	database *sqlx.DB
}

func NewGRPCServer(config config.AppConfiguration, logger logger.Logger, db *sqlx.DB) (GRPCServer, func()) {
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	return GRPCServer{
		server:   grpcServer,
		config:   config,
		logger:   logger,
		database: db,
	}, grpcServer.GracefulStop
}

func (g GRPCServer) Start() {
	// wiring
	fRepo := db.NewFinanceRepository(g.database, g.logger)
	fService := services.NewFinanceService(fRepo, g.logger)
	fServer := newFinanceServiceServer(fService, g.logger)

	grpcServer := g.server
	logger := g.logger
	address := g.config.Address
	port := g.config.Port
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", address, port))
	if err != nil {
		logger.Fatal("failed to listen: ", err)
	}

	// Register service to grpc server
	pb.RegisterFinanceServiceServer(grpcServer, fServer)

	// Start server
	logger.Infof("Starting gRPC server at %v:%v...", address, port)
	grpcServer.Serve(lis)
}
