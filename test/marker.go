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

type MinLength struct {
	// +govalid:minlength=3
	Name string `validate:"min=3" json:"name"`
}

type GTE struct {
	// +govalid:gte=18
	Age int `validate:"gte=18" json:"age"`
}

type LTE struct {
	// +govalid:lte=100
	Age int `validate:"lte=100" json:"age"`
}

type MaxItems struct {
	// +govalid:maxitems=5
	Items []string `validate:"max=5" json:"items"`

	// +govalid:maxitems=3
	Metadata map[string]string `json:"metadata"`

	// +govalid:maxitems=2
	ChanField chan int `json:"chan_field"`
}

type MinItems struct {
	// +govalid:minitems=2
	Items []string `validate:"min=2" json:"items"`

	// +govalid:minitems=1
	Metadata map[string]string `json:"metadata"`

	// +govalid:minitems=1
	ChanField chan int `json:"chan_field"`
}
// Custom types for enum testing
type UserRole string
type Priority int

type Enum struct {
	// String enum
	// +govalid:enum=admin,user,guest
	Role string `json:"role"`

	// Numeric enum
	// +govalid:enum=1,2,3
	Level int `json:"level"`

	// Custom string type enum
	// +govalid:enum=manager,developer,tester
	UserRole UserRole `json:"user_role"`

	// Custom numeric type enum
	// +govalid:enum=10,20,30
	Priority Priority `json:"priority"`
}

type Email struct {
	// +govalid:email
	Email string `validate:"email" json:"email"`
}

type UUID struct {
	// +govalid:uuid
	UUID string `validate:"uuid" json:"uuid"`
}

type URL struct {
	// +govalid:url
	URL string `validate:"url" json:"url"`
}
