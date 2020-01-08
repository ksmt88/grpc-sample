package main

import (
	"context"
	"fmt"
	"github.com/ksmt88/grpc-go-course-master/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello Client!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculateClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculateClient) {
	req := &calculatorpb.SumRequest{
		NumberA: 3,
		NumberB: 10,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculate RPC: %v", err)
	}
	log.Printf("Response from Calculate: %v", res)
}
