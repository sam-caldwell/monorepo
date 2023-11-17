package metrics

import (
	"sync"
	"time"
)

type TimeMetric struct {
	lock  sync.Mutex
	value time.Duration
}
