//go:generate govalid .

package multiple_markers

type User struct {
	// +govalid:required
	Name string `json:"name"`

	// +govalid:required
	// +govalid:email
	Email string `json:"email"`

	// +govalid:gte=18
	Age int `json:"age"`
}
