package utils

import (
	"encoding/json"
	"net/http"
)

// Response es una estructura estandarizada para todas las respuestas HTTP
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

// ErrorInfo contiene información detallada sobre un error
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// RespondWithError envía una respuesta de error estandarizada
func RespondWithError(w http.ResponseWriter, statusCode int, errorCode string, message string) {
	response := Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    errorCode,
			Message: message,
		},
	}
	RespondWithJSON(w, statusCode, response)
}

// RespondWithJSON envía una respuesta JSON exitosa estandarizada
func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	var response Response
	switch v := payload.(type) {
	case Response:
		response = v
	default:
		response = Response{
			Success: true,
			Data:    payload,
		}
	}

	json.NewEncoder(w).Encode(response)
}

// Common error codes
const (
	ErrInvalidInput     = "INVALID_INPUT"
	ErrNotFound         = "NOT_FOUND"
	ErrInternalServer   = "INTERNAL_SERVER_ERROR"
	ErrUnauthorized     = "UNAUTHORIZED"
	ErrForbidden        = "FORBIDDEN"
)

