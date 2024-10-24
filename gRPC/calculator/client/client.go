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
	"time"
)

func main() {

	certFile := "calculator/ssl/ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")

	if sslErr != nil {
		log.Fatalf("failed to create credentials: %v", sslErr)
	}

	cc, err := grpc.Dial("localhost:50069", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("Error while dial %v", err)
	}

	defer cc.Close()

	client := calculatorpb.NewCalculatorServiceClient(cc)

	//log.Printf("service client %f", client)
	callSum(client)
	callPND(client)
	callAverage(client)
	callMax(client)
	callSqrt(client, -4)
	callSumWithDeadline(client, 1*time.Second)
	callSumWithDeadline(client, 5*time.Second)
}

func callSum(c calculatorpb.CalculatorServiceClient) {
	log.Println("Calling sum api")

	resp, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 5,
		Num2: 10,
	})

	if err != nil {
		log.Fatalf("Call sum api error: %v", err)
	}

	fmt.Printf("sum api response %v\n", resp.GetResult())
}

func callSumWithDeadline(c calculatorpb.CalculatorServiceClient, timeout time.Duration) {
	log.Println("Calling sum with deadline api...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	resp, err := c.SumWithDeadline(ctx, &calculatorpb.SumRequest{
		Num1: 5,
		Num2: 10,
	})

	if err != nil {
		if statusErr, ok := status.FromError(err); ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("calling sum with deadline DeadlineExceeded")
			} else {
				log.Printf("calling with deadline api err %v", err)
			}
		} else {
			log.Printf("Calling with deadline unknown err %v", err)
		}

		return
	}

	log.Printf("Sum with deadling api response %v\n", resp)
}

func callPND(c calculatorpb.CalculatorServiceClient) {
	log.Println("PrimeNumberDecomposition sum")

	stream, err := c.PrimeNumberDecomposition(context.Background(), &calculatorpb.PNDRequest{
		Number: 120,
	})

	if err != nil {
		log.Fatalf("Call PrimeNumberDecomposition api is error: %f", err)
	}

	for {
		resp, errRcv := stream.Recv()

		if errRcv == io.EOF {
			log.Println("server finish streaming")
			return
		}

		fmt.Printf("PrimeNumberDecomposition api response %v\n", resp.GetResult())
	}
}

func callAverage(c calculatorpb.CalculatorServiceClient) {
	log.Println("Average sum")

	listReqs := []calculatorpb.AverageRequest{
		calculatorpb.AverageRequest{
			Num: 5,
		},
		calculatorpb.AverageRequest{
			Num: 10,
		},
		calculatorpb.AverageRequest{
			Num: 15,
		},
		calculatorpb.AverageRequest{
			Num: 3,
		},
		calculatorpb.AverageRequest{
			Num: 5.2,
		},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Call Average api is error: %f", err)
	}

	for _, req := range listReqs {
		err = stream.Send(&req)

		if err != nil {
			log.Fatalf("send average request err %v", err)
		}
	}

	response, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("receive average response err %v", err)
	}

	fmt.Println(response.GetResult())
}

func callMax(c calculatorpb.CalculatorServiceClient) {
	log.Println("Call max")

	//var wg sync.WaitGroup

	// work with chan
	waitc := make(chan struct{})

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Call Max api is error: %f", err)
	}

	//wg.Add(2)
	// request stream data
	go func() {
		//defer wg.Done()

		listReq := []calculatorpb.MaxRequest{
			calculatorpb.MaxRequest{
				Num: 10,
			},
			calculatorpb.MaxRequest{
				Num: 20,
			},
			calculatorpb.MaxRequest{
				Num: 30,
			},
			calculatorpb.MaxRequest{
				Num: 2,
			},
			calculatorpb.MaxRequest{
				Num: 5,
			},
			calculatorpb.MaxRequest{
				Num: 110,
			},
		}

		for _, req := range listReq {
			err = stream.Send(&req)

			if err != nil {
				log.Fatalf("send max api request err %v", err)
			}
		}

		stream.CloseSend()
	}()

	go func() {
		//defer wg.Done()

		for {
			resp, errRcv := stream.Recv()

			if errRcv == io.EOF {
				//log.Fatalf("Read max api request err %v", err)
				fmt.Println("Ending fine max api...")
				break
			}

			if errRcv != nil {
				log.Fatalf("recv find max err %v", err)

				break
			}

			fmt.Printf("max: %v\n", resp.GetResult())
		}
		close(waitc)
	}()

	//wg.Wait()
	<-waitc
}

func callSqrt(c calculatorpb.CalculatorServiceClient, num int32) {
	log.Println("Calling square root api")

	resp, err := c.Square(context.Background(), &calculatorpb.SquareRequest{Num: num})

	if err != nil {
		log.Printf("Send square api err %v\n", err)

		if errStatus, ok := status.FromError(err); ok {
			log.Printf("err msg: %v\n", errStatus.Message())
			log.Printf("err code: %v\n", errStatus.Code())

			if errStatus.Code() == codes.InvalidArgument {
				log.Printf("InvalidArgument num %v", num)

				return
			}
		}

	}

	fmt.Printf("sum square root response %v\n", resp.GetSquareRoot())
}
