package servicediscovery

import (
	"fmt"
	"github.com/ntc-goer/microservice-examples/registry/servicediscovery/common"
	"github.com/ntc-goer/microservice-examples/registry/servicediscovery/consul"
	"github.com/ntc-goer/microservice-examples/registry/servicediscovery/inmem"
	"time"
)

func GenerateInstanceId(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, time.Now().UnixMilli())
}

func GetDiscovery(dcType common.DiscoveryType) (common.DiscoveryI, error) {
	if dcType == common.DC_INMEM {
		return consul.NewRegistry()
	} else {
		return inmem.NewRegistry()
	}
}
