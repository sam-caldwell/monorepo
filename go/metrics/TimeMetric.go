package metrics

import (
	"sync"
	"time"
)

/*
 * Theory of Use
 *	The TimeMetric is intended to store any time-based metrics.
 *
 *  (Technically time is scalar since it is a measurable thing we can represent on a line.
 *  But time is special as a type and in how we track and analyze it.  I may be wrong here
 *  and this may get folded into Scalar)
 */

type TimeMetric struct {
	lock  sync.Mutex
	value time.Duration
}
