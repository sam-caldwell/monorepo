package LogEvent

import "time"

// RFC5424Message represents an RFC5424 syslog message
//
//	(c) 2023 Sam Caldwell.  MIT License
type RFC5424Message struct {
	Priority      uint         `json:"priority"`
	Version       string       `json:"version"`
	Timestamp     time.Time    `json:"timestamp"`
	Hostname      string       `json:"hostname"`
	AppName       string       `json:"appName"`
	ProcID        int          `json:"procID"`
	MsgID         string       `json:"msgID"`
	RateLimit     uint         `json:"rateLimit"`
	RateAvailable uint         `json:"rateAvailable"`
	Message       MessageValue `json:"message,omitempty"`
}
