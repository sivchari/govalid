package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sivchari/govalid"
)

type ContextKey string

const ValidatedBodyKey ContextKey = "validatedBody"

// ValidateRequest creates an HTTP middleware that validates incoming JSON request bodies
// using govalid validators. The validated body is stored in the request context with ValidatedBodyKey.
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

		ctx := context.WithValue(r.Context(), ValidatedBodyKey, body)
		next(w, r.WithContext(ctx))
	}
}
