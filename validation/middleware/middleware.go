package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/sivchari/govalid"
)

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
