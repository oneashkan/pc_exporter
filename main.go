package main

import (
	"fmt"
	"go-version/common"
	"go-version/cpu"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	if common.CanMonitor() {
		//power.GetBatteryInfo()
		for {
			cpu.UpdateCpuMetrics()
			//disk.DiskInfo()
			//memory.MemInfo()
			//network.GetNetworkIO()
			//network.NetInfo()
			http.Handle("/metrics/", promhttp.Handler())
			http.ListenAndServe(":3001", nil)
			fmt.Println("Server Started")

		}
	}
}
