package utils

import (
	"fmt"
	"os"
	"runtime"
)


var (
	CURRENT_SYSTEM = runtime.GOOS
	WINDOWS = false
	LINUX = false
)


var InitialDirectory, err = os.Getwd()




func GetOperatingSystem() {
	switch CURRENT_SYSTEM {
	case "windows":
		WINDOWS = true
	case "linux":
		LINUX = true
	default:
		fmt.Println("Operating system not supported by mass")
		os.Exit(1)
	}
}
