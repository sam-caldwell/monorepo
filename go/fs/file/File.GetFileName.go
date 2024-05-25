package file

import "github.com/sam-caldwell/monorepo/go/misc/words"

// GetFileName - Return the internal file/path state
//
//	     The goal is to create an abstraction layer so changes to golang will not require
//	     significant effort to update all projects.
//
//			(c) 2023 Sam Caldwell.  MIT License
func (fp *File) GetFileName() string {
	if fp.handle == nil {
		return words.EmptyString
	}
	fp.lock.Lock()
	defer fp.lock.Unlock()
	return fp.handle.Name()
}
