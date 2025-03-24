package messageHelpers

import (
	"time"
)

type SessionMessage struct {
	SessionID    string      `json:"SessionID"`
	Expiry       time.Time   `json:"Expiry"`
	UserKey      string      `json:"UserKey"`
	UserCode     string      `json:"UserCode"`
	SessionToken any         `json:"SessionToken"`
	User         UserMessage `json:"User"`
	Payload      any         `json:"Payload"`
	Locale       string      `json:"Locale"`
	Spare1       string      `json:"Spare1"`
	Spare2       string      `json:"Spare2"`
}

func (m *SessionMessage) Request(sessionID string, expiry time.Time, user UserMessage) SessionMessage {
	message := SessionMessage{}
	message.SessionID = sessionID
	message.Expiry = expiry
	message.User = user
	message.UserKey = user.Key
	message.UserCode = user.Code
	message.Locale = user.Locale
	return message
}

func (m *SessionMessage) Response(payload any) SessionMessage {
	m.Payload = payload
	return *m
}
