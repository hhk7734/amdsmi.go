package dl

// #cgo LDFLAGS: -ldl
// #include <stdlib.h>
// #include <dlfcn.h>
import "C"
import (
	"fmt"
	"unsafe"
)

const (
	RTLD_LAZY     = C.RTLD_LAZY
	RTLD_NOW      = C.RTLD_NOW
	RTLD_GLOBAL   = C.RTLD_GLOBAL
	RTLD_LOCAL    = C.RTLD_LOCAL
	RTLD_NODELETE = C.RTLD_NODELETE
	RTLD_NOLOAD   = C.RTLD_NOLOAD
)

type DynamicLibrary struct {
	handle unsafe.Pointer
}

func (d *DynamicLibrary) Open(name string, flags int) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	handle := C.dlopen(cname, C.int(flags))
	if handle == nil {
		return dlError()
	}
	d.handle = handle

	return nil
}

func (d *DynamicLibrary) Close() error {
	if d.handle == nil {
		return nil
	}

	if C.dlclose(d.handle) != 0 {
		return dlError()
	}
	d.handle = nil

	return nil
}

func dlError() error {
	if err := C.dlerror(); err != nil {
		return fmt.Errorf("dl: %s", C.GoString(err))
	}
	return nil
}
