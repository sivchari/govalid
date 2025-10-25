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

type Alpha struct {
	// +govalid:alpha
	FirstName string `validate:"alpha" json:"first_name"`

	// +govalid:alpha
	LastName string `validate:"alpha" json:"last_name"`

	// +govalid:alpha
	CountryCode string `validate:"alpha" json:"country_code"`
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

type CEL struct {
	// +govalid:cel=value >= 18
	Age int `json:"age"`

	// +govalid:cel=size(value) > 0
	Name string `json:"name"`

	// +govalid:cel=value > 0.0
	Score float64 `json:"score"`

	// +govalid:cel=value == true
	IsActive bool `json:"is_active"`
}

type CELCrossField struct {
	// Cross-field validation: Price must be less than MaxPrice
	// +govalid:cel=value < this.MaxPrice
	Price float64 `json:"price"`

	MaxPrice float64 `json:"max_price"`

	// Cross-field validation: Quantity * Price <= Budget
	// +govalid:cel=value * this.Price <= this.Budget
	Quantity float64 `json:"quantity"`

	Budget float64 `json:"budget"`
}

type Length struct {
	// +govalid:length=7
	Name string `validate:"len=7" json:"name"`
}

type Numeric struct {
	// +govalid:numeric
	Number string `json:"Number"`
}

type MultipleErrors struct {
	// +govalid:required
	URL string `validate:"required" json:"url"`

	// +govalid:maxlength=1
	TooLong string `validate:"max=1" json:"too_long"`
}

type IPV4 struct {
	// +govalid:ipv4
	IP string `validate:"ipv4" json:"ip"`
}
