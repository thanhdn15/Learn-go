package main

import (
	"context"
	"fmt"
	"github.com/thanhdn15/concrete_lean_go/gRPC/microservice/micropb"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type sendMessageServer struct {
	micropb.UnimplementedServiceSendMessageServer
}

func (*sendMessageServer) SendMessage(ctx context.Context, req *micropb.SendRequestMessage) (*micropb.SendResponseMessage, error) {
	return &micropb.SendResponseMessage{
		Result: "service send message " + req.GetMessage(),
	}, nil
}

func startGRPCService() {
	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Service Send failed to listen %v", err)
	}

	s := grpc.NewServer()

	micropb.RegisterServiceSendMessageServer(s, &sendMessageServer{})

	fmt.Println("Register Service send message server is calling...!")

	errServ := s.Serve(lis)

	if errServ != nil {
		log.Fatalf("Server error %v", errServ)
	}
}

func main() {
	go startGRPCService()

	fmt.Println("Starting HTTP server on :8081")

	http.HandleFunc("/receiveMessage", func(writer http.ResponseWriter, request *http.Request) {
		receiveMessage()
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func receiveMessage() {
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	defer cc.Close()

	if err != nil {
		log.Fatalf("call port err %v\n", err)
	}

	client := micropb.NewServiceReceivedMessageClient(cc)

	message, err := client.ReceiveMessage(context.Background(), &micropb.ReceiveRequestMessage{
		Message: "Message from send service thanhdn \n",
	})

	if err != nil {
		log.Fatalf("call grpc api err %v\n", err)
	}

	fmt.Println(message)
}
