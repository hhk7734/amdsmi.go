package amdsmi

/*
#include "amd_smi/amdsmi.h"
*/
import "C"
import (
	"errors"
	"fmt"
)

type initFlag uint32

const (
	INIT_ALL_PROCESSORS initFlag = C.AMDSMI_INIT_ALL_PROCESSORS
	INIT_AMD_CPUS       initFlag = C.AMDSMI_INIT_AMD_CPUS
	INIT_AMD_GPUS       initFlag = C.AMDSMI_INIT_AMD_GPUS
	INIT_NON_AMD_CPUS   initFlag = C.AMDSMI_INIT_NON_AMD_CPUS
	INIT_NON_AMD_GPUS   initFlag = C.AMDSMI_INIT_NON_AMD_GPUS
	INIT_AMD_APUS       initFlag = C.AMDSMI_INIT_AMD_APUS
)

type processorType uint32

const (
	UNKNOWN      processorType = C.UNKNOWN
	AMD_GPU      processorType = C.AMD_GPU
	AMD_CPU      processorType = C.AMD_CPU
	NON_AMD_GPU  processorType = C.NON_AMD_GPU
	NON_AMD_CPU  processorType = C.NON_AMD_CPU
	AMD_CPU_CORE processorType = C.AMD_CPU_CORE
	AMD_APU      processorType = C.AMD_APU
)

func (t processorType) String() string {
	switch t {
	case AMD_GPU:
		return "AMD GPU"
	case AMD_CPU:
		return "AMD CPU"
	case NON_AMD_GPU:
		return "Non AMD GPU"
	case NON_AMD_CPU:
		return "Non AMD CPU"
	case AMD_CPU_CORE:
		return "AMD CPU Core"
	case AMD_APU:
		return "AMD APU"
	default:
		return "Unknown"
	}
}

type amdsmiStatus uint32

var (
	ErrInvalidArgs          = errors.New("invalid arguments")
	ErrUnsupported          = errors.New("operation not supported")
	ErrNotImplemented       = errors.New("feature not yet implemented")
	ErrModuleLoadFailed     = errors.New("failed to load module")
	ErrSymbolLoadFailed     = errors.New("failed to load symbol")
	ErrDRMCallFailed        = errors.New("error occurred when calling libdrm")
	ErrAPICallFailed        = errors.New("API call failed")
	ErrAPITimeout           = errors.New("API call timed out")
	ErrRetryOperation       = errors.New("operation requires retry")
	ErrPermissionDenied     = errors.New("permission denied")
	ErrOperationInterrupted = errors.New("operation interrupted")
	ErrIOError              = errors.New("I/O error")
	ErrBadAddress           = errors.New("invalid address")
	ErrFileAccessFailed     = errors.New("failed to access file")
	ErrOutOfMemory          = errors.New("insufficient memory")
	ErrInternalException    = errors.New("internal exception occurred")
	ErrOutOfBounds          = errors.New("input out of bounds")
	ErrInitializationFailed = errors.New("initialization failed")
	ErrRefCountOverflow     = errors.New("reference count overflow")
	ErrDeviceBusy           = errors.New("device is busy")
	ErrDeviceNotFound       = errors.New("device not found")
	ErrDeviceNotInitialized = errors.New("device not initialized")
	ErrNoFreeSlot           = errors.New("no free slot available")
	ErrDriverNotLoaded      = errors.New("driver not loaded")
	ErrNoData               = errors.New("no data available")
	ErrInsufficientSize     = errors.New("insufficient size")
	ErrUnexpectedSize       = errors.New("unexpected size")
	ErrUnexpectedData       = errors.New("unexpected data")
	ErrNonAMDCPU            = errors.New("non-AMD CPU detected")
	ErrNoEnergyDriver       = errors.New("no energy driver found")
	ErrNoMSRDriver          = errors.New("no MSR driver found")
	ErrNoHSMPDriver         = errors.New("no HSMP driver found")
	ErrHSMPNotSupported     = errors.New("HSMP not supported")
	ErrHSMPMsgNotSupported  = errors.New("HSMP message not supported")
	ErrHSMPTimeout          = errors.New("HSMP operation timed out")
	ErrDriverMissing        = errors.New("no energy and HSMP drivers present")
	ErrFileNotFound         = errors.New("file not found")
	ErrNilArgumentPtr       = errors.New("argument pointer is null")
	ErrAMDGPUResetFailed    = errors.New("AMD GPU reset failed")
	ErrSettingUnavailable   = errors.New("setting unavailable")
	ErrMappingError         = errors.New("internal error did not map to status code")
)

func (s amdsmiStatus) Err() error {
	switch s {
	case C.AMDSMI_STATUS_SUCCESS:
		return nil
	case C.AMDSMI_STATUS_INVAL:
		return ErrInvalidArgs
	case C.AMDSMI_STATUS_NOT_SUPPORTED:
		return ErrUnsupported
	case C.AMDSMI_STATUS_NOT_YET_IMPLEMENTED:
		return ErrNotImplemented
	case C.AMDSMI_STATUS_FAIL_LOAD_MODULE:
		return ErrModuleLoadFailed
	case C.AMDSMI_STATUS_FAIL_LOAD_SYMBOL:
		return ErrSymbolLoadFailed
	case C.AMDSMI_STATUS_DRM_ERROR:
		return ErrDRMCallFailed
	case C.AMDSMI_STATUS_API_FAILED:
		return ErrAPICallFailed
	case C.AMDSMI_STATUS_TIMEOUT:
		return ErrAPITimeout
	case C.AMDSMI_STATUS_RETRY:
		return ErrRetryOperation
	case C.AMDSMI_STATUS_NO_PERM:
		return ErrPermissionDenied
	case C.AMDSMI_STATUS_INTERRUPT:
		return ErrOperationInterrupted
	case C.AMDSMI_STATUS_IO:
		return ErrIOError
	case C.AMDSMI_STATUS_ADDRESS_FAULT:
		return ErrBadAddress
	case C.AMDSMI_STATUS_FILE_ERROR:
		return ErrFileAccessFailed
	case C.AMDSMI_STATUS_OUT_OF_RESOURCES:
		return ErrOutOfMemory
	case C.AMDSMI_STATUS_INTERNAL_EXCEPTION:
		return ErrInternalException
	case C.AMDSMI_STATUS_INPUT_OUT_OF_BOUNDS:
		return ErrOutOfBounds
	case C.AMDSMI_STATUS_INIT_ERROR:
		return ErrInitializationFailed
	case C.AMDSMI_STATUS_REFCOUNT_OVERFLOW:
		return ErrRefCountOverflow
	case C.AMDSMI_STATUS_BUSY:
		return ErrDeviceBusy
	case C.AMDSMI_STATUS_NOT_FOUND:
		return ErrDeviceNotFound
	case C.AMDSMI_STATUS_NOT_INIT:
		return ErrDeviceNotInitialized
	case C.AMDSMI_STATUS_NO_SLOT:
		return ErrNoFreeSlot
	case C.AMDSMI_STATUS_DRIVER_NOT_LOADED:
		return ErrDriverNotLoaded
	case C.AMDSMI_STATUS_NO_DATA:
		return ErrNoData
	case C.AMDSMI_STATUS_INSUFFICIENT_SIZE:
		return ErrInsufficientSize
	case C.AMDSMI_STATUS_UNEXPECTED_SIZE:
		return ErrUnexpectedSize
	case C.AMDSMI_STATUS_UNEXPECTED_DATA:
		return ErrUnexpectedData
	case C.AMDSMI_STATUS_NON_AMD_CPU:
		return ErrNonAMDCPU
	case C.AMDSMI_NO_ENERGY_DRV:
		return ErrNoEnergyDriver
	case C.AMDSMI_NO_MSR_DRV:
		return ErrNoMSRDriver
	case C.AMDSMI_NO_HSMP_DRV:
		return ErrNoHSMPDriver
	case C.AMDSMI_NO_HSMP_SUP:
		return ErrHSMPNotSupported
	case C.AMDSMI_NO_HSMP_MSG_SUP:
		return ErrHSMPMsgNotSupported
	case C.AMDSMI_HSMP_TIMEOUT:
		return ErrHSMPTimeout
	case C.AMDSMI_NO_DRV:
		return ErrDriverMissing
	case C.AMDSMI_FILE_NOT_FOUND:
		return ErrDeviceNotFound
	case C.AMDSMI_ARG_PTR_NULL:
		return ErrNilArgumentPtr
	case C.AMDSMI_STATUS_AMDGPU_RESTART_ERR:
		return ErrAMDGPUResetFailed
	case C.AMDSMI_STATUS_SETTING_UNAVAILABLE:
		return ErrSettingUnavailable
	case C.AMDSMI_STATUS_MAP_ERROR:
		return ErrMappingError
	case C.AMDSMI_STATUS_UNKNOWN_ERROR:
		fallthrough
	default:
		return fmt.Errorf("amdsmi: unknown status %d", s)
	}
}

type temperatureType uint32

const (
	TEMP_TYPE_EDGE     temperatureType = C.TEMPERATURE_TYPE_EDGE
	TEMP_TYPE_FIRST    temperatureType = C.TEMPERATURE_TYPE_FIRST
	TEMP_TYPE_HOTSPOT  temperatureType = C.TEMPERATURE_TYPE_HOTSPOT
	TEMP_TYPE_JUNCTION temperatureType = C.TEMPERATURE_TYPE_JUNCTION
	TEMP_TYPE_VRAM     temperatureType = C.TEMPERATURE_TYPE_VRAM
	TEMP_TYPE_HBM_0    temperatureType = C.TEMPERATURE_TYPE_HBM_0
	TEMP_TYPE_HBM_1    temperatureType = C.TEMPERATURE_TYPE_HBM_1
	TEMP_TYPE_HBM_2    temperatureType = C.TEMPERATURE_TYPE_HBM_2
	TEMP_TYPE_HBM_3    temperatureType = C.TEMPERATURE_TYPE_HBM_3
	TEMP_TYPE_PLX      temperatureType = C.TEMPERATURE_TYPE_PLX
	TEMP_TYPE__MAX     temperatureType = C.TEMPERATURE_TYPE__MAX
)

type temperatureMetric uint32

const (
	TEMP_CURRENT        temperatureMetric = C.AMDSMI_TEMP_CURRENT
	TEMP_FIRST          temperatureMetric = C.AMDSMI_TEMP_FIRST
	TEMP_MAX            temperatureMetric = C.AMDSMI_TEMP_MAX
	TEMP_MIN            temperatureMetric = C.AMDSMI_TEMP_MIN
	TEMP_MAX_HYST       temperatureMetric = C.AMDSMI_TEMP_MAX_HYST
	TEMP_MIN_HYST       temperatureMetric = C.AMDSMI_TEMP_MIN_HYST
	TEMP_CRITICAL       temperatureMetric = C.AMDSMI_TEMP_CRITICAL
	TEMP_CRITICAL_HYST  temperatureMetric = C.AMDSMI_TEMP_CRITICAL_HYST
	TEMP_EMERGENCY      temperatureMetric = C.AMDSMI_TEMP_EMERGENCY
	TEMP_EMERGENCY_HYST temperatureMetric = C.AMDSMI_TEMP_EMERGENCY_HYST
	TEMP_CRIT_MIN       temperatureMetric = C.AMDSMI_TEMP_CRIT_MIN
	TEMP_CRIT_MIN_HYST  temperatureMetric = C.AMDSMI_TEMP_CRIT_MIN_HYST
	TEMP_OFFSET         temperatureMetric = C.AMDSMI_TEMP_OFFSET
	TEMP_LOWEST         temperatureMetric = C.AMDSMI_TEMP_LOWEST
	TEMP_HIGHEST        temperatureMetric = C.AMDSMI_TEMP_HIGHEST
	TEMP_LAST           temperatureMetric = C.AMDSMI_TEMP_LAST
)

type memoryType uint32

const (
	MEM_TYPE_FIRST    memoryType = C.AMDSMI_MEM_TYPE_FIRST
	MEM_TYPE_VRAM     memoryType = C.AMDSMI_MEM_TYPE_VRAM
	MEM_TYPE_VIS_VRAM memoryType = C.AMDSMI_MEM_TYPE_VIS_VRAM
	MEM_TYPE_GTT      memoryType = C.AMDSMI_MEM_TYPE_GTT
	MEM_TYPE_LAST     memoryType = C.AMDSMI_MEM_TYPE_LAST
)
