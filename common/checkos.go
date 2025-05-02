package common

import (
	"runtime"
)

func DetectOs() string {
	return runtime.GOOS
}

func CanMonitor() bool {
	os := DetectOs()
	switch os {
	case "linux":
		return true
	case "windows":
		return false
	default:
		return false
	}

}
