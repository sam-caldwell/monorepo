package file

import "io"

const (
	/*
	   This set of constants creates an abstraction from the continuous breaking changes
	   of golang's evolution.  Not long ago os.SEEK_CUR, etc. moved to io.SeekStart, et seq.

	   This guarantees that a future change will only require me to make one change here to
	   fix all the things that use this package...because I am lazy.
	*/

	//SeekFromStart - Abstraction from golang constants (see notes, above)
	SeekFromStart = io.SeekStart

	//SeekFromCurrent - Abstraction from golang constants (see notes, above)
	SeekFromCurrent = io.SeekCurrent

	//SeekFromEnd - Abstraction from golang constants (see notes, above)
	SeekFromEnd = io.SeekEnd
)
