package findCollision

import (
	"bufio"
	"encoding/hex"
	"github.com/sam-caldwell/monorepo/go/counters"
	"github.com/sam-caldwell/monorepo/go/fs/file"
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

func NewQuickTable(keySpaceSize, TableSize int) (t *QuickTable, lastSequence []byte) {
	var cycleStart time.Time
	var pos int
	var mode string
	var tableReady bool
	table := make(QuickTable)
	c, _ := counters.NewByteCounter(keySpaceSize)
	generatorStart := time.Now()
	defer func() { tableReady = true }()
	go func() {
		t := time.NewTicker(1 * time.Second)
		defer t.Stop()
		for !tableReady {
			select {
			case <-t.C:
				size := table.SizeOf()
				szUom := "bytes"
				if size > 1024 {
					szUom = "KB"
				}
				if size > 1048576 {
					szUom = "MB"
				}
				if size > 1073741824 {
					szUom = "GB"
				}
				progress := 100 * float64(pos) / float64(TableSize)
				log.Printf("lookup table init (mode:%6s). "+
					"(progress %12d/%012d %8.4f%%) (t/op:%6dns) elapsed:%v "+
					"size: %d%s",
					mode, pos, TableSize, progress, time.Since(cycleStart).Nanoseconds(),
					time.Since(generatorStart).Seconds(), size, szUom)
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

type QuickTable map[byte]QuickTable

func (o *QuickTable) Store(s [20]byte) {
	(*o)[s[0]] = QuickTable{
		s[1]: QuickTable{
			s[2]: QuickTable{
				s[3]: QuickTable{
					s[4]: QuickTable{
						s[5]: QuickTable{
							s[6]: QuickTable{
								s[7]: QuickTable{
									s[8]: QuickTable{
										s[9]: QuickTable{
											s[10]: QuickTable{
												s[11]: QuickTable{
													s[12]: QuickTable{
														s[13]: QuickTable{
															s[14]: QuickTable{
																s[15]: QuickTable{
																	s[16]: QuickTable{
																		s[17]: QuickTable{
																			s[18]: QuickTable{
																				s[19]: nil,
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	if !o.Lookup(s) {
		panic("store failed")
	}
}

func (o *QuickTable) Lookup(s [20]byte) bool {
	if n1, ok := (*o)[s[0]]; ok {
		if n2, ok := n1[s[1]]; ok {
			if n3, ok := n2[s[2]]; ok {
				if n4, ok := n3[s[3]]; ok {
					if n5, ok := n4[s[4]]; ok {
						if n6, ok := n5[s[5]]; ok {
							if n7, ok := n6[s[6]]; ok {
								if n8, ok := n7[s[7]]; ok {
									if n9, ok := n8[s[8]]; ok {
										if n10, ok := n9[s[9]]; ok {
											if n11, ok := n10[s[10]]; ok {
												if n12, ok := n11[s[11]]; ok {
													if n13, ok := n12[s[12]]; ok {
														if n14, ok := n13[s[13]]; ok {
															if n15, ok := n14[s[14]]; ok {
																if n16, ok := n15[s[15]]; ok {
																	if n17, ok := n16[s[16]]; ok {
																		if n18, ok := n17[s[17]]; ok {
																			if n19, ok := n18[s[18]]; ok {
																				if _, ok := n19[s[19]]; ok {
																					return true
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}

func getNodeSize(n map[byte]QuickTable) int64 {
	sz := int64(unsafe.Sizeof(n))
	for _, thisN := range n {
		if thisN != nil {
			sz += getNodeSize(thisN)
		}
	}
	return sz
}

func (o *QuickTable) SizeOf() int64 {
	mutex.Lock()
	defer mutex.Unlock()
	sz := int64(unsafe.Sizeof(*o))
	for _, n := range *o {
		if n != nil {
			sz = getNodeSize(n)
		}
	}
	return sz
}
