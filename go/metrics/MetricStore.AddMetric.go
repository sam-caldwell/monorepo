package metrics

import (
	"fmt"
	"math/big"
)

// AddMetric adds a new metric to the store.
func (store *MetricStore) AddMetric(name string, metric interface{}) *MetricStore {

	store.lock.Lock()
	defer store.lock.Unlock()

	// Ensure that we have no duplicate metric name
	// If one exists, we will use it, but we will save the error
	if _, ok := store.data[name]; ok {
		store.err = append(store.err, fmt.Errorf("duplicate metric found (%s)", name))
	}

	switch metric.(type) {
	case NumericMetric[big.Int, big.Int], NumericMetric[big.Float, big.Int], NumericMetric[big.Float, big.Float],
		NumericMetric[int, uint64], NumericMetric[int8, uint64], NumericMetric[int16, uint64],
		NumericMetric[int32, uint64], NumericMetric[int64, uint64], NumericMetric[uint, uint64],
		NumericMetric[uint8, uint64], NumericMetric[uint16, uint64], NumericMetric[uint32, uint64],
		NumericMetric[uint64, uint64], NumericMetric[float32, uint64], NumericMetric[float64, uint64],
		TimeMetric, TextMetric:
		store.data[name] = metric
	default:
		// Handle unsupported metric type, or log an error.
		// You can choose to panic, log, or take any appropriate action here.
		store.err = fmt.Errorf("unsupported metric type")
	}
	return store
}
