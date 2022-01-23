package main

import (
	"fmt"
	"greet/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
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
