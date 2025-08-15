package middleware_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sivchari/govalid/validation/middleware"
	"github.com/sivchari/govalid/validation/middleware/fixture"
)

func TestValidateRequest(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		expectedBody   string
		handlerCalled  bool
	}{
		{
			name:           "valid request",
			requestBody:    fixture.PersonRequest{Name: "John", Email: "john@example.com"},
			expectedStatus: http.StatusOK,
			expectedBody:   "Person created: John",
			handlerCalled:  true,
		},
		{
			name:           "invalid email - validation error",
			requestBody:    fixture.PersonRequest{Name: "John", Email: "invalid-email"},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Validation error",
			handlerCalled:  false,
		},
		{
			name:           "invalid json",
			requestBody:    "invalid json", // Raw string for invalid JSON
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid JSON",
			handlerCalled:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			handlerCalled := false
			testHandler := func(w http.ResponseWriter, r *http.Request) {
				handlerCalled = true
				person := r.Context().Value(middleware.ValidatedBodyKey).(*fixture.PersonRequest)
				w.Write([]byte("Person created: " + person.Name))
			}

			handler := middleware.ValidateRequest[*fixture.PersonRequest](testHandler)

			var req *http.Request
			if str, ok := tt.requestBody.(string); ok {
				req = httptest.NewRequest("POST", "/test", bytes.NewBufferString(str))
			} else {
				jsonBody, _ := json.Marshal(tt.requestBody)
				req = httptest.NewRequest("POST", "/test", bytes.NewBuffer(jsonBody))
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
			if !bytes.Contains(rr.Body.Bytes(), []byte(tt.expectedBody)) {
				t.Errorf("Expected body to contain '%s', got '%s'", tt.expectedBody, rr.Body.String())
			}
			if handlerCalled != tt.handlerCalled {
				t.Errorf("Expected handler called: %v, got: %v", tt.handlerCalled, handlerCalled)
			}
		})
	}
}
