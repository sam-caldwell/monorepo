package cli

import (
    "os"
    "strings"
)

// GetDebugFlag - Scan os.Args for --debug and return true.
func GetDebugFlag() bool {
    for arg := range os.Args {
        s := strings.ToLower(os.Args[arg])
        if s == "--debug" {
            return true
        }
    }
    return false
}
