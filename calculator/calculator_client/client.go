package main

import (
	"context"
	"fmt"
	"github.com/ksmt88/grpc-go-course-master/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	//doClientStreaming(c)
	//doBiDiStreaming(c)
	doErrorCall(c, -1)
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

	numbers := []int32{6, 9, 7, 0}

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

func doBiDiStreaming(c calculatorpb.CalculateClient) {
	fmt.Println("Starting")

	numbers := []int32{1, 5, 3, 6, 2, 20}

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("error while calling FindMaximum: %v", err)
	}

	go func() {
		for _, number := range numbers {
			fmt.Printf("Sending message: %v\n", number)
			stream.Send(&calculatorpb.FindMaximumRequest{
				Number: number,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	waitc := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res)
		}
		close(waitc)
	}()

	<-waitc
}

func doErrorCall(c calculatorpb.CalculateClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{
		Number: n,
	})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Error while calling SquareRoot RPC: %v", err)
		}
	}
	log.Printf("Result: %v", res)
}
