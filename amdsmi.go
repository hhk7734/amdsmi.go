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

type socket struct {
	handle C.amdsmi_socket_handle
}

func (a *AMDSMI) Sockets() ([]*socket, error) {
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

	sockets := make([]*socket, socketCount)
	for i := 0; i < int(socketCount); i++ {
		sockets[i] = &socket{
			handle: socketHandles[i],
		}
	}

	return sockets, nil
}

func (s *socket) Info() (string, error) {
	info := make([]byte, 128)
	if err := amdsmiStatus(C.amdsmi_get_socket_info(
		s.handle, C.size_t(len(info)), (*C.char)(unsafe.Pointer(&info[0])))).Err(); err != nil {
		return "", err
	}

	return string(info), nil
}

type processor struct {
	handle C.amdsmi_processor_handle
}

func (s *socket) Processors() ([]*processor, error) {
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

	processors := make([]*processor, processorCount)
	for i := 0; i < int(processorCount); i++ {
		processors[i] = &processor{
			handle: processorHandles[i],
		}
	}

	return processors, nil
}

func (p *processor) Type() (processorType, error) {
	var type_ C.processor_type_t
	if err := amdsmiStatus(C.amdsmi_get_processor_type(
		p.handle, &type_)).Err(); err != nil {
		return UNKNOWN, err
	}

	return processorType(type_), nil
}

func (p *processor) GPUID() (uint16, error) {
	var gpuID C.uint16_t
	if err := amdsmiStatus(C.amdsmi_get_gpu_id(
		p.handle, &gpuID)).Err(); err != nil {
		return 0, err
	}

	return uint16(gpuID), nil
}

func (p *processor) Temperature(type_ temperatureType, metric temperatureMetric) (int64, error) {
	var temp C.int64_t
	if err := amdsmiStatus(C.amdsmi_get_temp_metric(
		p.handle, C.amdsmi_temperature_type_t(type_), C.amdsmi_temperature_metric_t(metric), &temp)).Err(); err != nil {
		return 0, err
	}

	return int64(temp), nil
}
