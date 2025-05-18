package cpu

import (
	prometheus "github.com/prometheus/client_golang/prometheus"
)

var (
	CpuTotalUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_cpu_usage_percentage",
		Help: "Total CPU usage as a percentage",
	})

	CpuPerCoreUsage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pc_cpu_usage_per_core_percent",
		Help: "CPU usage per logical core.",
	}, []string{"core"})

	CpuLogicalCores = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pc_cpu_logical_cores",
		Help: "Number of logical CPU cores.",
	})

	CpuPhysicalCores = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pc_cpu_physical_cores",
		Help: "Number of physical CPU cores.",
	})
)

func init() {
	prometheus.MustRegister(CpuTotalUsage)
	prometheus.MustRegister(CpuPerCoreUsage)
	prometheus.MustRegister(CpuLogicalCores)
	prometheus.MustRegister(CpuPhysicalCores)
}
