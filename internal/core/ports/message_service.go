package ports

import "hexapp/internal/core/domain"

// MessageService defines the port (interface) for message services
type MessageService interface {
	GetGreetingMessage() *domain.Message
}