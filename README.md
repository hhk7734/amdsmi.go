## Go Bindings for ROCm System Management Interface (ROCm SMI) Library

This repository provides Go bindings for [ROCm System Management Interface (ROCm SMI) Library](https://github.com/ROCm/rocm_smi_lib)

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/hhk7734/rocm_smi.go/pkg/rsmi"
)

func main() {
	smi := rsmi.New()
	if err := smi.Init(0); err != nil {
		panic(err)
	}

	num, err := smi.NumMonitorDevices()
	if err != nil {
		panic(err)
	}

	for deviceIndex := uint32(0); deviceIndex < num; deviceIndex++ {
		temp, _ := smi.DevTempMetric(deviceIndex, rsmi.TEMP_TYPE_EDGE, rsmi.TEMP_CURRENT)

		total, _ := smi.DevMemoryTotal(deviceIndex, rsmi.MEM_TYPE_VRAM)
		usage, _ := smi.DevMemoryUsage(deviceIndex, rsmi.MEM_TYPE_VRAM)

		vramUtil := 0.0
		if total != 0 {
			vramUtil = float64(usage) / float64(total) * 100
		}

		fmt.Printf("DeviceID: %d\tTemp: %.1fC\tVRAM: %.1f%%\n", deviceIndex, temp, vramUtil)
	}

	if err := smi.Shutdown(); err != nil {
		panic(err)
	}
}
```

```shell
$ go run main.go
DeviceID: 0     Temp: 70.0C     VRAM: 25.7%
DeviceID: 1     Temp: 66.0C     VRAM: 25.7%
DeviceID: 2     Temp: 50.0C     VRAM: 13.9%
DeviceID: 3     Temp: 58.0C     VRAM: 13.9%
DeviceID: 4     Temp: 43.0C     VRAM: 14.7%
DeviceID: 5     Temp: 42.0C     VRAM: 14.7%
DeviceID: 6     Temp: 52.0C     VRAM: 17.0%
DeviceID: 7     Temp: 56.0C     VRAM: 17.0%
```