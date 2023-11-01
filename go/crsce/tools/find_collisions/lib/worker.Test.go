package hashScanner

/*
 * hashScanner worker.Test() method
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * The worker.Test() method performs a comparison between
 * every member of the lhs queue and every possible value less
 * than this lhs value.
 */

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"log"
	"os"
)

func (w *Worker) Test(source *CandidateQueue) {
	log.Println("Test() running")
	count := 0
	for {
		lhs := source.Pop()
		count++
		log.Printf("consumed lhs: %v (count %d)", lhs.hash, count)
		for {
			if err := w.counter.Increment(); err != nil {
				if err.Error() != errors.OverflowError {
					log.Printf("Unexpected Error: %v", err)
					os.Exit(1)
				}
				break
			}
			if lhs.raw == w.counter.String() {
				// Our RHS counter has reached our LHS value
				// Break the inner loop, grab another LHS value and start over.
				log.Printf("Pass finished for %v", lhs.hash)
				break
			}
			if w.counter.Sha1() == lhs.hash {
				// We have a matching hash (i.e. a collision)
				log.Printf("collision found at %v\n\n---\n%v\n---\n", lhs.hash, lhs.raw)
				os.Exit(0)
			}
		}
		log.Printf("get next lhs (count:%d)", count)
		_ = os.Stdout.Sync()
	}
}
