package logger

import (
	"bufio"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
)

func (log *Logger) ToFile(fileName string) *Logger {
	if file.Exists(fileName) {
		log.target, log.err = os.Open(fileName)
	} else {
		log.target, log.err = os.Create(fileName)
	}
	log.buffer = bufio.NewWriterSize(log.target, log.bufferSize)
	return log
}
