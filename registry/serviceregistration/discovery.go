package serviceregistration

import (
	"fmt"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/consul"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/inmem"
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
