package hashScanner

/*
 * HashScanner
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * hashScanner is a reusable library that creates a worker-based solution for collision detection in SHA1.
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/counters"
	"time"
)

const (
	maxSupportedWorkers = 256 // This maximum number of workers possible in the system.
)

type Worker struct {
	id      uint8
	offset  uint
	counter counters.ByteCounter
}

// Initialize - initialize a worker.
// Bounds check the worker count.  Then initialize the ByteCounter for the worker with an initial value equal to the
// worker's index value (i).  This value will start each worker one value above the previous worker.  The offset
// will then increment each counter by a count equal to the number of workers.
func (w *Worker) Initialize(i, offset, keySpace uint) (err error) {
	if i > (maxSupportedWorkers - 1) {
		return fmt.Errorf("this system does not support %d workers (max:%d)", i, maxSupportedWorkers)
	}
	if offset > maxSupportedWorkers {
		return fmt.Errorf("offset cannot be more than  %d workers", maxSupportedWorkers)
	}
	w.id = uint8(i)
	w.offset = offset
	if w.counter, err = counters.NewByteCounter(int(keySpace)); err != nil {
		return fmt.Errorf("error initializing counter (%v)", err)
	}
	return w.counter.Set(0, byte(i))
}

// Start - start the worker.
func (w *Worker) Start(exit chan bool) {
	fmt.Printf("worker (%v) started\n", w.id)
	time.Sleep(10)
	exit <- true
}
