package rsmi

import "fmt"

type (
	temperatureMetric = rsmi_temperature_metric
	temperatureType   = rsmi_temperature_type
	memoryType        = rsmi_memory_type
)

func (s rsmi_status) Err() error {
	switch s {
	case STATUS_SUCCESS:
		return nil
	default:
		return fmt.Errorf("unknow status %d", s)
	}
}
