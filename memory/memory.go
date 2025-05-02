package memory

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func MemInfo() {
	VMemory, _ := mem.VirtualMemory() //Virtual Memory = RAM + Swap
	fmt.Println("Memory Info:")
	fmt.Println("total: ", float64(VMemory.Total)/1e9)
	fmt.Println("free: ", float64(VMemory.Free)/1e9)
	fmt.Printf("Available: %.2f GB\n", float64(VMemory.Available)/1e9)
	fmt.Printf("UsedPercent: %.2f%%\n", VMemory.UsedPercent)
	swap, _ := mem.SwapMemory()
	fmt.Println("have swap:", swap.Total > 0)
}
