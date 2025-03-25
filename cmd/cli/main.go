package main

import (
	"hexapp/internal/adapters/handlers"
	"hexapp/internal/core/domain"
)

func main() {
	// Initialize the domain service
	messageService := domain.NewMessageService()

	// Initialize the CLI handler
	cliHandler := handlers.NewCLIHandler(messageService)

	// Print the message
	cliHandler.PrintMessage()
}