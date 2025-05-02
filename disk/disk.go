package disk

import (
	"fmt"

	disk "github.com/shirou/gopsutil/v3/disk"
)

func DiskInfo() {
	Partitions, _ := disk.Partitions(false)
	for _, partition := range Partitions {
		fmt.Println("Mountpoint: ", partition.Mountpoint)
		fmt.Println("Device: ", partition.Device)
		fmt.Println("File System Type:", partition.Fstype)
		usage, _ := disk.Usage(partition.Mountpoint)
		fmt.Printf("   Total: %.2f GB\n", float64(usage.Total)/1e9)
		fmt.Printf("   Used: %.2f GB (%.2f%%)\n", float64(usage.Used)/1e9, usage.UsedPercent)
		fmt.Printf("   Free: %.2f GB\n", float64(usage.Free)/1e9)

	}
}
