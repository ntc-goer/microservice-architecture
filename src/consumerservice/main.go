package main

import (
	"context"
	"fmt"
	"github.com/ntc-goer/microservice-examples/consumerservice/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const _SERVICENAME = "consumerservice"

type ServerImpl struct {
	proto.UnimplementedConsumerServiceServer
}

func (s *ServerImpl) VerifyUser(ctx context.Context, req *proto.VerifyUserRequest) (*proto.VerifyUserResponse, error) {
	fmt.Printf("Verifing user with ID %s", req.Id)
	time.Sleep(5 * time.Second)
	fmt.Printf("Verify user with ID %s valid", req.Id)
	return &proto.VerifyUserResponse{
		IsOk: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("Listen port fail %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterConsumerServiceServer(grpcServer, &ServerImpl{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Start server fail")
	}
}
