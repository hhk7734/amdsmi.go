package amdsmi

type MetricsTableHeader struct {
	StructureSize   uint16
	FormatRevision  uint8
	ContentRevision uint8
}

type GPUMetrics struct {
	CommonHeader MetricsTableHeader

	// Edge temperature (°C)
	TemperatureEdge uint16
	// Hotspot temperature (°C)
	TemperatureHotspot uint16
	// Memory temperature (°C)
	TemperatureMem uint16
	// VRGFX temperature (°C)
	TemperatureVRGFX uint16
	// VRSoC temperature (°C)
	TemperatureVRSoC uint16
	// VRMem temperature (°C)
	TemperatureVRMem uint16

	// Average GPU activity (%)
	AverageGFXActivity uint16
	// Average Universal Memory Controller activity (%)
	AverageUMCActivity uint16
	// Average MultiMedia activity (%)
	AverageMMActivity uint16

	// Average socket power (W)
	AverageSocketPower uint16
	// Energy accumulator
	EnergyAccumulator uint64

	// System clock counter (ns)
	SystemClockCounter uint64

	// Average GFX clock frequency (MHz)
	AverageGFXCLKFrequency uint16
	// Average SoC clock frequency (MHz)
	AverageSoCCLKFrequency uint16
	// Average UCLK frequency (MHz)
	AverageUCLKFrequency uint16
	// Average VCLK0 frequency (MHz)
	AverageVCLK0Frequency uint16
	// Average DCLK0 frequency (MHz)
	AverageDCLK0Frequency uint16
	// Average VCLK1 frequency (MHz)
	AverageVCLK1Frequency uint16
	// Average DCLK1 frequency (MHz)
	AverageDCLK1Frequency uint16

	// Current GFX clock frequency (MHz)
	CurrentGFXCLK uint16
	// Current SoC clock frequency (MHz)
	CurrentSoCCLK uint16
	// Current UCLK frequency (MHz)
	CurrentUCLK uint16
	// Current VCLK0 frequency (MHz)
	CurrentVCLK0 uint16
	// Current DCLK0 frequency (MHz)
	CurrentDCLK0 uint16
	// Current VCLK1 frequency (MHz)
	CurrentVCLK1 uint16
	// Current DCLK1 frequency (MHz)
	CurrentDCLK1 uint16

	// Current Throttle status
	ThrottleStatus uint32

	// Current fan speed (RPM)
	CurrentFanSpeed uint16

	// PCIe link width
	PCIeLinkWidth uint16
	// PCIe link speed (0.1 GT/s)
	PCIeLinkSpeed uint16

	// GFX accumulated activity
	GFXActivityAcc uint32
	// Memory accumulated activity
	MemActivityAcc uint32
	// HBM temperature (°C)
	TemperatureHBM [NUM_HBM_INSTANCE]uint16

	// Firmware timestamp (10 ns)
	FirmwareTimestamp uint64

	// SoC voltage (mV)
	VoltageSoC uint16
	// GFX voltage (mV)
	VoltageGFX uint16
	// Memory voltage (mV)
	VoltageMem uint16

	// ASIC independent throttle status
	IndepThrottleStatus uint64

	// Current socket power (W)
	CurrentSocketPower uint16

	// GFX clock lock status
	GFXCLKLockStatus uint32

	// XGMI link width
	XGMILinkWidth uint16
	// XGMI link speed (GB/s)
	XGMILinkSpeed uint16

	// PCIe accumulated bandwidth (GB/s)
	PCIeBandwidthAcc uint64

	// PCIe instantaneous bandwidth (GB/s)
	PCIeBandwidthInst uint64

	// PCIe L0 to recovery state transition accumulated count
	PCIeL0ToRecovCountAcc uint64

	// PCIe replay accumulated count
	PCIeReplayCountAcc uint64

	// PCIe replay rollover accumulated count
	PCIeReplayRoverCountAcc uint64

	// XGMI accumulated read data size (KB)
	XGMIReadDataAcc [MAX_NUM_XGMI_LINKS]uint64
	// XGMI accumulated write data size (KB)
	XGMIWriteDataAcc [MAX_NUM_XGMI_LINKS]uint64

	// Current GFX clock frequencies (MHz)
	CurrentGFXCLKs [MAX_NUM_GFX_CLKS]uint16
	// Current SoC clock frequencies (MHz)
	CurrentSoCCLKs [MAX_NUM_CLKS]uint16
	// Current VCLK0 frequencies (MHz)
	CurrentVCLK0s [MAX_NUM_CLKS]uint16
	// Current DCLK0 frequencies (MHz)
	CurrentDCLK0s [MAX_NUM_CLKS]uint16

	// JPEG activity (% / AID)
	JPEGActivity [MAX_NUM_JPEG]uint16

	// PCIe NAK sent accumulated count
	PCIeNAKSentCountAcc uint32

	// PCIe NAK received accumulated count
	PCIeNAKRcvdCountAcc uint32
}

type Frequencies struct {
	HasDeepSleep bool
	NumSupported uint32
	Current      uint32
	Frequency    [MAX_NUM_FREQUENCIES]uint64
}

type PCIeBandwidth struct {
	// Transfer rates (T/s) that are possible
	TransferRate Frequencies
	Lanes        [MAX_NUM_FREQUENCIES]uint32
}
