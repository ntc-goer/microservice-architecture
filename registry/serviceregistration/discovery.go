package serviceregistration

import (
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/consul"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/inmem"
)

func GetDiscovery(dcType common.DiscoveryType) (common.DiscoveryI, error) {
	if dcType == common.DC_INMEM {
		return inmem.NewRegistry()
	} else {
		return consul.NewRegistry()
	}
}
