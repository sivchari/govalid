//go:generate govalid ./marker.go

package test

type Required struct {
	// +govalid:required
	Name string `validate:"required" json:"name"`

	// +govalid:required
	Age int `validate:"required" json:"age"`

	// +govalid:required
	Items []string `validate:"required" json:"items"`
}

type LT struct {
	// +govalid:lt=10
	Age int `validate:"lt=10" json:"age"`
}

type GT struct {
	// +govalid:gt=100
	Age int `validate:"gt=100" json:"age"`
}

type MaxLength struct {
	// +govalid:maxlength=50
	Name string `validate:"max=50" json:"name"`
}

// Test instances for validation
type TestRequired Required
type TestLT LT
type TestGT GT
type TestMaxLength MaxLength