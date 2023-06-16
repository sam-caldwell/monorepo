package descriptormap

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/argparse/descriptormap/descriptor"
	"testing"
)

func TestDescriptorMap_Count(t *testing.T) {
	var m Map
	if m.Count() != 0 {
		t.Fatal("expect map count 0 initially")
	}
	m.data = make(map[string]descriptor.Descriptor)

	for i := 1; i <= 10; i++ {
		item := fmt.Sprintf("item%d", i)
		m.data[item] = descriptor.Descriptor{}
		if m.Count() != i {
			t.Fatalf("expect map count %d initially", i)
		}
	}

}
