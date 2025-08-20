//go:build test

//go:generate govalid ./fixture.go

// Package testfixture contains request fixtures used by middleware tests.
package testfixture

// PersonRequest is the request payload used in middleware tests.
// +govalid:required
type PersonRequest struct {
	Name string `json:"name"`
	// +govalid:email
	Email string `json:"email"`
}
