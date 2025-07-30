package enum

//go:generate govalid ./enum.go

// Custom string type
type UserRole string

// Custom int type
type Priority int

type Enum struct {
	// String enum
	// +govalid:enum=admin,user,guest
	Role string

	// Numeric enum
	// +govalid:enum=1,2,3
	Level int

	// Custom string type enum
	// +govalid:enum=manager,developer,tester
	UserRole UserRole

	// Custom numeric type enum
	// +govalid:enum=10,20,30
	Priority Priority
}