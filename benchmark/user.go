package benchmark

type User struct {
	// +govalid:required
	Name string `validate:"required" json:"name"`

	// +govalid:required
	Age int `validate:"required" json:"age"`

	// +govalid:required
	Items []string `validate:"required" json:"items"`
}
