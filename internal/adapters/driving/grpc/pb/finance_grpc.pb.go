// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: internal/adapters/driving/grpc/proto/finance.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FinanceServiceClient is the client API for FinanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FinanceServiceClient interface {
	Withdraw(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	Deposit(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	GetBalance(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Balance, error)
	GetOverviewMonthlyStatement(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OverviewStatement, error)
	GetOverviewAnnualStatement(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OverviewStatement, error)
}

type financeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFinanceServiceClient(cc grpc.ClientConnInterface) FinanceServiceClient {
	return &financeServiceClient{cc}
}

func (c *financeServiceClient) Withdraw(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/FinanceService/Withdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) Deposit(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/FinanceService/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, "/FinanceService/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) GetBalance(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Balance, error) {
	out := new(Balance)
	err := c.cc.Invoke(ctx, "/FinanceService/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) GetOverviewMonthlyStatement(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OverviewStatement, error) {
	out := new(OverviewStatement)
	err := c.cc.Invoke(ctx, "/FinanceService/GetOverviewMonthlyStatement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) GetOverviewAnnualStatement(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OverviewStatement, error) {
	out := new(OverviewStatement)
	err := c.cc.Invoke(ctx, "/FinanceService/GetOverviewAnnualStatement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinanceServiceServer is the server API for FinanceService service.
// All implementations must embed UnimplementedFinanceServiceServer
// for forward compatibility
type FinanceServiceServer interface {
	Withdraw(context.Context, *TransactionRequest) (*TransactionResponse, error)
	Deposit(context.Context, *TransactionRequest) (*TransactionResponse, error)
	Transfer(context.Context, *TransferRequest) (*TransferResponse, error)
	GetBalance(context.Context, *emptypb.Empty) (*Balance, error)
	GetOverviewMonthlyStatement(context.Context, *emptypb.Empty) (*OverviewStatement, error)
	GetOverviewAnnualStatement(context.Context, *emptypb.Empty) (*OverviewStatement, error)
	mustEmbedUnimplementedFinanceServiceServer()
}

// UnimplementedFinanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFinanceServiceServer struct {
}

func (UnimplementedFinanceServiceServer) Withdraw(context.Context, *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedFinanceServiceServer) Deposit(context.Context, *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedFinanceServiceServer) Transfer(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedFinanceServiceServer) GetBalance(context.Context, *emptypb.Empty) (*Balance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedFinanceServiceServer) GetOverviewMonthlyStatement(context.Context, *emptypb.Empty) (*OverviewStatement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOverviewMonthlyStatement not implemented")
}
func (UnimplementedFinanceServiceServer) GetOverviewAnnualStatement(context.Context, *emptypb.Empty) (*OverviewStatement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOverviewAnnualStatement not implemented")
}
func (UnimplementedFinanceServiceServer) mustEmbedUnimplementedFinanceServiceServer() {}

// UnsafeFinanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinanceServiceServer will
// result in compilation errors.
type UnsafeFinanceServiceServer interface {
	mustEmbedUnimplementedFinanceServiceServer()
}

func RegisterFinanceServiceServer(s grpc.ServiceRegistrar, srv FinanceServiceServer) {
	s.RegisterService(&FinanceService_ServiceDesc, srv)
}

func _FinanceService_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).Withdraw(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).Deposit(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).GetBalance(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_GetOverviewMonthlyStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).GetOverviewMonthlyStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/GetOverviewMonthlyStatement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).GetOverviewMonthlyStatement(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_GetOverviewAnnualStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).GetOverviewAnnualStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/GetOverviewAnnualStatement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).GetOverviewAnnualStatement(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// FinanceService_ServiceDesc is the grpc.ServiceDesc for FinanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FinanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FinanceService",
	HandlerType: (*FinanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Withdraw",
			Handler:    _FinanceService_Withdraw_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _FinanceService_Deposit_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _FinanceService_Transfer_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _FinanceService_GetBalance_Handler,
		},
		{
			MethodName: "GetOverviewMonthlyStatement",
			Handler:    _FinanceService_GetOverviewMonthlyStatement_Handler,
		},
		{
			MethodName: "GetOverviewAnnualStatement",
			Handler:    _FinanceService_GetOverviewAnnualStatement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/adapters/driving/grpc/proto/finance.proto",
}
