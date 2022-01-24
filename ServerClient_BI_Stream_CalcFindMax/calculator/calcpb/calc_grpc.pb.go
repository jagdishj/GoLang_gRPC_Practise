// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package calcpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNumberDecomposition(ctx context.Context, in *PrimaNumberDecoRequest, opts ...grpc.CallOption) (CalcService_PrimeNumberDecompositionClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalcService_ComputeAverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalcService_FindMaximumClient, error)
}

type calcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcServiceClient(cc grpc.ClientConnInterface) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calc.CalcService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimaNumberDecoRequest, opts ...grpc.CallOption) (CalcService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcService_ServiceDesc.Streams[0], "/calc.CalcService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalcService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberDecoResponse, error)
	grpc.ClientStream
}

type calcServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calcServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberDecoResponse, error) {
	m := new(PrimeNumberDecoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalcService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcService_ServiceDesc.Streams[1], "/calc.CalcService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceComputeAverageClient{stream}
	return x, nil
}

type CalcService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calcServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calcServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalcService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcService_ServiceDesc.Streams[2], "/calc.CalcService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceFindMaximumClient{stream}
	return x, nil
}

type CalcService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calcServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calcServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcServiceServer is the server API for CalcService service.
// All implementations must embed UnimplementedCalcServiceServer
// for forward compatibility
type CalcServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNumberDecomposition(*PrimaNumberDecoRequest, CalcService_PrimeNumberDecompositionServer) error
	ComputeAverage(CalcService_ComputeAverageServer) error
	FindMaximum(CalcService_FindMaximumServer) error
	mustEmbedUnimplementedCalcServiceServer()
}

// UnimplementedCalcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalcServiceServer struct {
}

func (UnimplementedCalcServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalcServiceServer) PrimeNumberDecomposition(*PrimaNumberDecoRequest, CalcService_PrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecomposition not implemented")
}
func (UnimplementedCalcServiceServer) ComputeAverage(CalcService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (UnimplementedCalcServiceServer) FindMaximum(CalcService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}
func (UnimplementedCalcServiceServer) mustEmbedUnimplementedCalcServiceServer() {}

// UnsafeCalcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcServiceServer will
// result in compilation errors.
type UnsafeCalcServiceServer interface {
	mustEmbedUnimplementedCalcServiceServer()
}

func RegisterCalcServiceServer(s grpc.ServiceRegistrar, srv CalcServiceServer) {
	s.RegisterService(&CalcService_ServiceDesc, srv)
}

func _CalcService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalcService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimaNumberDecoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcServiceServer).PrimeNumberDecomposition(m, &calcServicePrimeNumberDecompositionServer{stream})
}

type CalcService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberDecoResponse) error
	grpc.ServerStream
}

type calcServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calcServicePrimeNumberDecompositionServer) Send(m *PrimeNumberDecoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalcService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).ComputeAverage(&calcServiceComputeAverageServer{stream})
}

type CalcService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calcServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calcServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalcService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).FindMaximum(&calcServiceFindMaximumServer{stream})
}

type CalcService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calcServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calcServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcService_ServiceDesc is the grpc.ServiceDesc for CalcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calc.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalcService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalcService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalcService_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalcService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calcpb/calc.proto",
}
