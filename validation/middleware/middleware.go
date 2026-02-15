// Package middleware provides HTTP middleware for validating request payloads.
package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sivchari/govalid"
)

// ValidateRequest returns a middleware that validates the request body using the provided Validator type.
// If validation fails, it responds with a 400 Bad Request.
func ValidateRequest[T govalid.Validator](next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body T
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)

			return
		}

		if err := body.Validate(); err != nil {
			http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)

			return
		}

		next(w, r)
	}
}

// ValidateRequestContext returns a middleware that validates the request body using the provided Validator ValidateContext method
// Honors *http.Request provided context.
// Returns 408 for context errors (timeout or cancellation), and 400 for validation errors.
func ValidateRequestContext[T govalid.Validator](next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body T
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)

			return
		}

		if err := body.ValidateContext(r.Context()); err != nil {
			statusCode := http.StatusBadRequest
			if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
				statusCode = http.StatusRequestTimeout
			}

			http.Error(w, "Validation error: "+err.Error(), statusCode)

			return
		}

		next(w, r)
	}
}
