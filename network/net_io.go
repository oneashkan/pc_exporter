package network

import (
	"fmt"

	net "github.com/shirou/gopsutil/v3/net"
)

func GetNetworkIO() {

	ioStats, _ := net.IOCounters(true)
	fmt.Println("Network IO:")
	for _, stat := range ioStats {
		fmt.Println("Name:", stat.Name)
		fmt.Printf("Bytes Sent: %.2f MB\n", float64(stat.BytesSent)/1024/1024)
		fmt.Printf("Bytes Recv: %.2f MB\n", float64(stat.BytesRecv)/1024/1024)
	}
}
