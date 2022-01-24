package main

import (
	"context"
	"fmt"
	"greet/calculator/calcpb"
	"io"
	"log"
	"time"

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

	//doUnary(c)

	//doServerStreaming(c)

	//doClientStreaming(c)

	doServerClientStreaming(c)
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

func doServerStreaming(c calcpb.CalcServiceClient) {

	fmt.Println("Starting to do a ServerStreaming RPC...")

	req := &calcpb.PrimaNumberDecoRequest{
		Number: 123903928,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling ServerStreaming RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}

}

func doClientStreaming(c calcpb.CalcServiceClient) {
	fmt.Println("Starting to do a ClientStreaming RPC...")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		fmt.Printf("Sending number: %v \n", number)
		stream.Send(&calcpb.ComputeAverageRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving the response %v", err)
	}

	fmt.Printf("THe Average is: %v ", res.GetAverage())
}

func doServerClientStreaming(c calcpb.CalcServiceClient) {
	fmt.Println("Starting to do a ServerClientStreaming(FindMaximum) RPC...")

	steam, err := c.FindMaximum(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream and calling findMaximum: %v", err)
	}

	waitc := make(chan struct{})
	//Send go routine
	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			SendErr := steam.Send(&calcpb.FindMaximumRequest{
				Number: number,
			})
			fmt.Printf("Sending number %v\n", number)
			time.Sleep(1000 * time.Millisecond)
			if SendErr != nil {
				log.Fatalf("Error in Sending the request to Server: %v", SendErr)
			}
		}
		steam.CloseSend()
	}()
	//Receive go routing
	go func() {
		for {
			res, err := steam.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Problem while reading server stream: %v", err)
			}
			maximum := res.GetMaximum()
			fmt.Printf("Received a new maximum of ...%v\n", maximum)
		}
		close(waitc)
	}()

	<-waitc
}
