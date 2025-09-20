//go:generate govalid .

package multiple

// +govalid:required
type Multiple struct {
	Name string `json:"name"`

	// +govalid:email
	Email string `json:"email"`

	// +govalid:gte=18
	Age int `json:"age"`
}
