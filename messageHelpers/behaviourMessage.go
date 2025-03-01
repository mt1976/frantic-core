package messageHelpers

type BehaviourMessage struct {
	Key     string `json:"Key"`
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
