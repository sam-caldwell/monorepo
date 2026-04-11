package file

// Close - if file handle not nil, close the file handle
//
//	     The goal is to create an abstraction layer so changes to golang will not require
//	     significant effort to update all projects.
//
//			(c) 2023 Sam Caldwell.  MIT License
func (fp *File) Close() {
	fp.lock.Lock()
	defer fp.lock.Unlock()
	if fp.handle != nil {
		_ = fp.handle.Close()
		fp.handle = nil
	}
}
