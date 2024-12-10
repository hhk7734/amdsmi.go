package amdsmi

type MetricsTableHeader struct {
	StructureSize   uint16
	FormatRevision  uint8
	ContentRevision uint8
}

type GPUMetrics struct {
	CommonHeader MetricsTableHeader

	TemperatureEdge    uint16
	TemperatureHotspot uint16
	TemperatureMem     uint16
	TemperatureVRGFX   uint16
	TemperatureVRSoC   uint16
	TemperatureVRMem   uint16

	AverageGFXActivity uint16
	AverageUMCActivity uint16
	AverageMMActivity  uint16

	AverageSocketPower uint16
	EnergyAccumulator  uint64

	SystemClockCounter uint64

	AverageGFXCLKFrequency uint16
	AverageSoCCLKFrequency uint16
	AverageUCLKFrequency   uint16
	AverageVCLK0Frequency  uint16
	AverageDCLK0Frequency  uint16
	AverageVCLK1Frequency  uint16
	AverageDCLK1Frequency  uint16

	CurrentGFXCLK uint16
	CurrentSoCCLK uint16
	CurrentUCLK   uint16
	CurrentVCLK0  uint16
	CurrentDCLK0  uint16
	CurrentVCLK1  uint16
	CurrentDCLK1  uint16

	ThrottleStatus uint32

	CurrentFanSpeed uint16

	PCIELinkWidth uint16
	PCIELinkSpeed uint16

	GFXActivityAcc uint32
	MemActivityAcc uint32
	TemperatureHBM [NUM_HBM_INSTANCE]uint16

	FirmwareTimestamp uint64

	VoltageSoC uint16
	VoltageGFX uint16
	VoltageMem uint16

	IndepThrottleStatus uint64

	CurrentSocketPower uint16

	GFXCLKLockStatus uint32

	XGMILinkWidth uint16
	XGMILinkSpeed uint16

	PCIEBandwidthAcc uint64

	PCIEBandwidthInst uint64

	PCIEL0ToRecovCountAcc uint64

	PCIEReplayCountAcc uint64

	PCIEReplayRoverCountAcc uint64

	XGMIReadDataAcc  [MAX_NUM_XGMI_LINKS]uint64
	XGMIWriteDataAcc [MAX_NUM_XGMI_LINKS]uint64

	CurrentGFXCLKs [MAX_NUM_GFX_CLKS]uint16
	CurrentSoCCLKs [MAX_NUM_CLKS]uint16
	CurrentVCLK0s  [MAX_NUM_CLKS]uint16
	CurrentDCLK0s  [MAX_NUM_CLKS]uint16

	JPEGActivity [MAX_NUM_JPEG]uint16

	PCIENAKSentCountAcc uint32

	PCIENAKRcvdCountAcc uint32
}
