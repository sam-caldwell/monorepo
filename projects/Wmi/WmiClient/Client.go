//go:build windows
// +build windows

package wmiclient

import (
	"github.com/sam-caldwell/go/v2/projects/Wmi/swbemservices"
	"sync"
)

// Client An WMI query client.
//
// Its zero value (defaultClient) is a usable client.
type Client struct {
	lock sync.Mutex
	// NonePtrZero specifies if nil values for fields which aren't pointers
	// should be returned as the field types zero value.
	//
	// Setting this to true allows structs without pointer fields to be used
	// without the risk failure should a nil value returned from WMI.
	NonePtrZero bool

	// PtrNil specifies if nil values for pointer fields should be returned
	// as nil.
	//
	// Setting this to true will set pointer fields to nil where WMI
	// returned nil, otherwise the types zero value will be returned.
	PtrNil bool

	// AllowMissingFields specifies that struct fields not present in the
	// query result should not result in an error.
	//
	// Setting this to true allows custom queries to be used with full
	// struct definitions instead of having to define multiple structs.
	AllowMissingFields bool

	// SWbemServiceClient is an optional SWbemServices object that can be
	// initialized and then reused across multiple queries. If it is null
	// then the method will initialize a new temporary client each time.
	SWbemServicesClient *swbemservices.SWbemServices
}

func (client *Client) Lock() {
	client.lock.Lock()
}
func (client *Client) Unlock() {
	defer client.lock.Unlock()
}
