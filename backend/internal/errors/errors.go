package errors

import (
	"encoding/json"
	"net/http"
)

// APIError struct to represent API errors
type APIError struct {
    Code    int     `json:"code"`
	Message string  `json:"error"`
}

// RespondWithError sends a JSON error response
func RespondWithError(w http.ResponseWriter, code int, message string) {
    RespondWithDetailedError(w, code, message)
}

// RespondWithDetailedError sends a JSON error response with status code and message
func RespondWithDetailedError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
    apiError := APIError{Code: code, Message: message}
	json.NewEncoder(w).Encode(apiError)
}


// NewBadRequestError creates a new 400 bad request error
func NewBadRequestError(message string) APIError {
   return APIError{Code: http.StatusBadRequest, Message: message}
}
// NewUnauthorizedError creates a new 401 unauthorized error
func NewUnauthorizedError(message string) APIError {
  return APIError{Code: http.StatusUnauthorized, Message: message}
}
// NewNotFoundError creates a new 404 not found error
func NewNotFoundError(message string) APIError{
  return APIError{Code: http.StatusNotFound, Message: message}
}
// NewInternalServerError creates a new 500 internal server error
func NewInternalServerError(message string) APIError {
   return APIError{Code: http.StatusInternalServerError, Message: message}
}