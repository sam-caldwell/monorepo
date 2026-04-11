package matrix

import (
	"gocv.io/x/gocv"
	"sync"
)

// Matrix represents a 2D matrix.
//
// (c) 2024 Sam Caldwell.  See License.txt
type Matrix struct {
	lock sync.RWMutex
	data gocv.Mat
}
