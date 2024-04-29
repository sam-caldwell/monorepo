package persistence

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/types"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestEventWrite(t *testing.T) {
	const (
		tmpDir = "/tmp"
	)

	var e EventCollector
	var dir directory.Path
	dir.Set(tmpDir)

	eventTime := time.Now()
	sendTime := time.Now()
	testField := types.FieldId(0xFF)
	testResource := types.ResourceId(0xEA)

	evt := types.Event{
		EventTime: eventTime,
		Field:     testField,
		SendTime:  sendTime,
		Resource:  testResource,
		Value:     []byte("event value"),
	}

	t.Run("initialize EventCollector (happy)", func(t *testing.T) {
		if err := e.Init(&dir); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("test the write method", func(t *testing.T) {
		if err := e.Write(evt); err != nil {
			t.Fatal(err)
		}
	})
	fileName := filepath.Join(e.path.String(), fmt.Sprintf("event-1"))
	t.Run("verify the event file is written", func(t *testing.T) {
		if !file.Exists(fileName) {
			t.Fatalf("file not found (%s)", fileName)
		}
	})
	t.Run("verify the event file content", func(t *testing.T) {
		var (
			err     error
			content []byte
		)
		if content, err = os.ReadFile(fileName); err != nil {
			t.Fatal(err)
		}
		t.Log(content)
	})

}
