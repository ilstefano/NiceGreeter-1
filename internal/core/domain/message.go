package domain

// Message represents the core domain entity
type Message struct {
	Content string
}

// NewMessage creates a new Message with the given content
func NewMessage(content string) *Message {
	return &Message{
		Content: content,
	}
}

// GetContent returns the message content
func (m *Message) GetContent() string {
	return m.Content
}