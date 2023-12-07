package findCollision

import (
	"log"
	"math"
	"strings"
	"time"
)

// TimedLogger - Log the metrics via a goroutine using a timer
func TimedLogger(interval time.Duration, workerCount, startingWorkerId, timeWindow int, collector *Collector) {
	var currCount int64

	log.Printf("logging starting...")

	metricReportingTimer := time.NewTicker(interval)
	defer metricReportingTimer.Stop()

	for {
		select {
		case <-metricReportingTimer.C:

			// Capture the amount of time elapsed since the last log event.
			duration := time.Now().Unix() - collector.StartTime

			// For some metrics (such as counts) we need an aggregation.
			// To improve speed, we unroll the loop a bit.
			currCount += collector.Metrics[0].LhsCount +
				collector.Metrics[1].LhsCount +
				collector.Metrics[2].LhsCount +
				collector.Metrics[3].LhsCount +
				collector.Metrics[4].LhsCount +
				collector.Metrics[5].LhsCount +
				collector.Metrics[6].LhsCount +
				collector.Metrics[7].LhsCount +
				collector.Metrics[8].LhsCount
			for id := 9; id < workerCount; id++ {
				currCount += collector.Metrics[id].LhsCount
			}
			for workerId := 0; workerId < workerCount; workerId++ {
				// To reduce overhead, we sample certain worker information on each cycle.
				//workerId := int(sequentialNumber % uint(workerCount))
				//sequentialNumber++

				lhsSample := strings.TrimLeft(collector.Metrics[workerId].LhsSample[:], "0")
				rhsSample := strings.TrimLeft(collector.Metrics[workerId].RhsSample[:], "0")

				log.Printf("t:%4d, currCount: %12d, prevCount: %12d, chg Ops: %12.f, "+
					"id: %2d, lhs:%d:%s (%.f bytes) rhs: %s (%.f bytes)",
					duration,
					currCount,
					collector.PrevCount,
					float64(currCount-collector.PrevCount)/float64(timeWindow),
					workerId,
					startingWorkerId,
					lhsSample,
					math.Ceil(float64(len(lhsSample))/2),
					rhsSample,
					math.Ceil(float64(len(rhsSample))/2))
			}
			collector.PrevCount = currCount
			currCount = 0
		}
	}
}
