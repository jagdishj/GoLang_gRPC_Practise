package main

import (
	"context"
	"fmt"
	"greet/calculator/calcpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calcpb.CalcServiceServer
}

func (*server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Println("Received Sum RPC: %v", req)
	firstno := req.FirstNumber
	secno := req.SecondNumber
	sum := firstno + secno
	res := &calcpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calcpb.PrimaNumberDecoRequest, stream calcpb.CalcService_PrimeNumberDecompositionServer) error {
	fmt.Println("Received PrimeNumberDecomposition RPC: %v", req)
	number := req.GetNumber()
	divisor := int32(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calcpb.PrimeNumberDecoResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Println("Divisor has increased to %v", divisor)
		}
	}

	return nil
}

func main() {
	fmt.Println("Calculation Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to Connect: %v ", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
