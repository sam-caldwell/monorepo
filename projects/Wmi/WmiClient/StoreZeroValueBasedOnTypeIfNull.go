//go:build windows
// +build windows

package wmiclient

import (
	"fmt"
	"reflect"
)

// StoreZeroValueBasedOnTypeIfNull - if the field or property are null, etc., store a safe type-specific zero value
func StoreZeroValueBasedOnTypeIfNull(isPointer bool, client *Client, originalFieldValue reflect.Value) error {
	if (isPointer && client.PtrNil) || (!isPointer && client.NonePtrZero) {
		originalFieldValue.Set(reflect.Zero(originalFieldValue.Type()))
	}
	return fmt.Errorf("break")
}
