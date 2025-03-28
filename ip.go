package gotool

import (
	"io"
	"net"
	"net/http"
	"strings"
)

// 获取本机IP
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

// 获取网络IP
func GetNetworkIP() string {
	resp, err := http.Get("http://ipv4.icanhazip.com")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	ip := string(body)
	// 替换ip的空格和换行
	ip = strings.ReplaceAll(ip, "\n", "")
	ip = strings.ReplaceAll(ip, " ", "")
	return ip
}
