package network

import (
	"fmt"

	net "github.com/shirou/gopsutil/v3/net"
)

func NetInfo() {

	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		fmt.Println("Name:", iface.Name)
		fmt.Println("Hardware Addr (MAC):", iface.HardwareAddr)
		fmt.Println("Flags:", iface.Flags)

		for _, addr := range iface.Addrs {
			fmt.Println("	Address:", addr.Addr)
		}

	}
}
