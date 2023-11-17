package metrics

import "sync"

type TextMetric struct {
	lock  sync.Mutex
	value []byte
}
