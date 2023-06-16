package descriptormap

import (
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse/descriptormap/descriptor"
	"testing"
)

func TestDescriptorMap_List(t *testing.T) {
	var m Map
	m.data = make(map[string]descriptor.Descriptor)
	m.data["a"] = descriptor.Descriptor{}
	m.data["b"] = descriptor.Descriptor{}
	m.data["c"] = descriptor.Descriptor{}
	result := m.List()
	if len(result) != 3 {
		t.Fatal("3 results expected")
	}
	for _, key := range []string{"a", "b", "c"} {
		if _, ok := m.data[key]; !ok {
			t.Fatalf("Failed on %s", key)
		}
	}

}
