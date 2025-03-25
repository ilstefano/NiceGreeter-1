package handlers

import (
	"encoding/json"
	"net/http"

	"hexapp/internal/core/ports"
)

// HTTPHandler handles HTTP requests
type HTTPHandler struct {
	messageService ports.MessageService
}

// NewHTTPHandler creates a new HTTP handler
func NewHTTPHandler(messageService ports.MessageService) *HTTPHandler {
	return &HTTPHandler{
		messageService: messageService,
	}
}

// GetMessageHandler handles GET requests for messages
func (h *HTTPHandler) GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	message := h.messageService.GetGreetingMessage()
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	response := map[string]string{
		"message": message.GetContent(),
	}
	
	json.NewEncoder(w).Encode(response)
}

// SetupRoutes sets up the HTTP routes
func (h *HTTPHandler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/message", h.GetMessageHandler)
}