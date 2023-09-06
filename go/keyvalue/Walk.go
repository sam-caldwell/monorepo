package keyvalue

import "fmt"

/*
 * projects/keyvalue/Walk.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the keyvalue.Walk() method which will
 * allow us to 'walk' over a key-value store and execute a given
 * function for each key-value pair.
 *
 * See OpSys.Network.software.Memory.Disk.Cpu.README.md
 */

func (kv *KeyValue) Walk(fn func(key string, value interface{}) error) error {
	if kv.data == nil {
		return fmt.Errorf("keyvalue uninitialized")
	}
	for key, value := range kv.data {
		if err := fn(key, value); err != nil {
			return err
		}
	}
	return nil
}
