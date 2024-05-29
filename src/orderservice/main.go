package main

import (
	"fmt"
	"github.com/ntc-goer/microservice-examples/consumerservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", "8080"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connect consumer service")
	}
	defer conn.Close()
	client := proto.NewConsumerServiceClient(conn)
}
