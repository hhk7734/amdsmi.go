package main

import (
	"fmt"

	"github.com/hhk7734/amdsmi.go"
)

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

// type GPUMetrics struct {
// 	CommonHeader MetricsTableHeader

// 	TemperatureEdge    uint16
// 	TemperatureHotspot uint16
// 	TemperatureMem     uint16
// 	TemperatureVRGFX   uint16
// 	TemperatureVRSoC   uint16
// 	TemperatureVRMem   uint16

// 	AverageGFXActivity uint16
// 	AverageUMCActivity uint16
// 	AverageMMActivity  uint16

// 	AverageSocketPower uint16
// 	EnergyAccumulator  uint64

// 	SystemClockCounter uint64

// 	AverageGFXCLKFrequency uint16
// 	AverageSoCCLKFrequency uint16
// 	AverageUCLKFrequency   uint16
// 	AverageVCLK0Frequency  uint16
// 	AverageDCLK0Frequency  uint16
// 	AverageVCLK1Frequency  uint16
// 	AverageDCLK1Frequency  uint16

// 	CurrentGFXCLK uint16
// 	CurrentSoCCLK uint16
// 	CurrentUCLK   uint16
// 	CurrentVCLK0  uint16
// 	CurrentDCLK0  uint16
// 	CurrentVCLK1  uint16
// 	CurrentDCLK1  uint16

// 	ThrottleStatus uint32

// 	CurrentFanSpeed uint16

// 	PCIeLinkWidth uint16
// 	PCIeLinkSpeed uint16

// 	GFXActivityAcc uint32
// 	MemActivityAcc uint32
// 	TemperatureHBM [NUM_HBM_INSTANCE]uint16

// 	FirmwareTimestamp uint64

// 	VoltageSoC uint16
// 	VoltageGFX uint16
// 	VoltageMem uint16

// 	IndepThrottleStatus uint64

// 	CurrentSocketPower uint16

// 	GFXCLKLockStatus uint32

// 	XGMILinkWidth uint16
// 	XGMILinkSpeed uint16

// 	PCIeBandwidthAcc uint64

// 	PCIeBandwidthInst uint64

// 	PCIeL0ToRecovCountAcc uint64

// 	PCIeReplayCountAcc uint64

// 	PCIeReplayRoverCountAcc uint64

// 	XGMIReadDataAcc  [MAX_NUM_XGMI_LINKS]uint64
// 	XGMIWriteDataAcc [MAX_NUM_XGMI_LINKS]uint64

// 	CurrentGFXCLKs [MAX_NUM_GFX_CLKS]uint16
// 	CurrentSoCCLKs [MAX_NUM_CLKS]uint16
// 	CurrentVCLK0s  [MAX_NUM_CLKS]uint16
// 	CurrentDCLK0s  [MAX_NUM_CLKS]uint16

// 	JPEGActivity [MAX_NUM_JPEG]uint16

// 	PCIeNAKSentCountAcc uint32

// 	PCIeNAKRcvdCountAcc uint32
// }

func main() {
	smi := amdsmi.New()
	assert(smi.Init(amdsmi.INIT_AMD_GPUS))

	defer func() {
		assert(smi.Shutdown())
	}()

	sockets, err := smi.Sockets()
	assert(err)

	processors := make([]*amdsmi.Processor, 0)

	for _, socket := range sockets {
		socketInfo, err := socket.Info()
		assert(err)
		fmt.Printf("Socket Info: %s\n", socketInfo)

		ps, err := socket.Processors()
		assert(err)
		processors = append(processors, ps...)
	}

	for deviceIndex, processor := range processors {
		type_, _ := processor.Type()
		gpuID, _ := processor.GPUID()
		gpuRevision, _ := processor.GPURevision()
		gpuVendorName, _ := processor.GPUVendorName()
		gpuVRAMVendor, _ := processor.GPUVRAMVendor()
		gpuSubsystemID, _ := processor.GPUSubsystemID()
		gpuSubsystemName, _ := processor.GPUSubsystemName()
		gpuPCIeBandwidth, _ := processor.GPUPCIeBandwidth()
		gpuBDFID, _ := processor.GPUBDFID()
		gpuMemoryTotalFirst, _ := processor.GPUMemoryTotal(amdsmi.MEM_TYPE_FIRST)
		gpuMemoryUsageFirst, _ := processor.GPUMemoryUsage(amdsmi.MEM_TYPE_FIRST)
		gpuFanRPM, _ := processor.GPUFanRPM(0)
		gpuFanSpeed, _ := processor.GPUFanSpeed(0)
		gpuFanSpeedMax, _ := processor.GPUFanSpeedMax(0)
		temperatureCurrent, _ := processor.Temperature(amdsmi.TEMP_TYPE_FIRST, amdsmi.TEMP_CURRENT)
		gpuMetrics, _ := processor.GPUMetricsInfo()

		fmt.Printf("Device Index: %d\n", deviceIndex)
		fmt.Printf("\tType: %s\n", type_)
		fmt.Printf("\tGPU ID: %d\n", gpuID)
		fmt.Printf("\tGPU Revision: %d\n", gpuRevision)
		fmt.Printf("\tGPU Vendor Name: %s\n", gpuVendorName)
		fmt.Printf("\tGPU VRAM Vendor: %s\n", gpuVRAMVendor)
		fmt.Printf("\tGPU Subsystem ID: %d\n", gpuSubsystemID)
		fmt.Printf("\tGPU Subsystem Name: %s\n", gpuSubsystemName)
		fmt.Printf("\tGPU PCIe Bandwidth: %v\n", gpuPCIeBandwidth)
		fmt.Printf("\tGPU BDFID: %d\n", gpuBDFID)
		fmt.Printf("\tGPU Memory Total(first): %d MB\n", gpuMemoryTotalFirst/1024/1024)
		fmt.Printf("\tGPU Memory Usage(first): %d MB\n", gpuMemoryUsageFirst/1024/1024)
		fmt.Printf("\tGPU Fan RPM(index: 0): %d RPM\n", gpuFanRPM)
		fmt.Printf("\tGPU Fan Speed(index: 0): %d\n", gpuFanSpeed)
		fmt.Printf("\tGPU Fan Speed Max(index: 0): %d\n", gpuFanSpeedMax)
		fmt.Printf("\tTemperature(first, current): %d\n", temperatureCurrent)
		fmt.Printf("\tGPU Metrics Info:\n")
		fmt.Printf("\t\tCommon Header: %v\n", gpuMetrics.CommonHeader)
		fmt.Printf("\t\tTemperature Edge: %d °C\n", gpuMetrics.TemperatureEdge)
		fmt.Printf("\t\tTemperature Hotspot: %d °C\n", gpuMetrics.TemperatureHotspot)
		fmt.Printf("\t\tTemperature Memory: %d °C\n", gpuMetrics.TemperatureMem)
		fmt.Printf("\t\tTemperature VRGFX: %d °C\n", gpuMetrics.TemperatureVRGFX)
		fmt.Printf("\t\tTemperature VRSoC: %d °C\n", gpuMetrics.TemperatureVRSoC)
		fmt.Printf("\t\tTemperature VRMem: %d °C\n", gpuMetrics.TemperatureVRMem)
		fmt.Printf("\t\tAverage GFX Activity: %d %%\n", gpuMetrics.AverageGFXActivity)
		fmt.Printf("\t\tAverage UMC Activity: %d %%\n", gpuMetrics.AverageUMCActivity)
		fmt.Printf("\t\tAverage MM Activity: %d %%\n", gpuMetrics.AverageMMActivity)
		fmt.Printf("\t\tAverage Socket Power: %d W\n", gpuMetrics.AverageSocketPower)
		fmt.Printf("\t\tEnergy Accumulator: %d\n", gpuMetrics.EnergyAccumulator)
		fmt.Printf("\t\tSystem Clock Counter: %d ns\n", gpuMetrics.SystemClockCounter)
		fmt.Printf("\t\tAverage GFX CLK Frequency: %d MHz\n", gpuMetrics.AverageGFXCLKFrequency)
		fmt.Printf("\t\tAverage SoC CLK Frequency: %d MHz\n", gpuMetrics.AverageSoCCLKFrequency)
		fmt.Printf("\t\tAverage UCLK Frequency: %d MHz\n", gpuMetrics.AverageUCLKFrequency)
		fmt.Printf("\t\tAverage VCLK0 Frequency: %d MHz\n", gpuMetrics.AverageVCLK0Frequency)
		fmt.Printf("\t\tAverage DCLK0 Frequency: %d MHz\n", gpuMetrics.AverageDCLK0Frequency)
		fmt.Printf("\t\tAverage VCLK1 Frequency: %d MHz\n", gpuMetrics.AverageVCLK1Frequency)
		fmt.Printf("\t\tAverage DCLK1 Frequency: %d MHz\n", gpuMetrics.AverageDCLK1Frequency)
		fmt.Printf("\t\tCurrent GFX CLK: %d MHz\n", gpuMetrics.CurrentGFXCLK)
		fmt.Printf("\t\tCurrent SoC CLK: %d MHz\n", gpuMetrics.CurrentSoCCLK)
		fmt.Printf("\t\tCurrent UCLK: %d MHz\n", gpuMetrics.CurrentUCLK)
		fmt.Printf("\t\tCurrent VCLK0: %d MHz\n", gpuMetrics.CurrentVCLK0)
		fmt.Printf("\t\tCurrent DCLK0: %d MHz\n", gpuMetrics.CurrentDCLK0)
		fmt.Printf("\t\tCurrent VCLK1: %d MHz\n", gpuMetrics.CurrentVCLK1)
		fmt.Printf("\t\tCurrent DCLK1: %d MHz\n", gpuMetrics.CurrentDCLK1)
		fmt.Printf("\t\tThrottle Status: %d\n", gpuMetrics.ThrottleStatus)
		fmt.Printf("\t\tCurrent Fan Speed: %d RPM\n", gpuMetrics.CurrentFanSpeed)
		fmt.Printf("\t\tPCIe Link Width: %d\n", gpuMetrics.PCIeLinkWidth)
		fmt.Printf("\t\tPCIe Link Speed: %.1f GT/s\n", float32(gpuMetrics.PCIeLinkSpeed)*0.1)
		fmt.Printf("\t\tGFX accumulated Activity: %d\n", gpuMetrics.GFXActivityAcc)
		fmt.Printf("\t\tMem accumulated Activity: %d\n", gpuMetrics.MemActivityAcc)
		fmt.Printf("\t\tTemperature HBM: %v °C\n", gpuMetrics.TemperatureHBM)
		fmt.Printf("\t\tFirmware Timestamp: %d ns\n", gpuMetrics.FirmwareTimestamp*10)
		fmt.Printf("\t\tVoltage SoC: %d mV\n", gpuMetrics.VoltageSoC)
		fmt.Printf("\t\tVoltage GFX: %d mV\n", gpuMetrics.VoltageGFX)
		fmt.Printf("\t\tVoltage Mem: %d mV\n", gpuMetrics.VoltageMem)
		fmt.Printf("\t\tIndep Throttle Status: %d\n", gpuMetrics.IndepThrottleStatus)
		fmt.Printf("\t\tCurrent Socket Power: %d W\n", gpuMetrics.CurrentSocketPower)
		fmt.Printf("\t\tGFX CLK Lock Status: %d\n", gpuMetrics.GFXCLKLockStatus)
		fmt.Printf("\t\tXGMI Link Width: %d\n", gpuMetrics.XGMILinkWidth)
		fmt.Printf("\t\tXGMI Link Speed: %d GB/s\n", gpuMetrics.XGMILinkSpeed)
		fmt.Printf("\t\tPCIe accumulated Bandwidth: %d GB/s\n", gpuMetrics.PCIeBandwidthAcc)
		fmt.Printf("\t\tPCIe instantaneous Bandwidth: %d GB/s\n", gpuMetrics.PCIeBandwidthInst)
		fmt.Printf("\t\tPCIe L0 To Recovery accumulated Count: %d\n", gpuMetrics.PCIeL0ToRecovCountAcc)
		fmt.Printf("\t\tPCIe Replay accumulated Count: %d\n", gpuMetrics.PCIeReplayCountAcc)
		fmt.Printf("\t\tPCIe Replay Rover accumulated Count: %d\n", gpuMetrics.PCIeReplayRoverCountAcc)
		fmt.Printf("\t\tXGMI Read Data: %v KB\n", gpuMetrics.XGMIReadDataAcc)
		fmt.Printf("\t\tXGMI Write Data: %v KB\n", gpuMetrics.XGMIWriteDataAcc)
		fmt.Printf("\t\tCurrent GFX CLKs: %v MHz\n", gpuMetrics.CurrentGFXCLKs)
		fmt.Printf("\t\tCurrent SoC CLKs: %v MHz\n", gpuMetrics.CurrentSoCCLKs)
		fmt.Printf("\t\tCurrent VCLK0s: %v MHz\n", gpuMetrics.CurrentVCLK0s)
		fmt.Printf("\t\tCurrent DCLK0s: %v MHz\n", gpuMetrics.CurrentDCLK0s)
		fmt.Printf("\t\tJPEG Activity: %v\n", gpuMetrics.JPEGActivity)
		fmt.Printf("\t\tPCIe NAK Sent accumulated Count: %d\n", gpuMetrics.PCIeNAKSentCountAcc)
		fmt.Printf("\t\tPCIe NAK Rcvd accumulated Count: %d\n", gpuMetrics.PCIeNAKRcvdCountAcc)
	}
}
