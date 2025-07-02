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

type MaxItems struct {
	// +govalid:maxitems=5
	Items []string `validate:"max=5" json:"items"`
}

type MinLength struct {
	// +govalid:minlength=3
	Name string `validate:"min=3" json:"name"`
}

type MinItems struct {
	// +govalid:minitems=2
	Items []string `validate:"min=2" json:"items"`
}

type GTE struct {
	// +govalid:gte=18
	Age int `validate:"gte=18" json:"age"`
}

type LTE struct {
	// +govalid:lte=100
	Score int `validate:"lte=100" json:"score"`
}

type Enum struct {
	// +govalid:enum=admin user guest
	Role string `json:"role"`
}
