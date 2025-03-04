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
}

func (m *UserMessage) Request(key, code, source string) UserMessage {
	message := UserMessage{}
	message.Key = key
	message.Code = code
	message.Source = source
	return message
}

func (m *UserMessage) Response(payload any) UserMessage {
	m.Payload = payload
	return *m
}

func (m *UserMessage) Validate(log *log.Logger) error {
	if m.Key == "" {
		log.Println("user Key is required")
		return fmt.Errorf("user Key is required")
	}
	if m.Code == "" {
		log.Println("user Code is required")
		return fmt.Errorf("user Code is required")
	}
	if m.Source == "" {
		log.Println("user Source is required")
		return fmt.Errorf("user Source is required")
	}
	return nil
}
