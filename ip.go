package gotool

import (
	"io"
	"net"
	"net/http"
	"strings"
)

func GetLocalIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range ifaces {
		if strings.Contains(iface.Name, "vm") {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ip := ipnet.IP
				return ip.String()
			}
		}
	}
	return ""
}

func GetNetworkIP() string {
	resp, err := http.Get("http://ifconfig.me/ip")
	if err == nil {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			return string(body)
		}
	}
	return ""
}
