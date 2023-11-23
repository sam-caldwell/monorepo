package findCollision

import (
    "bufio"
    "encoding/hex"
    "github.com/sam-caldwell/monorepo/go/counters"
    "github.com/sam-caldwell/monorepo/go/fs/file"
    "github.com/sam-caldwell/monorepo/go/types/hashes"
    "log"
    "math"
    "os"
    "sync"
    "time"
    "unsafe"
)

const (
    //writeBufferSize = 4294967296 // 4GB Buffer for writes
    writeBufferSize = 1048576 * 20 // 10MB Buffer for writes
    hashFileName    = "/tmp/PreComputedHashes.txt"
)

var mutex sync.Mutex
var lookupTime time.Duration
var lookupCount int64
var lookupTotalTime time.Duration

func NewQuickTable(keySpaceSize, TableSize int) (t *QuickTable, lastSequence []byte) {
    var cycleStart time.Time
    var pos int
    var mode string
    var tableReady bool
    var table QuickTable
    table.Init()

    c, _ := counters.NewByteCounter(keySpaceSize)
    generatorStart := time.Now()
    defer func() { tableReady = true }()
    go func() {
        t := time.NewTicker(1 * time.Second)
        defer t.Stop()
        for !tableReady {
            select {
            case <-t.C:
                headerSize := int64(unsafe.Sizeof(table.data))
                recordSize := int64(unsafe.Sizeof(true) + unsafe.Sizeof(hashes.Sha1{}))
                recordCount := int64(len(table.data))
                size := headerSize + recordSize*recordCount
                var szUom string
                if size > 1073741824 {
                    szUom = "GB"
                    size = size / 1073741824
                } else if size > 1048576 {
                    szUom = "MB"
                    size = size / 1048576
                } else if size > 1024 {
                    szUom = "KB"
                    size = size / 1024
                } else {
                    szUom = "bytes"
                }

                avgLookupTime := time.Duration(float64(lookupTotalTime.Nanoseconds()) / float64(lookupCount))

                progress := 100 * float64(pos) / float64(TableSize)
                log.Printf("lookup table init (mode:%6s). "+
                    "(progress %12d/%012d %8.4f%%) (t/op:%6dns) elapsed:%6.2fs "+
                    "tableSize: %3d %s (lookup time: %4dns avg: %4dns)",
                    mode, pos, TableSize, progress, time.Since(cycleStart).Nanoseconds(),
                    time.Since(generatorStart).Seconds(), size, szUom,
                    lookupTime.Nanoseconds(), avgLookupTime.Nanoseconds())
            }
        }
    }()

    if file.Exists(hashFileName) {
        /*
         * The file exists, load it into memory
         */
        mode = "load"

        fileHandle, err := os.Open(hashFileName)
        if err != nil {
            panic(err)
        }
        defer func() { _ = fileHandle.Close() }()
        scanner := bufio.NewScanner(fileHandle)
        pos = 0
        TableSize = func() int {
            info, err := fileHandle.Stat()
            if err != nil {
                panic(err)
            }
            sz := info.Size()
            if sz > math.MaxInt {
                panic("pre-computed table is too large (past MaxInt)")
            }
            return int(sz)
        }()
        for scanner.Scan() {
            cycleStart = time.Now()
            line := scanner.Text()
            hash, err := hex.DecodeString(line)
            if err != nil {
                panic(err)
            }
            mutex.Lock()
            table.Store([20]byte(hash))
            mutex.Unlock()
            lastSequence = hash
            pos++
        }
        if err := scanner.Err(); err != nil {
            panic(err)
        }
    } else {
        /*
         * Create the new hash file for future and present use.
         */
        mode = "create"
        fileHandle, err := os.Create(hashFileName)
        if err != nil {
            panic(err)
        }
        writer := bufio.NewWriterSize(fileHandle, writeBufferSize)
        defer func() {
            _ = writer.Flush()
            _ = fileHandle.Close()
        }()
        for i := 0; i < TableSize; i++ {
            mutex.Lock()
            cycleStart = time.Now()
            hash := c.Sha1Bytes()
            table.Store(hash)
            if !table.Lookup(hash) {
                panic("Failed to look-up table after store")
            }
            if _, err := writer.WriteString(hex.EncodeToString(hash[:]) + "\n"); err != nil {
                panic(err)
            }
            _ = c.Increment()
            pos = i
            mutex.Unlock()
        }
    }
    lastSequence = c.Bytes()
    return &table, lastSequence
}

//type QuickMap map[byte]QuickMap

type QuickTable struct {
    data map[string]bool
}

func (o *QuickTable) Init() {
    o.data = make(map[string]bool)
}

func (o *QuickTable) Store(s [20]byte) {
    o.data[string(s[:])] = true
}

func (o *QuickTable) Lookup(s [20]byte) bool {
    lookupStart := time.Now()
    lookupCount++
    defer func() {
        lookupTime = time.Since(lookupStart)
        lookupTotalTime = lookupTotalTime + lookupTime
    }()
    _, ok := o.data[string(s[:])]
    return ok
}
