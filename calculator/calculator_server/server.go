package main

import (
	"context"
	"github.com/ksmt88/grpc-go-course-master/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (s *server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	numberA := request.GetNumberA()
	numberB := request.GetNumberB()
	result := numberA + numberB
	res := &calculatorpb.SumResponse{
		Res: result,
	}
	return res, nil
}

func main() {
	//fmt.Println("Hello World!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculateServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
