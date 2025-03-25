package handlers

import (
	"fmt"

	"hexapp/internal/core/ports"
)

// CLIHandler handles CLI commands
type CLIHandler struct {
	messageService ports.MessageService
}

// NewCLIHandler creates a new CLI handler
func NewCLIHandler(messageService ports.MessageService) *CLIHandler {
	return &CLIHandler{
		messageService: messageService,
	}
}

// PrintMessage prints the greeting message to the console
func (h *CLIHandler) PrintMessage() {
	message := h.messageService.GetGreetingMessage()
	fmt.Println(message.GetContent())
}