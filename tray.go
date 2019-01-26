package main

// #include <stdlib.h>
// extern int trayOpen(char* device);
import "C"
import (
	"fmt"
	"unsafe"
)

func isNativeTrayOpen(device string) (bool, error) {
	cs := C.CString(device)
	intRes := C.trayOpen(cs)
	C.free(unsafe.Pointer(cs))

	if intRes == 0 {
		return true, nil
	} else if intRes == 1 {
		return false, nil
	} else if intRes == 2 {
		return false, fmt.Errorf("device not found")
	}

	return false, fmt.Errorf("invalid response code")
}
