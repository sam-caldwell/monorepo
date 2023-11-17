package metrics

import (
	"github.com/sam-caldwell/monorepo/go/types/generic"
	"time"
)

type MetricType interface {
	string | time.Duration | generic.AnyNumber
}

// Metric - Implement a generic interface for the Metric structs
// this will ensure we have common interfaces.
type Metric[MT MetricType, CT generic.AnyNumber] interface {
	Push(value MT)
	Count() CT
}
