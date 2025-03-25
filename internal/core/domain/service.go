package domain

// MessageService represents the core business logic for messages
type MessageService struct {
	defaultMessage string
}

// NewMessageService creates a new MessageService with a default greeting
func NewMessageService() *MessageService {
	return &MessageService{
		defaultMessage: "Hello from hexagonal architecture!",
	}
}

// GetGreetingMessage returns a greeting message
func (s *MessageService) GetGreetingMessage() *Message {
	return NewMessage(s.defaultMessage)
}