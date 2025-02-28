package messagehelper

import (
	"time"
)

type Session struct {
	SessionID    string
	Expiry       time.Time
	UserKey      string
	UserCode     string
	SessionToken any
}
