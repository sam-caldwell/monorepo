package persistence

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/types"
	"log"
	"os"
	"path/filepath"
)

// Write - Write an event to file.
func (evt *EventCollector) Write(event types.Event) (err error) {
	fileName := filepath.Join(evt.path.String(), fmt.Sprintf("event-%d", evt.id.Get()))
	log.Printf("file created: %s contains: %v", fileName, []byte(event.ToString()))
	return os.WriteFile(fileName, []byte(event.ToString()), 0640)
}
