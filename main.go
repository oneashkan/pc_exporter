package main

import (
	"go-version/common"
	"go-version/cpu"
	"go-version/disk"
	"go-version/memory"
	"go-version/network"
	"go-version/power"
)

func main() {
	if common.CanMonitor() {
		power.GetBatteryInfo()
		cpu.GetCpuInfo()
		disk.DiskInfo()
		memory.MemInfo()
		network.GetNetworkIO()
		network.NetInfo()
	}
}
