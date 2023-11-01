package hashScanner

/*
 * HashScanner Worker.Initialize()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * hashScanner is initialized with this method.
 */

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"log"
	"os"
	"sync"
)

// EnumerateKeyspace - start the worker.  Iterate through the keyspace and push the hash and string to the
// candidate queue.
func (w *Worker) EnumerateKeyspace(wg *sync.WaitGroup, candidate *CandidateQueue) {
	//defer wg.Done()
	defer func() { _ = os.Stdout.Sync() }()
	for {
		if err := w.counter.Add(int(w.offset)); err != nil {
			log.Printf("Error incrementing (worker %d): %v", w.id, err)
			break
		}
		candidate.Push(w.counter.String(), w.counter.Sha1())
		//log.Printf("Worker %d added %v (sz:%d)", w.id, w.counter.Sha1(), len(candidate))
		_ = os.Stdout.Sync()
	}
	log.Printf("worker(%d) terminated\n", w.id)
}

func (w *Worker) Test(wg *sync.WaitGroup, source *CandidateQueue) {
	//defer wg.Done()
	defer func() { _ = os.Stdout.Sync() }()
	//log.Println("Start scanning worker")
	_ = os.Stdout.Sync()
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
				// We have reached the end of the rhs keyspace.
				// Break inner loop and get another LHS to start over.
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
