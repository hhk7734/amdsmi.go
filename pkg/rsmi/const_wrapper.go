package rsmi

import (
	"errors"
)

type (
	temperatureMetric = rsmi_temperature_metric
	temperatureType   = rsmi_temperature_type
	memoryType        = rsmi_memory_type
)

var (
	ErrInvalidArgs        = errors.New("invalid arguments")
	ErrNotSupported       = errors.New("not supported")
	ErrFileError          = errors.New("file error")
	ErrPermission         = errors.New("permission denied")
	ErrOutOfResources     = errors.New("out of resources")
	ErrInternalException  = errors.New("internal exception")
	ErrInputOutOfBounds   = errors.New("input out of bounds")
	ErrInitError          = errors.New("initialization error")
	ErrNotYetImplemented  = errors.New("not yet implemented")
	ErrNotFound           = errors.New("not found")
	ErrInsufficientSize   = errors.New("insufficient size")
	ErrInterrupt          = errors.New("interrupt")
	ErrUnexpectedSize     = errors.New("unexpected size")
	ErrNoData             = errors.New("no data")
	ErrUnexpectedData     = errors.New("unexpected data")
	ErrBusy               = errors.New("busy")
	ErrRefcountOverflow   = errors.New("refcount overflow")
	ErrSettingUnavailable = errors.New("setting unavailable")
	ErrAMDGPURestart      = errors.New("amdgpu restart error")
	ErrUnknown            = errors.New("unknown error")
)

func (s rsmi_status) Err() error {
	switch s {
	case STATUS_SUCCESS:
		return nil
	case STATUS_INVALID_ARGS:
		return ErrInvalidArgs
	case STATUS_NOT_SUPPORTED:
		return ErrNotSupported
	case STATUS_FILE_ERROR:
		return ErrFileError
	case STATUS_PERMISSION:
		return ErrPermission
	case STATUS_OUT_OF_RESOURCES:
		return ErrOutOfResources
	case STATUS_INTERNAL_EXCEPTION:
		return ErrInternalException
	case STATUS_INPUT_OUT_OF_BOUNDS:
		return ErrInputOutOfBounds
	case STATUS_INIT_ERROR:
		return ErrInitError
	case STATUS_NOT_YET_IMPLEMENTED:
		return ErrNotYetImplemented
	case STATUS_NOT_FOUND:
		return ErrNotFound
	case STATUS_INSUFFICIENT_SIZE:
		return ErrInsufficientSize
	case STATUS_INTERRUPT:
		return ErrInterrupt
	case STATUS_UNEXPECTED_SIZE:
		return ErrUnexpectedSize
	case STATUS_NO_DATA:
		return ErrNoData
	case STATUS_UNEXPECTED_DATA:
		return ErrUnexpectedData
	case STATUS_BUSY:
		return ErrBusy
	case STATUS_REFCOUNT_OVERFLOW:
		return ErrRefcountOverflow
	case STATUS_SETTING_UNAVAILABLE:
		return ErrSettingUnavailable
	case STATUS_AMDGPU_RESTART_ERR:
		return ErrAMDGPURestart
	default:
		return ErrUnknown
	}
}
