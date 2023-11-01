package hashScanner

/*
 * HashScanner Worker.Initialize()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * hashScanner is initialized with this method.
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/counters"
	"log"
	"os"
)

// Initialize - initialize a worker.
// Bounds check the worker count.  Then initialize the ByteCounter for the worker with an initial value equal to the
// worker's index value (i).  This value will start each worker one value above the previous worker.  The offset
// will then increment each counter by a count equal to the number of workers.
func (w *Worker) Initialize(i, offset, keySpace uint) (err error) {
	log.Printf("worker %d initializer done", i)
	defer func() { _ = os.Stdout.Sync() }()
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
