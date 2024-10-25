package main

import (
	"context"
	"fmt"
	"gateway.com/gatewaypb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	gatewaypb.UnimplementedGatewayServiceServer
}

func (*server) Echo(ctx context.Context, req *gatewaypb.StringMessage) (*gatewaypb.StringMessage, error) {
	log.Printf("receive message %s\n", req.GetMsg())

	return &gatewaypb.StringMessage{
		Msg: req.GetMsg(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50080")

	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	gatewaypb.RegisterGatewayServiceServer(s, &server{})

	fmt.Println("Demo gateway service is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
