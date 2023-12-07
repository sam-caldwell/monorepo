package metrics

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/types/generic"
	hashes2 "github.com/sam-caldwell/monorepo/go/types/hashes"
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
	case Scalar[big.Int], Scalar[big.Float],
		Scalar[int], Scalar[int8], Scalar[int16], Scalar[int32], Scalar[int64],
		Scalar[uint], Scalar[uint8], Scalar[uint16], Scalar[uint32], Scalar[uint64],
		Scalar[float32], Scalar[float64],
		BigScalar[big.Int], BigScalar[big.Float],
		BigScalar[int], BigScalar[int8], BigScalar[int16], BigScalar[int32], BigScalar[int64],
		BigScalar[uint], BigScalar[uint8], BigScalar[uint16], BigScalar[uint32], BigScalar[uint64],
		BigScalar[float32], BigScalar[float64],
		StateMetric[[]byte], StateMetric[string],
		StateMetric[hashes2.Sha1], StateMetric[hashes2.Sha256], StateMetric[hashes2.Sha512],
		StateMetric[generic.Block1KB], StateMetric[generic.Block2KB],
		StateMetric[generic.Block4KB], StateMetric[generic.Block8KB],
		StateMetric[generic.Block16KB]: //,
		//TimeMetric:
		//store.data[name] = metricType
	default:
		//
		// Record the error if an unsupported type is used.
		// But do not create the metric.
		//
		store.err = append(store.err, fmt.Errorf("unsupported metric type"))
	}
	//
	// Return the store reference, allowing AddMetric() calls to be chained
	// in a really nice way.
	//
	return store
}
