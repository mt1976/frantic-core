package messageHelpers

type TranslationMessage struct {
	Text        string `json:"Text"`
	Locale      string `json:"Locale"`
	Origin      string `json:"Origin"`
	Translation string `json:"Translation"`
	Payload     any    `json:"Payload"`
}

func (m *TranslationMessage) Request(text string, locale string, origin string) TranslationMessage {
	message := TranslationMessage{}
	message.Text = text
	message.Locale = locale
	message.Origin = origin
	return message
}

func (m *TranslationMessage) Response(translation string) TranslationMessage {
	m.Translation = translation
	return *m
}

func (m *TranslationMessage) ReponseWithPayload(translation string, payload any) TranslationMessage {
	message := m.Response(translation)
	message.Payload = payload
	return message
}
