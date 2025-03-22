package messageHelpers

import (
	"fmt"
	"log"
)

type UserMessage struct {
	Key     string `json:"Key"`
	Code    string `json:"Code"`
	Payload any    `json:"Payload"`
	Source  string `json:"Source"`
	Locale  string `json:"Locale"`
}

func (m *UserMessage) Request(key, code, source, locale string) UserMessage {
	message := UserMessage{}
	message.Key = key
	message.Code = code
	message.Source = source
	message.Locale = locale
	return message
}

func (m *UserMessage) Response(payload any) UserMessage {
	m.Payload = payload
	return *m
}

func (m *UserMessage) Validate(log *log.Logger) error {
	if m.Key == "" {
		log.Printf("[%v] user Key is required", "frantic-core")
		return fmt.Errorf("user Key is required")
	}
	if m.Code == "" {
		log.Printf("[%v] user Code is required", "frantic-core")
		return fmt.Errorf("user Code is required")
	}
	if m.Source == "" {
		log.Printf("[%v] user Source is required", "frantic-core")
	}
	return nil
}
