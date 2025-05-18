package cpu

import "fmt"

func UpdateCpuMetrics() {

	PerCpuUsage := CpuUsage(true)
	for index, value := range PerCpuUsage {
		CpuPerCoreUsage.WithLabelValues(
			"core_" + fmt.Sprint(index+1),
		).Set(value)
	}

	CpuTotalUsage.Set(CpuUsage(false)[0])

	CpuPhysicalCores.Set(float64(CountLogicalCores(true)))

	CpuLogicalCores.Set(float64(CountLogicalCores(false)))
}
