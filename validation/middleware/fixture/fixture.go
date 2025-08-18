//go:generate govalid ./fixture.go

// Package fixture contains request fixtures used by middleware tests.
package fixture

// PersonRequest is the request payload used in middleware tests.
// +govalid:required
type PersonRequest struct {
	Name string `json:"name"`
	// +govalid:email
	Email string `json:"email"`
}
