package LogEvent

import "time"

type LogFormat interface {
	ToJson() ([]byte, error)
	FromJson([]byte) error
}

// RFC5424Message represents an RFC5424 syslog message
type RFC5424Message struct {
	Priority  uint         `json:"priority"`
	Version   string       `json:"version"`
	Timestamp time.Time    `json:"timestamp"`
	Hostname  string       `json:"hostname"`
	AppName   string       `json:"appName"`
	ProcID    int          `json:"procID"`
	MsgID     string       `json:"msgID"`
	Message   MessageValue `json:"message,omitempty"`
}
