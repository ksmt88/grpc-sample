package main

import (
	"context"
	"github.com/ksmt88/grpc-go-course-master/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type server struct{}

func (s *server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	numberA := request.GetNumberA()
	numberB := request.GetNumberB()
	result := numberA + numberB
	res := &calculatorpb.SumResponse{
		Res: result,
	}
	return res, nil
}

func (s *server) PrimeNumberDecomposition(request *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.Calculate_PrimeNumberDecompositionServer) error {
	number := request.GetNumber()
	var divisor int32 = 2
	for {
		if number%divisor == int32(0) {
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				Res: divisor,
			}
			number = number / divisor
			if err := stream.Send(res); err != nil {
				return err
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			divisor = divisor + 1
		}

		if divisor > number {
			break
		}
	}
	return nil
}

func (s *server) Average(stream calculatorpb.Calculate_AverageServer) error {
	result := float32(0)
	count := float32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.AverageResponse{
				Res: result / count,
			})
		}

		if err != nil {
			log.Fatalf("Error while streaming: %v", err)
		}

		result += float32(req.GetNumber())
		count++
	}
}

func (s *server) FindMaximum(stream calculatorpb.Calculate_FindMaximumServer) error {
	current := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while streaming: %v", err)
		}

		if current < req.GetNumber() {
			current = req.GetNumber()

			sendErr := stream.Send(&calculatorpb.FindMaximumResponse{
				Res: current,
			})
			if sendErr != nil {
				log.Fatalf("Error sending data: %v", err)
				return err
			}
		}
	}
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
