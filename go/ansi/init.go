package ansi

var disabled bool
var debugMode bool

func init() {
	disabled = false  // Use color by default
	debugMode = false // Don't print debug messages by default
}
