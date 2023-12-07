package findCollision

import (
	"log"
	"os"
)

// CollectFindings - Block and collect all findings, reporting them to the logger
func CollectFindings(workerCount int, findings <-chan Finding) {

	log.Printf("Waiting on results")

	for workerCheckedIn := 0; workerCheckedIn < workerCount; {

		log.Printf("\tWaiting...")

		thisResult := <-findings

		if thisResult.Err != nil {

			log.Printf("Error(id: %d): %v", thisResult.Id, thisResult.Err)
			os.Exit(1)

		}

		if thisResult.Collision {

			//report the collision and keep searching.
			log.Printf("collision found (id: %2d) (%v) in %v",
				thisResult.Id, thisResult.Hash, thisResult.Raw)

		} else {

			log.Printf("worker %d finished", workerCheckedIn)
			workerCheckedIn++ // We finished, but we have not encountered any collisions or errors.

		}

	}

}
