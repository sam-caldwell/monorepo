package convert

import "fmt"

// Int64ToIntSafe - safely convert int64 to integer
func Int64ToIntSafe(num int64) (int, error) {
	const maxInt = int(^uint(0) >> 1)
	const minInt = -maxInt - 1

	if num > int64(maxInt) || num < int64(minInt) {
		return 0, fmt.Errorf("integer overflow")
	}

	return int(num), nil
}
