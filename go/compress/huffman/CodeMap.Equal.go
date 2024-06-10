package huffman

import "bytes"

func (cm *CodeMap) Equal(rhs CodeMap) bool {
	if len(*cm) != len(rhs) {
		return false
	}
	for key, value := range *cm {
		if !bytes.Equal(value, (rhs)[key]) {
			return false
		}
	}
	return true
}
