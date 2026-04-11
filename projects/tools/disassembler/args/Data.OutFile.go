package args

import "os"

func (data *Data) OutFile() *os.File {
	return data.out
}
