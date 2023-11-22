package findCollision

import (
	"github.com/sam-caldwell/monorepo/go/counters"
	"log"
	"strings"
)

// AsynchronousJob - Perform a collision search of a given pattern
func AsynchronousJob(segment, id, workerCount, segmentCount, keySpaceSize int, seed []byte,
	lookupTable *QuickTable, collector *Collector, result chan<- Finding) {

	var err error
	var lhs *counters.ByteCounter
	var rhs *counters.ByteCounter
	var initialRhsState []byte

	log.Printf("Start worker (%d, %d)", segment, id)
	// Create the LHS (Left-hand side) counter.  This will be the counter we compare against.
	if lhs, err = counters.NewByteCounter(keySpaceSize); err != nil {
		result <- Finding{
			Id:        id,
			Collision: false,
			Err:       err,
		}
		return
	}
	// Set the lhs to the seed value, which will be over-ridden by the workerId.
	log.Printf("initial seed: %v", seed)
	if err = lhs.SetBytes(0, seed); err != nil {
		result <- Finding{
			Id:        id,
			Collision: false,
			Err:       err,
		}
		return
	}
	// Set the initial value of LHS to be the worker id (0-255) at the least-significant byte.
	// This will offset each worker by a value of 1.  Our increment (later) will be id * workerCount
	// So that each worker's search pattern will interleave entire key space.
	log.Printf("id: %d, segment: %d", id, segment)
	if err = lhs.Add(id + segment); err != nil {
		result <- Finding{
			Id:        id,
			Collision: false,
			Err:       err,
		}
		return
	}
	log.Printf("lhsStart: %v", strings.TrimLeft(lhs.String(), "0"))
	// Create the Right-Hand Side counter.  This will be the counter we will use to
	// iterate over {0,...,LHS}.  Thus, for every LHS, we will iterate over every value from 0 to LHS
	// testing for a collision.
	if rhs, err = counters.NewByteCounter(keySpaceSize); err != nil {
		result <- Finding{
			Id:        id,
			Collision: false,
			Err:       err,
		}
		return
	}
	_ = rhs.Add(len(*lookupTable))
	initialRhsState = rhs.Bytes()

	// Loop through all values of LHS starting with the initial value.
	// When lhs.Add() is executed, the increment will be the product of the worker id
	// and the number of workers running.
	for {
		collector.Metrics[id].LhsCount++
		collector.Metrics[id].LhsSample = lhs.String()
		//
		// Lookup the first few hashes rather than calculate them
		// This optimization trades memory for cpu/time.
		//
		if _, ok := (*lookupTable)[lhs.Sha1()]; ok {
			log.Printf("prelookup hit: %s", lhs.Sha1())
			result <- Finding{
				Id:        id,
				Hash:      lhs.Sha1(),
				Raw:       lhs.String(),
				Collision: true,
			}
			return
		}
		for {
			collector.Metrics[id].RhsSample = rhs.String()

			// When LHS == RHS, get the next LHS
			if lhs.Equal(rhs) {
				break
			}

			// When LHS != RHS but the two have the same hash, we have a collision.
			if lhs.EqualSha1(rhs) {
				result <- Finding{
					Id:        id,
					Hash:      lhs.Sha1(),
					Raw:       lhs.String(),
					Collision: true,
				}
				return
			}

			// Increment RHS to the next value and carry on until exhausted.
			if err := rhs.Increment(); err != nil {
				break //All RHS values have been exhausted.
			}
		} /* RHS Loop */

		//LHS.Add() increments the LHS value by id*workerCount and when an error occurs, we are exhausted.
		if err := lhs.Add(workerCount * segmentCount); err != nil {
			result <- Finding{
				Id:        id,
				Err:       err,
				Collision: false,
			}
			break
		}
		rhs.Revert(&initialRhsState)
	} /* LHS Loop */
	defer log.Println("worker died")
}
