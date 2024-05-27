package file

import (
    "github.com/google/uuid"
    "github.com/sam-caldwell/monorepo/go/exit/errors"
    "os"
    "path/filepath"
    "testing"
)

func TestFile_WriteStringp(t *testing.T) {
    t.Run("test uninitialized state", func(t *testing.T) {
        var f File
        n, err := f.WriteString("test")
        if err == nil {
            t.Fatal("expected error when writing to uninitialized file (nil file handle)")
        }
        if err.Error() != errors.NotInitialized {
            t.Fatal("expected NotInitialized error")
        }
        if n != 0 {
            t.Errorf("WriteBytes failed: expected 0, got %v", n)
        }
    })

    t.Run("test initialized state", func(t *testing.T) {
        var (
            f                File
            name             string
            expectedFileName = filepath.Join("/tmp/", uuid.New().String())
        )

        t.Cleanup(func() {
            _ = f.handle.Close()
            _ = os.Remove(expectedFileName)
        })

        t.Run("test create test file", func(t *testing.T) {
            var err error
            if err = f.Open(expectedFileName, os.O_CREATE|os.O_RDWR, 0600); err != nil {
                t.Fatal(err)
            }
            if f.handle == nil {
                t.Fatal("handle is nil")
            }
            name = f.handle.Name()
            if name != expectedFileName {
                t.Fatal("name mismatch")
            }
        })

        t.Run("test that the file exists", func(t *testing.T) {
            if !Existsp(&name) {
                t.Fatal("file does not exist as expected")
            }
        })

        t.Run("write file bytes", func(t *testing.T) {
            var testData = "test1"
            n, err := f.WriteStringp(&testData)
            if err != nil {
                t.Fatal(err)
            }
            if n != len(testData) {
                t.Fatal("size mismatch")
            }
            pos, _ := f.handle.Seek(0, SeekFromCurrent)
            if pos != int64(len(testData)) {
                t.Fatal("position mismatch")
            }
            b := make([]byte, len(testData))
            n, err = f.handle.ReadAt(b, 0)
            if err != nil {
                t.Fatal(err)
            }
            if n != len(testData) {
                t.Fatal("read size mismatch")
            }
            if string(b[:n]) != testData {
                t.Fatal("read data mismatch")
            }
        })
    })
}
