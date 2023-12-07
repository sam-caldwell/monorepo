package metrics

import (
	"github.com/sam-caldwell/monorepo/go/types/generic"
	"math/big"
	"sync"
)

/*
 * Theory of Use
 *  The BigScalar is intended to store any scalar metric.
 *
 *  (For the new coder, we have all been there.  Let me explain.  A scalar metric is a measurable
 *   value (such as distance or temperature).  We just give it a fancy name, so we seem smarter
 *   and can confuse people.)
 */

// BigScalar - A generic metric for all things big and numeric with a larger counter type
type BigScalar[ValueType generic.AnyNumber] struct {

	// Lock - some safety for the concurrent stuff.
	lock sync.Mutex

	// value - The FIFO metric value bucket array for storing a moving window of metrics.
	value []ValueType

	// count - a total count of all instances of this metric for all time
	count big.Int

	// movingWindow - A moving window for calculating min/max and moving averages
	// This is used to determine the value bucket array size during initialization.
	movingWindow uint8 //The moving window for min/max
}
