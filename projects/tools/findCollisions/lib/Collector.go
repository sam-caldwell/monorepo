package findCollision

import "time"

// Collector - Global metric collector for findCollision
type Collector struct {
	StartTime int64
	PrevCount int64
	Metrics   []Metric
}

// NewCollector - Create an array of Metric objects (one per worker).
func NewCollector(sz int) Collector {
	return Collector{
		Metrics:   make([]Metric, sz),
		StartTime: time.Now().Unix(),
	}
}
