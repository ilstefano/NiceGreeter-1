package main

import (
        "log"
        "net/http"

        "hexapp/internal/adapters/handlers"
        "hexapp/internal/core/domain"
)

func main() {
        // Initialize the domain service
        messageService := domain.NewMessageService()

        // Initialize the HTTP handler
        httpHandler := handlers.NewHTTPHandler(messageService)

        // Setup the HTTP server
        mux := http.NewServeMux()
        httpHandler.SetupRoutes(mux)

        // Start the server
        log.Println("Starting HTTP server on :8080")
        log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))
}