package messageHelpers

import (
	"log"

	ce "github.com/mt1976/frantic-core/commonErrors"
)

type BehaviourMessage struct {
	Key     string `json:"Key"`
	Source  string `json:"Source"`
	Payload any    `json:"Payload"`
}

type DeclareMessage struct {
	Domain    string           `json:"Domain"`
	Behaviour BehaviourMessage `json:"Behaviour"`
}

func (m *DeclareMessage) Request(domain string, behaviour string) DeclareMessage {
	message := DeclareMessage{}
	message.Domain = domain
	message.Behaviour = BehaviourMessage{Key: behaviour}
	return message
}

func (m *DeclareMessage) Response(payload any) DeclareMessage {
	m.Behaviour.Payload = payload
	return *m
}

func (m *BehaviourMessage) Validate(log *log.Logger) error {
	if m.Key == "" {
		log.Printf("[%v] Behaviour Key is required", "frantic-core")
		return ce.ErrKeyRequiredWrapper("behaviour")
	}

	if m.Source == "" {
		log.Printf("[%v] Warning! Behaviour Source is not provided", "frantic-core")
		return nil
	}
	return nil
}
