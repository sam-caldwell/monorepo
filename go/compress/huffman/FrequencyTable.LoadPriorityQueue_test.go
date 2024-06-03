package huffman

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestLoadPriorityQueue(t *testing.T) {
	// Step 1: Create a frequency table for testing
	frequencyTable := FrequencyTable{
		'a': 5,
		'b': 9,
		'c': 12,
		'd': 13,
		'e': 16,
		'f': 45,
	}

	// Step 2: Load the priority queue from the frequency table
	pq := frequencyTable.LoadPriorityQueue()

	// Step 3: Validate the priority queue
	expectedNodes := []*Node{
		{symbol: 'a', frequency: 5},
		{symbol: 'b', frequency: 9},
		{symbol: 'c', frequency: 12},
		{symbol: 'd', frequency: 13},
		{symbol: 'e', frequency: 16},
		{symbol: 'f', frequency: 45},
	}

	if pq.Len() != len(expectedNodes) {
		t.Errorf("Priority queue has incorrect length, got: %d, want: %d", pq.Len(), len(expectedNodes))
	}

	for _, expectedNode := range expectedNodes {
		actualNode := heap.Pop(&pq).(*Node)
		if !reflect.DeepEqual(actualNode, expectedNode) {
			t.Errorf("Node does not match, got: %v, want: %v", actualNode, expectedNode)
		}
	}
}
