package messageHelpers

import (
	"time"
)

type SessionMessage struct {
	SessionID    string
	Expiry       time.Time
	UserKey      string
	UserCode     string
	SessionToken any
}
