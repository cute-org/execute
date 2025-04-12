package utils

import (
	"fmt"
	"net"
)

// GetLocalhostIP returns the loopback IP address (always 127.0.0.1)
func getLocalhostIP() string {
	return "localhost"
}

// GetNetworkIP returns the first non-loopback IPv4 address on the machine
func getNetworkIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				return ip4.String()
			}
		}
	}
	return "unknown"
}

func PrintIPs(addr string) {
	fmt.Printf("  ➜  Local:   http://%s%s/\n", getLocalhostIP(), addr)
	fmt.Printf("  ➜  Network: http://%s%s/\n", getNetworkIP(), addr)
}
