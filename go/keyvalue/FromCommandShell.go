package keyvalue

import (
	"os/exec"
)

// FromCommandShell - Execute a command shell and parse its results into the Key-Value Map.
//
//	     Note that this only supports KeyValue[string,string] and may produce unintended effects
//	     if used otherwise.
//
//		(c) 2023 Sam Caldwell.  MIT License
func (kv *KeyValue[KeyType, ValueType]) FromCommandShell(columnDelimiter, lineDelimiter, command string, args ...string) (err error) {
	var out []byte
	if out, err = exec.Command(command, args...).Output(); err != nil {
		return err
	}
	return kv.FromBytes(&out, lineDelimiter, columnDelimiter)
}
