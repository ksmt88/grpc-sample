package main

import (
	"context"
	"fmt"
	"github.com/ksmt88/grpc-go-course-master/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello Client!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculateClient(cc)

	//doUnary(c)
	//doServerStreaming(c)
	doClientStreaming(c)
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

func doServerStreaming(c calculatorpb.CalculateClient) {
	fmt.Println("Starting")

	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 2500,
	}
	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from PrimeNumberDecomposition: %v", msg.GetRes())
	}

}

func doClientStreaming(c calculatorpb.CalculateClient) {
	fmt.Println("Starting")

	numbers := []int32{6,9,7,0}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("error while calling Average: %v", err)
	}

	for _, number := range numbers {
		fmt.Printf("Send request: %v\n", number)
		stream.Send(&calculatorpb.AverageRequest{
			Number: number,
		})
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving Average: %v", err)
	}

	fmt.Printf("result: %.1f\n", res.GetRes())
}
