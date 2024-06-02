package bitreader

import (
	"fmt"
	"testing"
)

func TestBitReader_Error(t *testing.T) {
	var br BitReader
	br.err = fmt.Errorf("test error")
	if br.err == nil {
		t.Fatal("expected error")
	}
	if br.err.Error() != "test error" {
		t.Fatalf("expected 'test error', got '%s'", br.err.Error())
	}
	if err := br.Error(); err.Error() != "test error" {
		t.Fatalf("expected 'test error', got '%s'", err.Error())
	}
}
