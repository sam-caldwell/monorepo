package list

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

// FromStringP - Given a string pointer, parse by delimiter and clean up
func FromStringP(s *string, sep string, pruneQuotes bool) []string {
	// Set up our cleaner function so we are only doing our if pruneQuotes
	// check one time.
	if s == nil {
		return []string{}
	}
	var cleaner func(*string) *string
	if pruneQuotes {
		cleaner = func(i *string) *string {
			*i = strings.TrimSpace(*i)
			*i = strings.TrimPrefix(*i, words.SingleQuote)
			*i = strings.TrimPrefix(*i, words.DoubleQuote)
			*i = strings.TrimSuffix(*i, words.SingleQuote)
			*i = strings.TrimSuffix(*i, words.DoubleQuote)
			return i
		}
	} else {
		cleaner = func(i *string) *string {
			return i
		}
	}
	data := strings.Split(*(cleaner(s)), *sep)
	for i := range data {
		// Creating a pointer because we want to defeat
		// multiple copies of the string
		dp := &data[i]
		dp = cleaner(dp)
	}
	return data
}
