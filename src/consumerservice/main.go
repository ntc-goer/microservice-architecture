package main

import (
	"context"
	"fmt"
	"github.com/ntc-goer/microservice-examples/consumer_service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server struct {
	proto.ConsumerServiceServer
}

func (s *Server) VerifyUser(ctx context.Context, req *proto.VerifyUserRequest) (*proto.VerifyUserResponse, error) {
	fmt.Printf("Verifing user with ID %s", req.Id)
	time.Sleep(5 * time.Second)
	fmt.Printf("Verify user with ID %s valid", req.Id)
	return &proto.VerifyUserResponse{
		IsOk: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", 8080))
	if err != nil {
		log.Fatalf("Listen port fail %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterConsumerServiceServer(grpcServer, &Server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Start server fail")
	}
}
