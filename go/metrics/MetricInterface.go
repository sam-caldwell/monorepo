package metrics

import "time"

type MetricType interface {
	string | time.Duration | AnyNumber
}

// Metric - Implement a generic interface for the Metric structs
// this will ensure we have common interfaces.
type Metric[MT MetricType, CT AnyNumber] interface {
	Push(value MT)
	Count() CT
}
