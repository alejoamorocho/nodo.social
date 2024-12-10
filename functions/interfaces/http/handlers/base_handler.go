package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/kha0sys/nodo.social/functions/domain/errors"
)

// BaseHandler provides common HTTP handler functionality
type BaseHandler struct {
    // Common fields and dependencies
}

// ErrorResponse represents a response error
type ErrorResponse struct {
    Type    string `json:"type"`
    Message string `json:"message"`
}

// RespondWithError sends an error response
func (h *BaseHandler) RespondWithError(w http.ResponseWriter, err error) {
    var response ErrorResponse
    var statusCode int

    if domainErr, ok := err.(*errors.DomainError); ok {
        response = ErrorResponse{
            Type:    string(domainErr.Type),
            Message: domainErr.Message,
        }
        statusCode = domainErr.Code
    } else {
        response = ErrorResponse{
            Type:    string(errors.InternalError),
            Message: "Error interno del servidor",
        }
        statusCode = http.StatusInternalServerError
    }

    h.RespondWithJSON(w, statusCode, response)
}

// RespondWithJSON sends a JSON response
func (h *BaseHandler) RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, err := json.Marshal(payload)
    if err != nil {
        h.RespondWithError(w, errors.NewInternalError("Error al serializar respuesta", err))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

// ValidateRequest validates the request body
func (h *BaseHandler) ValidateRequest(r *http.Request, v interface{}) error {
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(v); err != nil {
        return errors.NewValidationError("Error al decodificar el cuerpo de la petición", err)
    }
    
    if validator, ok := v.(interface{ Validate() error }); ok {
        if err := validator.Validate(); err != nil {
            return err
        }
    }
    
    return nil
}

// ExtractPaginationParams extracts pagination parameters from the request
func (h *BaseHandler) ExtractPaginationParams(r *http.Request) (page, size int, err error) {
    pageStr := r.URL.Query().Get("page")
    sizeStr := r.URL.Query().Get("size")

    if pageStr != "" {
        page, err = strconv.Atoi(pageStr)
        if err != nil {
            return 0, 0, errors.NewValidationError("Parámetro 'page' inválido", err)
        }
    }

    if sizeStr != "" {
        size, err = strconv.Atoi(sizeStr)
        if err != nil {
            return 0, 0, errors.NewValidationError("Parámetro 'size' inválido", err)
        }
    }

    if page < 0 {
        page = 0
    }
    if size <= 0 {
        size = 10
    }
    if size > 100 {
        size = 100
    }

    return page, size, nil
}
