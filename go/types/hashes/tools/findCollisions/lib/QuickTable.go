package findCollision

import (
    "bufio"
    "compress/gzip"
    "encoding/hex"
    "github.com/sam-caldwell/monorepo/go/counters"
    "github.com/sam-caldwell/monorepo/go/fs/file"
    "github.com/sam-caldwell/monorepo/go/types/hashes"
    "log"
    "math"
    "os"
    "time"
    "unsafe"
)

const (
    //writeBufferSize = 4294967296 // 4GB Buffer for writes
    writeBufferSize = 1048576 * 20 // 10MB Buffer for writes
)

var storeTime time.Duration
var storeCount int64
var storeTotalTime time.Duration
var lookupTime time.Duration
var lookupCount int64
var lookupTotalTime time.Duration

func sizeFix(sz int64) (int64, string) {
    if sz > 1073741824 {
        return sz / 1073741824, "GB"
    } else if sz > 1048576 {
        return sz / 1048576, "MB"
    } else if sz > 1024 {
        return sz / 1024, "KB"
    } else {
        return sz, "bytes"
    }
}

func logMessage(
    mode *string,
    headerSize, recordSize, recordCount, pos, TableSize int64,
    cycleStart, generatorStart time.Time, lookupTime, storeTime time.Duration) {

    size, szUom := sizeFix(headerSize + recordSize*recordCount)

    avgLookupTime := time.Duration(float64(lookupTotalTime.Nanoseconds()) / float64(lookupCount))
    avgStoreTime := time.Duration(float64(storeTotalTime.Nanoseconds()) / float64(storeCount))

    progress := 100 * float64(pos) / float64(TableSize)
    log.Printf("lookup table init (mode:%6s). "+
        "(progress %12d/%012d %8.4f%%) (t/op:%6dns) elapsed:%6.2fs "+
        "tableSize: %3d %s "+
        "(lookup time: %4dns avg: %4dns)"+
        "(store time: %4dns avg: %4dns)",
        *mode, pos, TableSize, progress, time.Since(cycleStart).Nanoseconds(),
        time.Since(generatorStart).Seconds(), size, szUom,
        lookupTime.Nanoseconds(), avgLookupTime.Nanoseconds(),
        storeTime.Nanoseconds(), avgStoreTime.Nanoseconds())
}

func NewQuickTable(hashFileName string, keySpaceSize, TableSize int) (t *QuickTable, lastSequence []byte) {
    var cycleStart time.Time
    var pos int64
    var mode string
    var tableReady bool
    var table QuickTable
    table.Init(1048576)

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
                logMessage(&mode, headerSize, recordSize, recordCount, pos, int64(TableSize), cycleStart,
                    generatorStart, lookupTime, storeTime)
            }
        }
    }()

    if file.Exists(hashFileName) {
        /*
           		 * A pre-computed hash file exists on disk.  Load that file into our in-memory
                    * hash table.
        */
        log.Printf("hashFile exists (%s). loading...", hashFileName)
        mode = "load"

        fileHandle, err := os.Open(hashFileName)
        if err != nil {
            panic(err)
        }
        gzipReader, err := gzip.NewReader(fileHandle)
        if err != nil {
            panic(err)
        }
        defer func() {
            if gzipReader != nil {
                _ = gzipReader.Close()
            }
            if fileHandle != nil {
                _ = fileHandle.Close()
            }
        }()
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

        err = nil
        compressor, err := gzip.NewReader(fileHandle)
        if err != nil {
            panic(err)
        }
        scanner := bufio.NewScanner(compressor)
        for scanner.Scan() {
            cycleStart = time.Now()
            line := scanner.Text()
            hash, err := hex.DecodeString(line)
            if err != nil {
                panic(err)
            }
            table.Store([20]byte(hash))
            lastSequence = hash
            pos++
        }
        if err := scanner.Err(); err != nil {
            panic(err)
        }
    } else {
        /*
           		 * We have no precomputed hash file.  We need to create one.
                    * Create the in-memory map of hashes, then asynchronously write
                    * the hashes to a precomputed hash file we can use next time.
        */
        log.Printf("hashFile not found (%s). creating...", hashFileName)
        mode = "create"
        for i := 0; i < TableSize; i++ {
            cycleStart = time.Now()
            hash := c.Sha1Bytes()
            table.Store(hash)
            if !table.Lookup(hash) {
                panic("Failed to look-up table after store")
            }
            _ = c.FastIncrement()
            pos = int64(i)
        }
        log.Println("Asynchronously write file.")
        go func() {
            fileHandle, err := os.Create(hashFileName)
            if err != nil {
                log.Fatalf("Cannot create hash file (%s): %v", hashFileName, err)
            }
            compressor, err := gzip.NewWriterLevel(fileHandle, gzip.BestCompression)
            if err != nil {
                panic(err)
            }
            writer := bufio.NewWriterSize(compressor, writeBufferSize)
            defer func() {
                _ = writer.Flush()
                _ = compressor.Flush()
                _ = compressor.Close()
                _ = fileHandle.Close()
            }()
            pos = int64(0)
            for hash, _ := range table.data {
                cycleStart = time.Now()
                // compress and write to file
                if _, err := writer.Write([]byte(hash)); err != nil {
                    panic(err)
                }
                _ = c.FastIncrement()
                pos++
            }
        }()
    }
    lastSequence = c.Bytes()
    return &table, lastSequence
}

//type QuickMap map[byte]QuickMap

type QuickTable struct {
    data map[string]bool
}

func (o *QuickTable) Init(sz int) {
    o.data = make(map[string]bool, sz)
}

func (o *QuickTable) Store(s [20]byte) {
    storeStart := time.Now()
    storeCount++
    defer func() {
        storeTime = time.Since(storeStart)
        storeTotalTime = storeTotalTime + storeTime
    }()
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
