package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ntc-goer/microservice-examples/orderservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type Server struct {
	proto.OrderServiceServer
}

func (s *Server) Order(ctx context.Context, orderReq *proto.OrderRequest) (*proto.OrderResponse, error) {
	fmt.Println("Create Order")
	return &proto.OrderResponse{
		IsOk: true,
	}, nil
}

func main() {
	// Set up a connection to the consumer server.
	consumerServiceAddr := "localhost:50000"
	conn, err := grpc.Dial(consumerServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	defer conn.Close()

	// start listening to requests from the gateway server
	mux := runtime.NewServeMux()
	if err = proto.RegisterOrderServiceHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the order server: %v", err)
	}
	addr := "localhost:8080"
	fmt.Println("API gateway server is running on " + addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}
}
