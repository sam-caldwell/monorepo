package convert

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/metrics"
)

// NumericErrToStringErr - given a numeric type and error, return its string, error.  (useful as a wrapper)
func NumericErrToStringErr[V metrics.Number](n V, err error) (string, error) {
	return fmt.Sprintf("%v", n), err
}
