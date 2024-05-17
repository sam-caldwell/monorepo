package keyvalue

import (
	"os/exec"
)

// FromCommandShell - Execute a command shell and parse its results into the Key-Value Map.
func (kv *KeyValue[KeyType, ValueType]) FromCommandShell(columnDelimiter, lineDelimiter, command string, args ...string) (err error) {
	var out []byte
	if out, err = exec.Command(command, args...).Output(); err != nil {
		return err
	}
	kv.lock.Lock()
	defer kv.lock.Unlock()
	return kv.FromBytes(&out, lineDelimiter, columnDelimiter)
}
