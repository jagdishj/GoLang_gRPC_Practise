package main

import (
	"context"
	"fmt"
	"greet/calculator/calcpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := calcpb.NewCalcServiceClient(cc)

	fmt.Printf("Created client: %f\n", c)

	doUnary(c)
}

func doUnary(c calcpb.CalcServiceClient) {

	fmt.Println("Starting to do a Sum Unary RPC...")

	req := &calcpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", res.SumResult)
}
