package messageHelpers

type UserMessage struct {
	Key     string `json:"Key"`
	Code    string `json:"Code"`
	Payload any    `json:"Payload"`
}

func (m *UserMessage) Request(key string, code string) UserMessage {
	message := UserMessage{}
	message.Key = key
	message.Code = code
	return message
}

func (m *UserMessage) Response(payload any) UserMessage {
	m.Payload = payload
	return *m
}
