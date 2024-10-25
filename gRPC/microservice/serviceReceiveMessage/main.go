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

type receiveServer struct {
	micropb.UnimplementedServiceReceivedMessageServer
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", "localhost:50052")

	if err != nil {
		log.Fatalln("Receive Failed to listen %v", err)
	}

	s := grpc.NewServer()

	micropb.RegisterServiceReceivedMessageServer(s, &receiveServer{})

	fmt.Println("Receive Service send message server is calling...!")

	errServ := s.Serve(lis)

	if errServ != nil {
		log.Fatalf("err while server %v", errServ)
	}
}

func (*receiveServer) ReceiveMessage(ctx context.Context, req *micropb.ReceiveRequestMessage) (*micropb.ReceiveResponseMessage, error) {
	return &micropb.ReceiveResponseMessage{
		Result: "service receive message success" + req.GetMessage(),
	}, nil
}

func main() {
	go startGRPCServer()

	fmt.Println("Starting HTTP server on :8080")
	http.HandleFunc("/sendMessage", func(writer http.ResponseWriter, request *http.Request) {
		sendMessage()
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sendMessage() {
	cc, err := grpc.Dial(":50051", grpc.WithInsecure())

	defer cc.Close()

	if err != nil {
		log.Fatalf("call port err %v\n", err)
	}

	client := micropb.NewServiceSendMessageClient(cc)

	message, err := client.SendMessage(context.Background(), &micropb.SendRequestMessage{
		Message: "Thanhdn \n",
	})

	if err != nil {
		log.Fatalf("call api send message error: %v\n", err)
	}

	fmt.Printf("Send message: %v\n", message)
}
