package benchmark

type Required struct {
	// +govalid:required
	Name string `validate:"required" json:"name"`

	// +govalid:required
	Age int `validate:"required" json:"age"`

	// +govalid:required
	Items []string `validate:"required" json:"items"`
}

type Min struct {
	// +govalid:min=10
	Age int `validate:"min=10" json:"age"`
}
