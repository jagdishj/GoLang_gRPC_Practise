package main

import (
	"context"
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

	fmt.Printf("Created client: %f\n", c)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do a Unary RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Jagdish",
			LastName:  "J",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
