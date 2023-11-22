package findCollision

import (
	"github.com/sam-caldwell/monorepo/go/counters"
	"log"
)

type QuickTable map[string]bool

func NewQuickTable(keySpaceSize, TableSize int) *QuickTable {
	log.Printf("initializing the lookup table for %d hashes", TableSize)
	table := make(QuickTable)
	c, _ := counters.NewByteCounter(keySpaceSize)

	for i := 0; i < TableSize; i++ {
		log.Printf("pre-populating lookup table.  row %f", 100*float64(i)/float64(TableSize))
		table[c.Sha1()] = true
		_ = c.Increment()
	}
	log.Printf("quickTable initialized")
	return &table
}
