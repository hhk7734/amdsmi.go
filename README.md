## Go Bindings for AMD System Management Interface (AMD SMI) Library

This repository provides Go bindings for [AMD System Management Interface (AMD SMI) Library](https://github.com/ROCm/amdsmi)

> [!IMPORTANT]
> This binding dynamically loads the shared library `libamd_smi.so` at runtime.
>
> You do not need the `libamd_smi.so` to build code that uses this binding. However, at runtime, the shared library must be available in the system library path. Typically, The shared library is installed in `/opt/rocm/lib` when you install the ROCm software stack.
>
> For more information on installing AMD SMI, see the [Install AMD SMI](https://rocm.docs.amd.com/projects/amdsmi/en/latest/install/install.html) guide.

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/hhk7734/amdsmi.go"
)

func main() {
	smi := amdsmi.New()
	if err := smi.Init(amdsmi.INIT_AMD_GPUS); err != nil {
		panic(err)
	}

	defer func() {
		if err := smi.Shutdown(); err != nil {
			panic(err)
		}
	}()

	sockets, err := smi.Sockets()
	if err != nil {
		panic(err)
	}

	for _, socket := range sockets {
		processors, err := socket.Processors()
		if err != nil {
			panic(err)
		}

		for deviceIndex, processor := range processors {
			temp, _ := processor.Temperature(amdsmi.TEMP_TYPE_EDGE, amdsmi.TEMP_CURRENT)

			fmt.Printf("DeviceID: %d\tTemp: %.1f°C\n", deviceIndex, float64(temp))
		}
	}
}
```

```shell
$ go run main.go
DeviceID: 0     Temp: 50.0°C
DeviceID: 0     Temp: 48.0°C
DeviceID: 0     Temp: 43.0°C
DeviceID: 0     Temp: 42.0°C
DeviceID: 0     Temp: 42.0°C
DeviceID: 0     Temp: 42.0°C
DeviceID: 0     Temp: 49.0°C
DeviceID: 0     Temp: 51.0°C
```
