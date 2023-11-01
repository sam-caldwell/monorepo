package hashScanner

/*
 * HashScanner Worker.Initialize()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * hashScanner is initialized with this method.
 */

import (
	"log"
	"os"
)

// EnumerateKeyspace - start the worker.  Iterate through the keyspace and push the hash and string to the
// candidate queue.
func (w *Worker) EnumerateKeyspace(candidate *CandidateQueue) {
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
