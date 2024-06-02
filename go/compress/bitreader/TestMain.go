package bitreader

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Run tests
	retCode := m.Run()
	// Additional teardown if needed
	// Exiting
	os.Exit(retCode)
}
