package cpu

import (
	"fmt"
	"time"

	cpu "github.com/shirou/gopsutil/v3/cpu"
)

func CountLogicalCores(physical bool) int {
	LogicalCores, _ := cpu.Counts(physical)
	return LogicalCores

}
func CpuUsage(percpu bool) []float64 {
	usage, _ := cpu.Percent(1*time.Minute, percpu)
	return usage
}

func GetCpuInfo() {

	fmt.Printf("Usage: %.2f%%\n", CpuUsage(false)[0])
	fmt.Println("Number of Logical Core: ", CountLogicalCores(true))
	fmt.Println("Number of physical Core: ", CountLogicalCores(false))
	for index, usage := range CpuUsage(true) {
		fmt.Printf("core index %d Usage: %.2f%%\n", index+1, usage)
	}
	/*
		info, _ := cpu.Info()
		for _, item := range info {
			fmt.Println("model name: ", item.ModelName)
			fmt.Println("\n", item.Mhz)
		}*/
}
