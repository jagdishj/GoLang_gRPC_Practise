package main

import (
	"fmt"
	"greet/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, i'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	fmt.Printf("Created client: %f", c)
}
