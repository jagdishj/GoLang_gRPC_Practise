package main

import (
	"context"
	"fmt"
	"greet/greet/greetpb"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	greetpb.GreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet Function is invoked with %v", req)
	firstname := req.GetGreeting().GetFirstName()
	result := "Hello " + firstname
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes Function is invoked with %v\n", req)
	firstname := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello" + firstname + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet Function is invoked \n")
	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//we have finished reading the client stream
			stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the client stream %v", err)
		}

		firstname := req.GetGreeting().GetFirstName()
		result += firstname + "! "
	}
	return nil
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone Function is invoked with streaming request\n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstname := req.GetGreeting().GetFirstName()
		result := "Hello " + firstname + "!"
		er := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if er != nil {
			log.Fatalf("Error while sending the data to client: %v", er)
			return er
		}
	}
}

func (*server) GreetwithDeadline(ctx context.Context, req *greetpb.GreetwithdeadlineRequest) (*greetpb.GreetwithdeadlineResponse, error) {
	fmt.Printf("GreetwithDeadline Function is invoked with %v \n", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			//the client cancelled the request
			fmt.Println("the client cancelled the request")
			return nil, status.Error(codes.Canceled, "the client cancelled the request")
		}
		time.Sleep(1 * time.Second)
	}

	firstname := req.GetGreeting().GetFirstName()
	result := "Hello " + firstname
	res := &greetpb.GreetwithdeadlineResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
