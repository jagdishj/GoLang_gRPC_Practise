package main

import (
	"context"
	"fmt"
	"greet/calculator/calcpb"
	"io"
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
func (*server) ComputeAverage(stream calcpb.CalcService_ComputeAverageServer) error {
	fmt.Println("Received ComputeAverage RPC\n")

	sum := 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			average := float64(sum) / float64(count)

			stream.SendAndClose(&calcpb.ComputeAverageResponse{
				Average: average,
			})
			//fmt.Printf("Average=%v\n", average)
			break
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		sum += int(req.GetNumber())
		count++

		fmt.Printf("values of sum = %v and count = %v \n", sum, count)
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
