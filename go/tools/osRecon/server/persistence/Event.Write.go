package persistence

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/types"
	"os"
	"path/filepath"
)

func (evt *EventCollector) Write(event types.Event) error {
	fileName := filepath.Join(evt.path.String(), fmt.Sprintf("event-%d", evt.id.Get()))
	return os.WriteFile(fileName, []byte(event.ToString()), 0640)
}
