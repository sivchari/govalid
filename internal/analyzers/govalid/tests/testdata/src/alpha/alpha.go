//go:generate ./alpha.go
package alpha

// Alpha is a struct for testing alpha validation
type Alpha struct {
	// +govalid:alpha
	FirstName string `json:"first_name"`

	// +govalid:alpha
	LastName string `json:"last_name"`

	// +govalid:alpha
	CountryCode string `json:"country_code"`
}