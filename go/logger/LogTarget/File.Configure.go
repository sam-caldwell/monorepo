package LogTarget

import (
	"os"
	"strconv"
)

func (out File) Configure(cfg ConfigureTarget) {
	var (
		ok         bool
		filename   string
		permission os.FileMode
	)
	if filename, ok = cfg["filename"]; !ok {
		panic("missing filename in logger config")
	}
	permission = func() (perm os.FileMode) {
		var permissionString string
		var permissionNumber int
		var err error
		if permissionString, ok = cfg["permission"]; ok {
			panic("missing permissions in logger config")
		}
		if permissionNumber, err = strconv.Atoi(permissionString); err != nil {
			panic(err)
		}
		return os.FileMode(permissionNumber)
	}()

	if err := out.file.Open(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, permission); err != nil {
		panic(err)
	}
}
