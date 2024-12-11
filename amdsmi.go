package amdsmi

/*
#cgo linux LDFLAGS: -Wl,--export-dynamic -Wl,--unresolved-symbols=ignore-in-object-files
#include "amd_smi/amdsmi.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"

	"github.com/hhk7734/amdsmi.go/pkg/dl"
)

type AMDSMI struct {
	dl *dl.DynamicLibrary
}

func New() *AMDSMI {
	return &AMDSMI{
		dl: &dl.DynamicLibrary{},
	}
}

func (a *AMDSMI) Init(flags initFlag) error {
	if err := a.dl.Open("libamd_smi.so", dl.RTLD_LAZY|dl.RTLD_GLOBAL); err != nil {
		return err
	}

	if err := amdsmiStatus(C.amdsmi_init(C.uint64_t(flags))).Err(); err != nil {
		return err
	}

	return nil
}

func (a *AMDSMI) Shutdown() error {
	if err := amdsmiStatus(C.amdsmi_shut_down()).Err(); err != nil {
		return err
	}

	if err := a.dl.Close(); err != nil {
		return err
	}

	return nil
}

type Socket struct {
	handle C.amdsmi_socket_handle
}

func (a *AMDSMI) Sockets() ([]*Socket, error) {
	var socketCount C.uint32_t
	if err := amdsmiStatus(C.amdsmi_get_socket_handles(
		&socketCount, nil)).Err(); err != nil {
		return nil, err
	}

	socketHandles := make([]C.amdsmi_socket_handle, socketCount)
	if err := amdsmiStatus(C.amdsmi_get_socket_handles(
		&socketCount, &socketHandles[0])).Err(); err != nil {
		return nil, err
	}

	sockets := make([]*Socket, socketCount)
	for i := 0; i < int(socketCount); i++ {
		sockets[i] = &Socket{
			handle: socketHandles[i],
		}
	}

	return sockets, nil
}

func (s *Socket) Info() (string, error) {
	info := make([]byte, 128)
	if err := amdsmiStatus(C.amdsmi_get_socket_info(
		s.handle, C.size_t(len(info)), (*C.char)(unsafe.Pointer(&info[0])))).Err(); err != nil {
		return "", err
	}

	return string(info), nil
}

type Processor struct {
	handle C.amdsmi_processor_handle
}

func (s *Socket) Processors() ([]*Processor, error) {
	var processorCount C.uint32_t
	if err := amdsmiStatus(C.amdsmi_get_processor_handles(
		s.handle, &processorCount, nil)).Err(); err != nil {
		return nil, err
	}

	processorHandles := make([]C.amdsmi_processor_handle, processorCount)
	if err := amdsmiStatus(C.amdsmi_get_processor_handles(
		s.handle, &processorCount, &processorHandles[0])).Err(); err != nil {
		return nil, err
	}

	processors := make([]*Processor, processorCount)
	for i := 0; i < int(processorCount); i++ {
		processors[i] = &Processor{
			handle: processorHandles[i],
		}
	}

	return processors, nil
}

func (p *Processor) Type() (processorType, error) {
	var type_ C.processor_type_t
	if err := amdsmiStatus(C.amdsmi_get_processor_type(
		p.handle, &type_)).Err(); err != nil {
		return UNKNOWN, err
	}

	return processorType(type_), nil
}

func (p *Processor) GPUID() (uint16, error) {
	var gpuID C.uint16_t
	if err := amdsmiStatus(C.amdsmi_get_gpu_id(
		p.handle, &gpuID)).Err(); err != nil {
		return 0, err
	}

	return uint16(gpuID), nil
}

func (p *Processor) GPURevision() (uint16, error) {
	var gpuRevision C.uint16_t
	if err := amdsmiStatus(C.amdsmi_get_gpu_revision(
		p.handle, &gpuRevision)).Err(); err != nil {
		return 0, err
	}

	return uint16(gpuRevision), nil
}

func (p *Processor) GPUVendorName() (string, error) {
	vendor := make([]byte, 128)
	if err := amdsmiStatus(C.amdsmi_get_gpu_vendor_name(
		p.handle, (*C.char)(unsafe.Pointer(&vendor[0])), C.size_t(len(vendor)))).Err(); err != nil {
		return "", err
	}

	return string(vendor), nil
}

func (p *Processor) GPUVRAMVendor() (string, error) {
	vramVendor := make([]byte, 128)
	if err := amdsmiStatus(C.amdsmi_get_gpu_vram_vendor(
		p.handle, (*C.char)(unsafe.Pointer(&vramVendor[0])), C.uint32_t(len(vramVendor)))).Err(); err != nil {
		return "", err
	}

	return string(vramVendor), nil
}

func (p *Processor) GPUSubsystemID() (uint16, error) {
	var subsystemID C.uint16_t
	if err := amdsmiStatus(C.amdsmi_get_gpu_subsystem_id(
		p.handle, &subsystemID)).Err(); err != nil {
		return 0, err
	}

	return uint16(subsystemID), nil
}

func (p *Processor) GPUSubsystemName() (string, error) {
	subsystem := make([]byte, 128)
	if err := amdsmiStatus(C.amdsmi_get_gpu_subsystem_name(
		p.handle, (*C.char)(unsafe.Pointer(&subsystem[0])), C.size_t(len(subsystem)))).Err(); err != nil {
		return "", err
	}

	return string(subsystem), nil
}

func (p *Processor) GPUPCIeBandwidth() (PCIeBandwidth, error) {
	var bandwidth C.amdsmi_pcie_bandwidth_t
	if err := amdsmiStatus(C.amdsmi_get_gpu_pci_bandwidth(
		p.handle, &bandwidth)).Err(); err != nil {
		return PCIeBandwidth{}, err
	}

	pcieBandwidth := PCIeBandwidth{
		TransferRate: Frequencies{
			HasDeepSleep: bool(bandwidth.transfer_rate.has_deep_sleep),
			NumSupported: uint32(bandwidth.transfer_rate.num_supported),
			Current:      uint32(bandwidth.transfer_rate.current),
			Frequency:    [MAX_NUM_FREQUENCIES]uint64{},
		},
		Lanes: [MAX_NUM_FREQUENCIES]uint32{},
	}

	for i := 0; i < MAX_NUM_FREQUENCIES; i++ {
		pcieBandwidth.TransferRate.Frequency[i] = uint64(bandwidth.transfer_rate.frequency[i])
		pcieBandwidth.Lanes[i] = uint32(bandwidth.lanes[i])
	}

	return pcieBandwidth, nil
}

func (p *Processor) GPUBDFID() (uint64, error) {
	var bdfid C.uint64_t
	if err := amdsmiStatus(C.amdsmi_get_gpu_bdf_id(
		p.handle, &bdfid)).Err(); err != nil {
		return 0, err
	}

	return uint64(bdfid), nil
}

// GPUMemoryTotal returns the total memory of type_ in bytes.
func (p *Processor) GPUMemoryTotal(type_ memoryType) (uint64, error) {
	var memoryTotal C.uint64_t
	if err := amdsmiStatus(C.amdsmi_get_gpu_memory_total(
		p.handle, C.amdsmi_memory_type_t(type_), &memoryTotal)).Err(); err != nil {
		return 0, err
	}

	return uint64(memoryTotal), nil
}

// GPUMemoryUsage returns the memory usage of type_ in bytes.
func (p *Processor) GPUMemoryUsage(type_ memoryType) (uint64, error) {
	var memoryUsed C.uint64_t
	if err := amdsmiStatus(C.amdsmi_get_gpu_memory_usage(
		p.handle, C.amdsmi_memory_type_t(type_), &memoryUsed)).Err(); err != nil {
		return 0, err
	}

	return uint64(memoryUsed), nil
}

func (p *Processor) GPUFanRPM(index uint32) (uint32, error) {
	var fanRPMs C.int64_t
	if err := amdsmiStatus(C.amdsmi_get_gpu_fan_rpms(
		p.handle, C.uint32_t(index), &fanRPMs)).Err(); err != nil {
		return 0, err
	}

	return uint32(fanRPMs), nil
}

func (p *Processor) GPUFanSpeed(index uint32) (uint32, error) {
	var fanSpeed C.int64_t
	if err := amdsmiStatus(C.amdsmi_get_gpu_fan_speed(
		p.handle, C.uint32_t(index), &fanSpeed)).Err(); err != nil {
		return 0, err
	}

	return uint32(fanSpeed), nil
}

func (p *Processor) GPUFanSpeedMax(index uint32) (uint32, error) {
	var fanSpeedMax C.uint64_t
	if err := amdsmiStatus(C.amdsmi_get_gpu_fan_speed_max(
		p.handle, C.uint32_t(index), &fanSpeedMax)).Err(); err != nil {
		return 0, err
	}

	return uint32(fanSpeedMax), nil
}

func (p *Processor) Temperature(type_ temperatureType, metric temperatureMetric) (int64, error) {
	var temp C.int64_t
	if err := amdsmiStatus(C.amdsmi_get_temp_metric(
		p.handle, C.amdsmi_temperature_type_t(type_), C.amdsmi_temperature_metric_t(metric), &temp)).Err(); err != nil {
		return 0, err
	}

	return int64(temp), nil
}

func (p *Processor) GPUMetricsInfo() (GPUMetrics, error) {
	var metrics C.amdsmi_gpu_metrics_t
	if err := amdsmiStatus(C.amdsmi_get_gpu_metrics_info(p.handle, &metrics)).Err(); err != nil {
		return GPUMetrics{}, err
	}

	gpuMetrics := GPUMetrics{
		CommonHeader: MetricsTableHeader{
			StructureSize:   uint16(metrics.common_header.structure_size),
			FormatRevision:  uint8(metrics.common_header.format_revision),
			ContentRevision: uint8(metrics.common_header.content_revision),
		},

		TemperatureEdge:    uint16(metrics.temperature_edge),
		TemperatureHotspot: uint16(metrics.temperature_hotspot),
		TemperatureMem:     uint16(metrics.temperature_mem),
		TemperatureVRGFX:   uint16(metrics.temperature_vrgfx),
		TemperatureVRSoC:   uint16(metrics.temperature_vrsoc),
		TemperatureVRMem:   uint16(metrics.temperature_vrmem),

		AverageGFXActivity: uint16(metrics.average_gfx_activity),
		AverageUMCActivity: uint16(metrics.average_umc_activity),
		AverageMMActivity:  uint16(metrics.average_mm_activity),

		AverageSocketPower: uint16(metrics.average_socket_power),
		EnergyAccumulator:  uint64(metrics.energy_accumulator),

		SystemClockCounter: uint64(metrics.system_clock_counter),

		AverageGFXCLKFrequency: uint16(metrics.average_gfxclk_frequency),
		AverageSoCCLKFrequency: uint16(metrics.average_socclk_frequency),
		AverageUCLKFrequency:   uint16(metrics.average_uclk_frequency),
		AverageVCLK0Frequency:  uint16(metrics.average_vclk0_frequency),
		AverageDCLK0Frequency:  uint16(metrics.average_dclk0_frequency),
		AverageVCLK1Frequency:  uint16(metrics.average_vclk1_frequency),
		AverageDCLK1Frequency:  uint16(metrics.average_dclk1_frequency),

		CurrentGFXCLK: uint16(metrics.current_gfxclk),
		CurrentSoCCLK: uint16(metrics.current_socclk),
		CurrentUCLK:   uint16(metrics.current_uclk),
		CurrentVCLK0:  uint16(metrics.current_vclk0),
		CurrentDCLK0:  uint16(metrics.current_dclk0),
		CurrentVCLK1:  uint16(metrics.current_vclk1),
		CurrentDCLK1:  uint16(metrics.current_dclk1),

		ThrottleStatus: uint32(metrics.throttle_status),

		CurrentFanSpeed: uint16(metrics.current_fan_speed),

		PCIeLinkWidth: uint16(metrics.pcie_link_width),
		PCIeLinkSpeed: uint16(metrics.pcie_link_speed),

		GFXActivityAcc: uint32(metrics.gfx_activity_acc),
		MemActivityAcc: uint32(metrics.mem_activity_acc),
		TemperatureHBM: [NUM_HBM_INSTANCE]uint16{},

		FirmwareTimestamp: uint64(metrics.firmware_timestamp),

		VoltageSoC: uint16(metrics.voltage_soc),
		VoltageGFX: uint16(metrics.voltage_gfx),
		VoltageMem: uint16(metrics.voltage_mem),

		IndepThrottleStatus: uint64(metrics.indep_throttle_status),

		CurrentSocketPower: uint16(metrics.current_socket_power),

		GFXCLKLockStatus: uint32(metrics.gfxclk_lock_status),

		XGMILinkWidth: uint16(metrics.xgmi_link_width),
		XGMILinkSpeed: uint16(metrics.xgmi_link_speed),

		PCIeBandwidthAcc: uint64(metrics.pcie_bandwidth_acc),

		PCIeBandwidthInst: uint64(metrics.pcie_bandwidth_inst),

		PCIeL0ToRecovCountAcc: uint64(metrics.pcie_l0_to_recov_count_acc),

		PCIeReplayCountAcc: uint64(metrics.pcie_replay_count_acc),

		PCIeReplayRoverCountAcc: uint64(metrics.pcie_replay_rover_count_acc),

		XGMIReadDataAcc:  [MAX_NUM_XGMI_LINKS]uint64{},
		XGMIWriteDataAcc: [MAX_NUM_XGMI_LINKS]uint64{},

		CurrentGFXCLKs: [MAX_NUM_GFX_CLKS]uint16{},
		CurrentSoCCLKs: [MAX_NUM_CLKS]uint16{},
		CurrentVCLK0s:  [MAX_NUM_CLKS]uint16{},
		CurrentDCLK0s:  [MAX_NUM_CLKS]uint16{},

		JPEGActivity: [MAX_NUM_JPEG]uint16{},

		PCIeNAKSentCountAcc: uint32(metrics.pcie_nak_sent_count_acc),
		PCIeNAKRcvdCountAcc: uint32(metrics.pcie_nak_rcvd_count_acc),
	}

	for i := 0; i < NUM_HBM_INSTANCE; i++ {
		gpuMetrics.TemperatureHBM[i] = uint16(metrics.temperature_hbm[i])
	}

	for i := 0; i < MAX_NUM_XGMI_LINKS; i++ {
		gpuMetrics.XGMIReadDataAcc[i] = uint64(metrics.xgmi_read_data_acc[i])
		gpuMetrics.XGMIWriteDataAcc[i] = uint64(metrics.xgmi_write_data_acc[i])
	}

	for i := 0; i < MAX_NUM_GFX_CLKS; i++ {
		gpuMetrics.CurrentGFXCLKs[i] = uint16(metrics.current_gfxclks[i])
	}

	for i := 0; i < MAX_NUM_CLKS; i++ {
		gpuMetrics.CurrentSoCCLKs[i] = uint16(metrics.current_socclks[i])
		gpuMetrics.CurrentVCLK0s[i] = uint16(metrics.current_vclk0s[i])
		gpuMetrics.CurrentDCLK0s[i] = uint16(metrics.current_dclk0s[i])
	}

	for i := 0; i < MAX_NUM_JPEG; i++ {
		gpuMetrics.JPEGActivity[i] = uint16(metrics.jpeg_activity[i])
	}

	return gpuMetrics, nil
}
