//go:generate govalid ./fixture.go

package fixture

// +govalid:required
type PersonRequest struct {
	Name string `json:"name"`
	// +govalid:email
	Email string `json:"email"`
}
