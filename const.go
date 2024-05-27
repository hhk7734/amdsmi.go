package amdsmi

/*
#include "amd_smi/amdsmi.h"
*/
import "C"
import "fmt"

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

func (s amdsmiStatus) Err() error {
	switch s {
	case C.AMDSMI_STATUS_SUCCESS:
		return nil
	default:
		return fmt.Errorf("amdsmi: unknown status %d", s)
	}
}

type temperatureType uint32

const (
	TEMP_TYPE_EDGE     temperatureType = C.TEMPERATURE_TYPE_EDGE
	TEMP_TYPE_FIRST    temperatureType = TEMP_TYPE_EDGE
	TEMP_TYPE_HOTSPOT  temperatureType = C.TEMPERATURE_TYPE_HOTSPOT
	TEMP_TYPE_JUNCTION temperatureType = TEMP_TYPE_HOTSPOT
	TEMP_TYPE_VRAM     temperatureType = C.TEMPERATURE_TYPE_VRAM
	TEMP_TYPE_HBM_0    temperatureType = C.TEMPERATURE_TYPE_HBM_0
	TEMP_TYPE_HBM_1    temperatureType = C.TEMPERATURE_TYPE_HBM_1
	TEMP_TYPE_HBM_2    temperatureType = C.TEMPERATURE_TYPE_HBM_2
	TEMP_TYPE_HBM_3    temperatureType = C.TEMPERATURE_TYPE_HBM_3
	TEMP_TYPE_PLX      temperatureType = C.TEMPERATURE_TYPE_PLX
	TEMP_TYPE__MAX     temperatureType = TEMP_TYPE_PLX
)

type temperatureMetric uint32

const (
	TEMP_CURRENT        temperatureMetric = C.AMDSMI_TEMP_CURRENT
	TEMP_FIRST          temperatureMetric = TEMP_CURRENT
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
	TEMP_LAST           temperatureMetric = TEMP_HIGHEST
)
