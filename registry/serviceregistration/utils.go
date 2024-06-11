package serviceregistration

import (
	"fmt"
	"net"
	"time"
)

const _LOCAL_SERVICE = true

func GenerateInstanceId(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, time.Now().UnixMilli())
}

func GetCurrentIP() string {
	if _LOCAL_SERVICE {
		return "localhost"
	}
	interfaces, err := net.Interfaces()
	if err != nil {
		return "localhost"
	}

	for _, iface := range interfaces {
		// Skip down interfaces
		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		// Skip loopback interfaces
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return "localhost"
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// Skip IPv6 and loopback addresses
			if ip == nil || ip.IsLoopback() || ip.To4() == nil {
				continue
			}

			return ip.String()
		}
	}

	return "localhost"
}
