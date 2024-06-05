package huffman

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
)

// Dump - Write out the frequency table to stdout.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (freq *FrequencyTable) Dump() {
	ansi.Blue().Println("Frequency Table").Reset()
	for symbol, count := range *freq {
		ansi.Blue().Printf("%s   %d", string(symbol), count).LF().Reset()
	}
	ansi.Blue().Line(words.Hyphen, 40).LF().Reset()
}
