package orderedset

import "fmt"

func (set *Set) Delete(pos int) error {
	if (pos >= 0) && (pos < len(set.data)) {
		set.data = append(set.data[:pos], set.data[pos+1:]...)
		return nil
	}
	return fmt.Errorf("index out of range")
}
