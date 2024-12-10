package errors

import (
    "fmt"
    "net/http"
)

// ErrorType representa el tipo de error del dominio
type ErrorType string

const (
    // Tipos de error comunes
    ValidationError   ErrorType = "VALIDATION_ERROR"
    NotFoundError    ErrorType = "NOT_FOUND"
    ConflictError    ErrorType = "CONFLICT"
    UnauthorizedError ErrorType = "UNAUTHORIZED"
    ForbiddenError   ErrorType = "FORBIDDEN"
    InternalError    ErrorType = "INTERNAL_ERROR"
)

// DomainError representa un error del dominio
type DomainError struct {
    Type    ErrorType
    Message string
    Cause   error
    Code    int
}

// Error implementa la interfaz error
func (e *DomainError) Error() string {
    if e.Cause != nil {
        return fmt.Sprintf("%s: %s (causa: %v)", e.Type, e.Message, e.Cause)
    }
    return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// NewValidationError crea un nuevo error de validación
func NewValidationError(message string, cause error) *DomainError {
    return &DomainError{
        Type:    ValidationError,
        Message: message,
        Cause:   cause,
        Code:    http.StatusBadRequest,
    }
}

// NewNotFoundError crea un nuevo error de recurso no encontrado
func NewNotFoundError(message string) *DomainError {
    return &DomainError{
        Type:    NotFoundError,
        Message: message,
        Code:    http.StatusNotFound,
    }
}

// NewConflictError crea un nuevo error de conflicto
func NewConflictError(message string) *DomainError {
    return &DomainError{
        Type:    ConflictError,
        Message: message,
        Code:    http.StatusConflict,
    }
}

// NewUnauthorizedError crea un nuevo error de no autorizado
func NewUnauthorizedError(message string) *DomainError {
    return &DomainError{
        Type:    UnauthorizedError,
        Message: message,
        Code:    http.StatusUnauthorized,
    }
}

// NewForbiddenError crea un nuevo error de acceso prohibido
func NewForbiddenError(message string) *DomainError {
    return &DomainError{
        Type:    ForbiddenError,
        Message: message,
        Code:    http.StatusForbidden,
    }
}

// NewInternalError crea un nuevo error interno
func NewInternalError(message string, cause error) *DomainError {
    return &DomainError{
        Type:    InternalError,
        Message: message,
        Cause:   cause,
        Code:    http.StatusInternalServerError,
    }
}

// IsDomainError verifica si un error es del tipo DomainError
func IsDomainError(err error) bool {
    _, ok := err.(*DomainError)
    return ok
}

// GetErrorCode obtiene el código HTTP correspondiente al error
func GetErrorCode(err error) int {
    if domainErr, ok := err.(*DomainError); ok {
        return domainErr.Code
    }
    return http.StatusInternalServerError
}
