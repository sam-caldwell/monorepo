package counters

import "fmt"

// IncrementIf - increment if ok is true otherwise return -1
func (counter *Conditional) IncrementIf(ok bool) (v int, err error) {
	const notOK = -1 //If not ok, return -1
	if ok {
		if counter.value > uint(int(^uint(0)>>1)) {
			err = fmt.Errorf("Conditional exceeds limit")
		}
		v = int(counter.value)
		counter.value++
		return v, err
	}
	return notOK, err
}
