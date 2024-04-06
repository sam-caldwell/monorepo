package types

import (
	"fmt"
	"strings"
)

// OperatingMode - The operating mode: client vs server
type OperatingMode uint8

const (
	Undefined OperatingMode = iota
	Client

	Server
)

// FromString - Convert string to operating mode
func (o *OperatingMode) FromString(v string) error {
	switch strings.ToLower(v) {
	case "client":
		*o = Client
	case "server":
		*o = Server
	default:
		return fmt.Errorf("unexpected operating mode: %s", v)
	}
	return nil
}

// ToString - Convert OperatingMode to string
func (o *OperatingMode) ToString() string {
	switch *o {
	case Client:
		return "client"
	case Server:
		return "server"
	default:
		panic("invalid operating mode")
	}
}
