package counters

import "fmt"

// Increment - Increment value
func (counter *Conditional) Increment() (v int, err error) {
	if counter.value > uint(int(^uint(0)>>1)) {
		err = fmt.Errorf("conditional exceeds limit")
	}
	v = int(counter.value)
	counter.value++
	return v, err
}
