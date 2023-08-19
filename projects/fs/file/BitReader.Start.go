package file

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"io"
	"log"
	"os"
	"sync"
)

const (
	FailedToOpen = 10
	ReadFailed   = 11
)

func (reader *BitReader) Close() {
	reader.done = true
	close(reader.buffer)
}

// Start - start reading the file bit-by-bit into the bitReader.buffer (channel).
func (reader *BitReader) Start(fileName string, wg *sync.WaitGroup) {

	// Make sure our BitReader is initialized.
	if (reader.buffer == nil) || (reader.readBuffer == nil) {
		ansi.Red().Println("BitReader not initialized").Fatal(exit.GeneralError)
	}

	// Open the file handle
	handle, err := os.Open(fileName)
	if err != nil {
		ansi.Red().Println(err.Error()).Fatal(FailedToOpen)
	}
	defer func() {
		_ = handle.Close()
	}()
	defer func() {
		reader.Close()
		wg.Done()
	}()
	// Iterate over the file, reading in a buffer full of bytes and for each byte
	// isolate each bit and push the bit to the buffer (channel)
	for {
		nBytesRead, err := handle.Read(reader.readBuffer)
		if err != nil && err != io.EOF {
			ansi.Red().Println(err.Error()).Fatal(ReadFailed)
		}
		if nBytesRead == 0 {
			break
		}
		for _, byteVal := range reader.readBuffer[:nBytesRead] {
			for i := uint(0); i < 8; i++ {
				bit := (byteVal & (1 << i)) != 0
				reader.buffer <- bit
				log.Printf("bit count: %02d value: %v\n", reader.count, bit)
				reader.count++
			}
		}
	}
}
