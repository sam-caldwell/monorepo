//go:build windows && amd64
// +build windows,amd64

package ole

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"syscall"
	"time"
	"unsafe"
)

// GetVariantDate converts COM Variant Time value to Go time.Time.
func GetVariantDate(value uint64) (time.Time, error) {
	var st syscall.Systemtime
	r, _, _ := procVariantTimeToSystemTime.Call(uintptr(value), uintptr(unsafe.Pointer(&st)))
	if r != 0 {
		return time.Date(int(st.Year), time.Month(st.Month), int(st.Day), int(st.Hour), int(st.Minute), int(st.Second), int(st.Milliseconds/1000), time.UTC), nil
	}
	return time.Now(), fmt.Errorf(errors.TimeConversionError)
}
