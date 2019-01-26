package main

// #include <stdlib.h>
// extern int trayOpen(char* device);
import "C"
import "unsafe"

func isNativeTrayOpen(device string) bool {
	cs := C.CString(device)
	intRes := C.trayOpen(cs)
	C.free(unsafe.Pointer(cs))
	return intRes == 0
}
