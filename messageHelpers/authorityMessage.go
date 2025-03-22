package messageHelpers

type AuthorityMessage struct {
	Key       string           `json:"Key"`
	User      UserMessage      `json:"User"`
	Behaviour BehaviourMessage `json:"Behaviour"`
	Payload   any              `json:"Payload"`
	Source    string           `json:"Source"`
}

type GrantMessage struct {
	User      UserMessage      `json:"User"`
	Behaviour BehaviourMessage `json:"Behaviour"`
}

func (m *GrantMessage) Request(user UserMessage, behaviour BehaviourMessage) GrantMessage {
	message := GrantMessage{}
	message.User = user
	message.Behaviour = behaviour
	return message
}

func (m *GrantMessage) Response(payload any) GrantMessage {
	m.Behaviour.Payload = payload
	return *m
}

type RevokeMessage struct {
	User      UserMessage      `json:"User"`
	Behaviour BehaviourMessage `json:"Behaviour"`
}

func (m *RevokeMessage) Request(user UserMessage, behaviour BehaviourMessage) RevokeMessage {
	message := RevokeMessage{}
	message.User = user
	message.Behaviour = behaviour
	return message
}

func (m *RevokeMessage) Response(payload any) RevokeMessage {
	m.Behaviour.Payload = payload
	return *m
}
