package metrics

import (
	"fmt"
	"math/big"
)

// AddMetric adds a new dynamically typed metric to the store.
func (store *MetricStore) AddMetric(name string, metricType interface{}) *MetricStore {

	store.lock.Lock()
	defer store.lock.Unlock()

	// Ensure that we have no duplicate metric name
	// If one exists, we will use it, but we will save the error
	if _, ok := store.data[name]; ok {
		store.err = append(store.err, fmt.Errorf("duplicate metric found (%s)", name))
	}

	// Verify that the metric is a valid type.
	// If it is, create the metric...
	// If it is not, then bail and store the error.  Do not create the metric.
	switch metricType.(type) {
	case NumericMetric[big.Int, big.Int], NumericMetric[big.Float, big.Int], NumericMetric[big.Float, big.Float],
		NumericMetric[int, big.Int], NumericMetric[int8, big.Int], NumericMetric[int16, big.Int],
		NumericMetric[int32, big.Int], NumericMetric[int64, big.Int], NumericMetric[uint, big.Int],
		NumericMetric[uint8, big.Int], NumericMetric[uint16, big.Int], NumericMetric[uint32, big.Int],
		NumericMetric[uint64, big.Int], NumericMetric[float32, big.Int], NumericMetric[float64, big.Int],
		NumericMetric[int, uint64], NumericMetric[int8, uint64], NumericMetric[int16, uint64],
		NumericMetric[int32, uint64], NumericMetric[int64, uint64], NumericMetric[uint, uint64],
		NumericMetric[uint8, uint64], NumericMetric[uint16, uint64], NumericMetric[uint32, uint64],
		NumericMetric[uint64, uint64], NumericMetric[float32, uint64], NumericMetric[float64, uint64],
		TimeMetric, TextMetric:
		store.data[name] = metricType
	default:
		store.err = append(store.err, fmt.Errorf("unsupported metric type"))
	}
	return store
}
