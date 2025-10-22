// Package errors provides enhanced error handling capabilities for the API.
// It allows for structured error responses with additional context and HTTP status codes.
package errors

import (
	"encoding/json"
	"fmt"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
)

// EnhancedError represents a structured error response with additional context.
// It includes an error code, message, HTTP status code, and optional details.
type EnhancedError struct {
	// ErrorCode is a unique identifier for the type of error
	ErrorCode string `json:"error_code"`
	// ErrorMessage provides a human-readable description of the error
	ErrorMessage string `json:"error_message,omitempty"`
	// ErrorHTTPCode represents the HTTP status code associated with this error
	ErrorHTTPCode int `json:"error_http_code,omitempty"`
	// Details can contain any additional information about the error
	Details interface{} `json:"details,omitempty"`
}

// New creates a new EnhancedError instance with the specified error code and details.
// Parameters:
//   - errorCode: string identifier for the type of error
//   - details: additional context or data related to the error
//
// Returns:
//   - *EnhancedError: a pointer to the newly created error
func New(errorCode string, details interface{}) *EnhancedError {
	return &EnhancedError{
		ErrorCode: errorCode,
		Details:   details,
	}
}

// ToMap converts the EnhancedError to a map representation.
// This method enriches the error with standard error information before conversion.
// Returns:
//   - map[string]interface{}: a map containing the error information
func (e *EnhancedError) ToMap() map[string]interface{} {
	e.enrichWithErrorKind()
	errorMap := map[string]interface{}{
		"error_code":    e.ErrorCode,
		"error_message": e.ErrorMessage,
		"details":       e.Details,
	}

	return errorMap
}

// ToJSON serializes the EnhancedError into a JSON byte array.
// This method enriches the error with standard error information before serialization.
// Returns:
//   - []byte: JSON representation of the error
//   - nil if serialization fails
func (e *EnhancedError) ToJSON() []byte {
	e.enrichWithErrorKind()
	responseData, err := json.Marshal(e)
	if err != nil {
		logging.Error(err, nil)
		return nil
	}
	return responseData
}

// ToError converts the EnhancedError to a standard error interface and its map representation.
// This method enriches the error with standard error information before conversion.
// Returns:
//   - map[string]interface{}: a map containing the error information
//   - error: a standard error interface containing the error message
func (e *EnhancedError) ToError() (map[string]interface{}, error) {
	e.enrichWithErrorKind()
	return e.ToMap(), fmt.Errorf("%s", e.ErrorMessage)
}

// enrichWithErrorKind internal method that populates the error with standard information
// based on the error code. If the error code is not found in the ErrorKinds enum,
// it falls back to InternalServerError.
// This method modifies the EnhancedError in place, setting the appropriate
// error message and HTTP status code.
func (e *EnhancedError) enrichWithErrorKind() {
	errorCode := e.ErrorCode
	errorMessage, ok := enum.ErrorKinds[errorCode]
	if !ok {
		errorCode = enum.InternalServerError
		errorMessage = enum.ErrorKinds[enum.InternalServerError]
	}
	e.ErrorCode = errorCode
	e.ErrorMessage = errorMessage.ErrorMessage
	e.ErrorHTTPCode = errorMessage.HTTPCode
}
