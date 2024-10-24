package main

import (
	"context"
	"fmt"
	"github.com/thanhdn15/concrete_lean_go/gRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"math"
	"net"
	"time"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("Sum called...")
	resp := &calculatorpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return resp, nil
}

func (s *server) SumWithDeadline(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("Sum with deadline...")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			log.Println("context.Canceled...")
			return nil, status.Error(codes.Canceled, "client canceled req")
		}

		time.Sleep(1 * time.Second)
	}

	resp := &calculatorpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return resp, nil
}

func (s *server) PrimeNumberDecomposition(req *calculatorpb.PNDRequest, stream grpc.ServerStreamingServer[calculatorpb.PNDResponse]) error {
	log.Println("PrimeNumberDecomposition called...")
	k := int32(2)
	N := req.GetNumber()
	for N > 1 {
		if N%k == 0 {
			N = N / k
			// sent to client
			stream.Send(&calculatorpb.PNDResponse{
				Result: k,
			})
		} else {
			k++
			log.Printf("k increase to %v", k)
		}
	}
	return nil
}

func (s *server) Average(stream grpc.ClientStreamingServer[calculatorpb.AverageRequest, calculatorpb.AverageResponse]) error {
	log.Println("Average called...")

	var total float32
	var count int
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			resp := calculatorpb.AverageResponse{
				Result: total / float32(count),
			}

			return stream.SendAndClose(&resp)
		}

		if err != nil {
			log.Fatalf("err whild Recv Average %v", err)

			return err
		}

		total += req.GetNum()
		count++
	}
}

func (s *server) Max(stream grpc.BidiStreamingServer[calculatorpb.MaxRequest, calculatorpb.MaxResponse]) error {
	log.Println("Find max called ...")

	max := int32(0)
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Println("EOF...")
			return nil
		}

		if err != nil {
			log.Fatalf("err whild Recv FindMax %v", err)

			return err
		}

		num := req.GetNum()
		log.Printf("recv num %v\n", num)
		if num > max {
			max = num
		}

		err = stream.Send(&calculatorpb.MaxResponse{
			Result: max,
		})

		if err != nil {
			log.Fatalf("send max err %v", err)

			return err
		}
	}
}

func (s *server) Square(ctx context.Context, req *calculatorpb.SquareRequest) (*calculatorpb.SquareResponse, error) {
	log.Println("Square call...")
	num := req.GetNum()

	if num < 0 {
		log.Printf("request num < 0, num=%v", num)
		return nil, status.Errorf(codes.InvalidArgument, "Expect num > 0, req num was %v", num)
	}

	return &calculatorpb.SquareResponse{
		SquareRoot: math.Sqrt(float64(num)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50069")

	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	certFile := "calculator/ssl/server.crt"
	keyFile := "calculator/ssl/server.pem"

	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)

	if sslErr != nil {
		log.Fatalf("failed to create credentials: %v", sslErr)
	}

	s := grpc.NewServer(grpc.Creds(creds))

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("Calculator server is running...!")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while server %v", err)
	}
}
