package main

import (
	"go-version/common"

	"go-version/cpu"
	"go-version/power"
)

func main() {
	if common.CanMonitor() {
		power.GetBatteryInfo()
		cpu.GetCpuInfo()
	}
}
