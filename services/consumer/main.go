package main

import (
	"context"
	"fmt"
	pb "github.com/ntc-goer/microservice-examples/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const _SERVICENAME = "consumer"

type ServerImpl struct {
	pb.UnimplementedConsumerServiceServer
}

func (s *ServerImpl) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*pb.VerifyUserResponse, error) {
	fmt.Printf("Verifing user with ID %s", req.Id)
	time.Sleep(5 * time.Second)
	fmt.Printf("Verify user with ID %s valid", req.Id)
	return &pb.VerifyUserResponse{
		IsOk: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("Listen port fail %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterConsumerServiceServer(grpcServer, &ServerImpl{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Start server fail")
	}
}
