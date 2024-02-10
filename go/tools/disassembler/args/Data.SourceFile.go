package args

import "os"

func (data *Data) SourceFile() *os.File {
	return data.source
}
