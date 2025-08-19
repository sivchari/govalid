//go:generate govalid .

package multiple_markers

// +govalid:required
type Multiple struct {
	Name string `json:"name"`

	// +govalid:email
	Email string `json:"email"`

	// +govalid:gte=18
	Age int `json:"age"`
}
