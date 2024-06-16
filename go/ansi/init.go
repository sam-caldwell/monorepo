package ansi

var disabled bool
var debugMode bool

// init
//
//	(c) 2023 Sam Caldwell.  MIT License
func init() {
	Reset()
	disabled = false  // Use color by default
	debugMode = false // Don't print debug messages by default
}
