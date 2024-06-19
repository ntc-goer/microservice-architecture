package main

import (
	"context"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"log"
)

func main() {
	dp, err := InitializeDependency("consul")
	if err != nil {
		log.Fatalf("fail to init dependency %v", err)
	}
	ctx := context.Background()

	// Register to discovery service
	instanceId := serviceregistration.GenerateInstanceId(dp.Config.MailServiceName)
	if err := dp.ServiceDiscovery.RegisterService(instanceId, dp.Config.MailServiceName, serviceregistration.GetCurrentIP(), dp.Config.ServicePort, common.GRPC_CHECK_TYPE); err != nil {
		log.Fatalf("RegisterService fail: %v", err)
	}
	defer dp.ServiceDiscovery.Deregister(ctx, instanceId)

	if err := dp.Queue.Connect(dp.Config.QueueAddress); err != nil {
		log.Fatalf("QueueConnect fail: %v", err)
	}
	defer dp.Queue.Close()
	// Start consuming message from queue
	for {
		if err := dp.Queue.QueueSubscribe(dp.Config.MailQueueSubject, dp.Config.MailQueueGroup, dp.Service.Handle); err != nil {
			log.Fatal(err)
		}
	}
}
