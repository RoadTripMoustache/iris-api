package errors

import (
	"encoding/json"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		errorCode string
		details   interface{}
	}{
		{
			name:      "Create with valid error code",
			errorCode: enum.InternalServerError,
			details:   map[string]string{"key": "value"},
		},
		{
			name:      "Create with nil details",
			errorCode: enum.InternalServerError,
			details:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := New(tt.errorCode, tt.details)
			assert.NotNil(t, err)
			assert.Equal(t, tt.errorCode, err.ErrorCode)
			assert.Equal(t, tt.details, err.Details)
		})
	}
}

func TestEnhancedError_ToMap(t *testing.T) {
	tests := []struct {
		name           string
		enhancedError  *EnhancedError
		expectedResult map[string]interface{}
	}{
		{
			name: "Valid error code",
			enhancedError: &EnhancedError{
				ErrorCode: enum.InternalServerError,
				Details:   map[string]string{"detail": "test"},
			},
			expectedResult: map[string]interface{}{
				"error_code":    enum.InternalServerError,
				"error_message": enum.ErrorKinds[enum.InternalServerError].ErrorMessage,
				"details":       map[string]string{"detail": "test"},
			},
		},
		{
			name: "Invalid error code",
			enhancedError: &EnhancedError{
				ErrorCode: "INVALID_CODE",
				Details:   nil,
			},
			expectedResult: map[string]interface{}{
				"error_code":    enum.InternalServerError,
				"error_message": enum.ErrorKinds[enum.InternalServerError].ErrorMessage,
				"details":       nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.enhancedError.ToMap()
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestEnhancedError_ToJSON(t *testing.T) {
	tests := []struct {
		name          string
		enhancedError *EnhancedError
		wantNil       bool
	}{
		{
			name: "Successful JSON conversion",
			enhancedError: &EnhancedError{
				ErrorCode: enum.InternalServerError,
				Details:   map[string]string{"detail": "test"},
			},
			wantNil: false,
		},
		{
			name: "JSON conversion without details",
			enhancedError: &EnhancedError{
				ErrorCode: enum.InternalServerError,
			},
			wantNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.enhancedError.ToJSON()
			if tt.wantNil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)

				var decoded EnhancedError
				err := json.Unmarshal(result, &decoded)
				assert.NoError(t, err)
				assert.Equal(t, tt.enhancedError.ErrorCode, decoded.ErrorCode)
			}
		})
	}
}

func TestEnhancedError_ToError(t *testing.T) {
	tests := []struct {
		name          string
		enhancedError *EnhancedError
	}{
		{
			name: "Valid error conversion",
			enhancedError: &EnhancedError{
				ErrorCode: enum.InternalServerError,
				Details:   map[string]string{"detail": "test"},
			},
		},
		{
			name: "Conversion with invalid code",
			enhancedError: &EnhancedError{
				ErrorCode: "INVALID_CODE",
				Details:   nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorMap, err := tt.enhancedError.ToError()

			assert.NotNil(t, errorMap)
			assert.Error(t, err)
			assert.Contains(t, errorMap, "error_code")
			assert.Contains(t, errorMap, "error_message")
			assert.Contains(t, errorMap, "details")
		})
	}
}

func TestEnhancedError_enrichWithErrorKind(t *testing.T) {
	tests := []struct {
		name             string
		enhancedError    *EnhancedError
		expectedCode     string
		expectedHTTPCode int
	}{
		{
			name: "Enrich with valid code",
			enhancedError: &EnhancedError{
				ErrorCode: enum.InternalServerError,
			},
			expectedCode:     enum.InternalServerError,
			expectedHTTPCode: enum.ErrorKinds[enum.InternalServerError].HTTPCode,
		},
		{
			name: "Enrich with invalid code",
			enhancedError: &EnhancedError{
				ErrorCode: "INVALID_CODE",
			},
			expectedCode:     enum.InternalServerError,
			expectedHTTPCode: enum.ErrorKinds[enum.InternalServerError].HTTPCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.enhancedError.enrichWithErrorKind()
			assert.Equal(t, tt.expectedCode, tt.enhancedError.ErrorCode)
			assert.Equal(t, tt.expectedHTTPCode, tt.enhancedError.ErrorHTTPCode)
			assert.NotEmpty(t, tt.enhancedError.ErrorMessage)
		})
	}
}
