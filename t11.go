package main

import (
	"bytes"
	"fmt"
	"net"
	"sort"
)

func main() {
	ips := []string{
		"192.168.1.5",
		"69.52.220.44",
		"10.152.16.23",
		"192.168.3.10",
		"192.168.1.4",
		"192.168.1.41",
	}

		realIPs := make([]net.IP, 0, len(ips))

		for _, ip := range ips {
			realIPs = append(realIPs, net.ParseIP(ip))
		}

		sort.Slice(realIPs, func(i, j int) bool {
			return bytes.Compare(realIPs[i], realIPs[j]) < 0
		})

		for _, ip := range realIPs {
			fmt.Printf("%s\n", ip)
		}

}
