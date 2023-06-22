//go:build windows
// +build windows

package wmiclient

import "github.com/go-ole/go-ole"

// isOleVariantPropertyNull - the name is pretty self-explanatory
func isOleVariantPropertyNull(property ole.VT) bool {
	return property == VtNull
}
