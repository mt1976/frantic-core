package messageHelpers

import (
	"log"

	ce "github.com/mt1976/frantic-core/commonErrors"
)

type UserMessage struct {
	Key      string `json:"Key"`
	Code     string `json:"Code"`
	Payload  any    `json:"Payload"`
	Source   string `json:"Source"`
	Locale   string `json:"Locale"`
	Theme    string `json:"Theme"`
	Timezone string `json:"Timezone"`
	Role     string `json:"Role"`
	Spare0   string `json:"Spare0"`
	Spare1   string `json:"Spare1"`
	Spare2   string `json:"Spare2"`
}

func (m *UserMessage) Request(key, code, source, locale, theme, timezone, role string) UserMessage {
	message := UserMessage{}
	message.Key = key
	message.Code = code
	message.Source = source
	message.Locale = locale
	if locale == "" {
		message.Locale = cfg.GetApplication_Locale()
	}
	message.Theme = theme
	if theme == "" {
		message.Theme = cfg.GetApplication_Theme()
	}
	message.Timezone = timezone
	if timezone == "" {
		message.Timezone = cfg.GetApplication_Timezone()
	}
	message.Role = role
	if role == "" {
		message.Role = "default"
	}
	return message
}

func (m *UserMessage) Response(payload any) UserMessage {
	m.Payload = payload
	return *m
}

func (m *UserMessage) Validate(log *log.Logger) error {
	if m.Key == "" {
		log.Printf("[%v] user Key is required", "frantic-core")
		return ce.ErrKeyRequiredWrapper("user")
	}
	if m.Code == "" {
		log.Printf("[%v] user Code is required", "frantic-core")
		return ce.ErrCodeRequiredWrapper("user")
	}
	if m.Source == "" {
		log.Printf("[%v] user Source is required", "frantic-core")
	}
	return nil
}
