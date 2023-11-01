package hashScanner

/*
 * HashScanner Worker
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * hashScanner is a reusable library that creates a worker-based solution for collision detection in SHA1.
 */

import (
	"github.com/sam-caldwell/monorepo/go/counters"
)

type Worker struct {
	id      uint8
	offset  uint
	counter counters.ByteCounter
}
