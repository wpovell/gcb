package wifi

// #include "ssid.h"
import "C"
import (
	"errors"
	"unsafe"
)

// Get current SSID for interface
// Returns an error if interface is down or if err occurs
func ssid(intf string) (string, error) {
	// Convert interface to CString
	cIntf := C.CString(intf)
	defer C.free(unsafe.Pointer(cIntf))

	// Get SSID
	cSsid := C.ssid(cIntf)
	defer C.free(unsafe.Pointer(cSsid))
	if cSsid == nil {
		return "", errors.New("ssid returned NULL")
	}
	ssid := C.GoString(cSsid)

	return ssid, nil
}
