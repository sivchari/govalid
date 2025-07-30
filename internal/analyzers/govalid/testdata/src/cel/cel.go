package cel

import "time"

//go:generate govalid ./cel.go

type CEL struct {
	// Basic comparison operators
	// +govalid:cel=value >= 18
	Age int

	// +govalid:cel=value > 0.0
	Score float64

	// +govalid:cel=value <= 100
	MaxScore int

	// +govalid:cel=value < 1000
	Limit int

	// +govalid:cel=value == 42
	Answer int

	// +govalid:cel=value != 0
	NonZero int

	// String operations
	// +govalid:cel=size(value) > 0
	Name string

	// +govalid:cel=size(value) >= 3 && size(value) <= 50
	Username string

	// +govalid:cel=value.startsWith('prefix_')
	PrefixedName string

	// +govalid:cel=value.endsWith('.com')
	Email string

	// +govalid:cel=value.contains('@')
	EmailAddress string

	// Boolean operations
	// +govalid:cel=value == true
	IsActive bool

	// +govalid:cel=value != false
	MustBeTrue bool

	// Complex expressions with multiple operators
	// +govalid:cel=value >= 0 && value <= 120
	ValidAge int

	// +govalid:cel=value > 0.0 && value <= 100.0
	Percentage float64

	// +govalid:cel=size(value) >= 8 && size(value) <= 256
	Password string

	// Cross-field validation (this references)
	// +govalid:cel=value >= this.Age
	MinAge int

	// +govalid:cel=value <= this.MaxScore
	CurrentScore int

	// +govalid:cel=size(value) >= size(this.Name)
	LongName string

	// Complex cross-field with multiple conditions
	// +govalid:cel=value > this.Age && value < this.Limit
	MiddleValue int

	// Arithmetic operations
	// +govalid:cel=value >= this.Age * 2
	DoubleAge int

	// +govalid:cel=value <= this.MaxScore / 2
	HalfScore int

	// +govalid:cel=value == this.Age + this.NonZero
	SumValue int

	// Complex boolean logic
	// +govalid:cel=(value >= 18 && value <= 65) || value == 100
	SpecialAge int

	// +govalid:cel=value > 0 || (value == 0 && this.IsActive)
	ConditionalValue int

	// String pattern matching
	// +govalid:cel=value.matches('^[A-Z][a-z]+$')
	ProperName string

	// Slice/Array operations (if supported)
	// +govalid:cel=size(value) >= 1 && size(value) <= 10
	Items []string

	// +govalid:cel=size(value) > 0
	NonEmptySlice []int

	// New advanced features - temporarily commented out due to type issues
	// +govalid:cel=value > 0
	PositiveValue int

	// +govalid:cel='admin' in value
	HasAdminRole []string

	// +govalid:cel=int(value) >= 18
	AgeFromString string

	// +govalid:cel=string(value) in ['active', 'inactive', 'pending']
	StatusCode int

	// Timestamp comparison - temporarily commented out due to struct comparison issues  
	// BirthDate time.Time

	// +govalid:cel=value > duration('1h')
	ProcessingTime time.Duration

	// List comprehensions
	// +govalid:cel=value.all(item, size(item) > 0)
	AllNonEmpty []string

	// +govalid:cel=value.exists(item, item == 'target')
	HasTarget []string

	// +govalid:cel=value.exists_one(item, item == 'unique')
	HasUniqueItem []string

	// +govalid:cel=value.all(item, item.startsWith('prefix'))
	AllPrefixed []string

	// +govalid:cel=value.exists(item, item.contains('@'))
	HasEmailFormat []string

	// +govalid:cel=size(value.filter(item, item.startsWith('prefix'))) > 0
	FilteredItems []string

	// +govalid:cel=size(value.map(item, size(item))) == size(value)
	MappedSizes []string
}
