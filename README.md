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

	processors := make([]*amdsmi.Processor, 0)

	for _, socket := range sockets {
		ps, err := socket.Processors()
		if err != nil {
			panic(err)
		}
		processors = append(processors, ps...)
	}

	fmt.Println("GPU  POWER  GPU_TEMP  MEM_TEMP  GFX_UTIL  GFX_CLOCK  MEM_UTIL  MEM_CLOCK  ENC_UTIL  ENC_CLOCK  DEC_UTIL  DEC_CLOCK  SINGLE_ECC  DOUBLE_ECC  PCIE_REPLAY  VRAM_USED  VRAM_TOTAL   PCIE_BW")

	for deviceIndex, processor := range processors {
		gpuMetrics, _ := processor.GPUMetricsInfo()
		vramUsed, _ := processor.GPUMemoryUsage(amdsmi.MEM_TYPE_VRAM)
		vramTotal, _ := processor.GPUMemoryTotal(amdsmi.MEM_TYPE_VRAM)

		fmt.Printf("%3d  %3d W    %3d °C    %3d °C     %3d %%   %4d MHz     %3d %%   %4d MHz                                                                                 %6d MB   %6d MB\n",
			deviceIndex,
			gpuMetrics.CurrentSocketPower,
			gpuMetrics.TemperatureHotspot,
			gpuMetrics.TemperatureMem,
			gpuMetrics.AverageGFXActivity,
			gpuMetrics.CurrentGFXCLK,
			gpuMetrics.AverageUMCActivity,
			gpuMetrics.CurrentUCLK,
			vramUsed/1024/1024,
			vramTotal/1024/1024,
		)
	}
}
```

```shell
$ amd-smi monitor
GPU  POWER  GPU_TEMP  MEM_TEMP  GFX_UTIL  GFX_CLOCK  MEM_UTIL  MEM_CLOCK  ENC_UTIL  ENC_CLOCK  DEC_UTIL  DEC_CLOCK  SINGLE_ECC  DOUBLE_ECC  PCIE_REPLAY  VRAM_USED  VRAM_TOTAL   PCIE_BW
  0  131 W     42 °C     34 °C       0 %    135 MHz       0 %    901 MHz     0.0 %     29 MHz       N/A     22 MHz           0           0            0     285 MB   196300 MB  166 Mb/s
  1  137 W     47 °C     40 °C       0 %    135 MHz       0 %    901 MHz     0.0 %     29 MHz       N/A     22 MHz           0           0            0     285 MB   196300 MB   28 Mb/s
  2  140 W     47 °C     41 °C       0 %    135 MHz       0 %    901 MHz     0.0 %     29 MHz       N/A     22 MHz           0           0            0     285 MB   196300 MB  177 Mb/s
  3  129 W     42 °C     34 °C       0 %    146 MHz       0 %    901 MHz     0.0 %     29 MHz       N/A     22 MHz           0           0            0     285 MB   196300 MB  137 Mb/s
  4  672 W     74 °C     57 °C     100 %   2007 MHz       1 %   1300 MHz     0.0 %     42 MHz       N/A     36 MHz           0           0            0   51666 MB   196300 MB 6213 Mb/s
  5  166 W     63 °C     57 °C       0 %    136 MHz       0 %    901 MHz     0.0 %     29 MHz       N/A     22 MHz           0           0            0     285 MB   196300 MB   58 Mb/s
  6  156 W     63 °C     57 °C       0 %    159 MHz       0 %    901 MHz     0.0 %     29 MHz       N/A     22 MHz           0           0            0     285 MB   196300 MB  190 Mb/s
  7  747 W     95 °C     58 °C     100 %   1732 MHz       2 %   1300 MHz     0.0 %     42 MHz       N/A     36 MHz           0           0            0   51408 MB   196300 MB 6346 Mb/s

$ go run ./monitor.go
GPU  POWER  GPU_TEMP  MEM_TEMP  GFX_UTIL  GFX_CLOCK  MEM_UTIL  MEM_CLOCK  ENC_UTIL  ENC_CLOCK  DEC_UTIL  DEC_CLOCK  SINGLE_ECC  DOUBLE_ECC  PCIE_REPLAY  VRAM_USED  VRAM_TOTAL   PCIE_BW
  0  131 W     42 °C     34 °C       0 %    182 MHz       0 %    900 MHz                                                                                    285 MB   196592 MB
  1  137 W     47 °C     40 °C       0 %    168 MHz       0 %    900 MHz                                                                                    285 MB   196592 MB
  2  141 W     47 °C     41 °C       0 %    167 MHz       0 %    900 MHz                                                                                    285 MB   196592 MB
  3  129 W     41 °C     34 °C       0 %    184 MHz       0 %    900 MHz                                                                                    285 MB   196592 MB
  4  644 W     76 °C     57 °C     100 %   2022 MHz       1 %   1300 MHz                                                                                  51666 MB   196592 MB
  5  166 W     63 °C     57 °C       0 %    169 MHz       0 %    900 MHz                                                                                    285 MB   196592 MB
  6  156 W     63 °C     57 °C       0 %    187 MHz       0 %    900 MHz                                                                                    285 MB   196592 MB
  7  596 W     91 °C     56 °C      74 %   1298 MHz       2 %   1300 MHz                                                                                  51408 MB   196592 MB
```
