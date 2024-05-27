package LogEvent

import (
	"encoding/json"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
	"time"
)

// ToJson - marshals RFC5424Message to a JSON byte slice
func (e *RFC5424Message) ToJson() ([]byte, error) {
	return json.Marshal(e)
}

// ToJsonString - marshals RFC5424Message to a JSON string
func (e *RFC5424Message) ToJsonString() (string, error) {
	b, err := e.ToJson()
	return string(b), err
}

func (e *RFC5424Message) Create(messagePriority LogLevel.Value, appName, msgId *string, message *MessageValue) *RFC5424Message {
	var (
		err      error
		hostname string
	)
	hostname, err = os.Hostname()
	if err != nil {
		hostname = "not_available"
	}
	return &(RFC5424Message{
		Priority:  uint(messagePriority),
		Version:   version.Version,
		Timestamp: time.Now(),
		Hostname:  hostname,
		AppName:   *appName,
		ProcID:    os.Getpid(),
		MsgID:     *msgId,
		Message:   message,
	})
}
