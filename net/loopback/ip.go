package loopback

import (
	"net"
)

func IP() string {
	ips := IPs()

	if len(ips) == 0 {
		return ""
	} else {
		return ips[0]
	}
}

func IPs() []string {
	ips := make([]string, 0)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	return ips
}
