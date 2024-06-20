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
	instanceId := serviceregistration.GenerateInstanceId(dp.Config.Service.MailServiceName)
	if err := dp.ServiceDiscovery.RegisterService(instanceId, dp.Config.Service.MailServiceName, serviceregistration.GetCurrentIP(), dp.Config.ServicePort, common.GRPC_CHECK_TYPE); err != nil {
		log.Fatalf("RegisterService fail: %v", err)
	}
	defer dp.ServiceDiscovery.Deregister(ctx, instanceId)

	// Connect to Broker
	if err := dp.Broker.Connect(dp.Config.Broker.Address); err != nil {
		log.Fatalf("BrokerConnect fail: %v", err)
	}
	defer dp.Broker.Close()
	// Start consuming message from broker
	for {
		if err := dp.Broker.QueueSubscribe(dp.Config.Broker.Subject.SendMail, dp.Config.Broker.Queue.Mail, dp.Service.Handle); err != nil {
			log.Fatal(err)
		}
	}
}
