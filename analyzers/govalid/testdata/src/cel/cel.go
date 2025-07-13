package cel

//go:generate govalid ./cel.go

type CEL struct {
	// Basic numeric validation
	// +govalid:cel=value >= 18
	Age int

	// Simple numeric check  
	// +govalid:cel=value > 0.0
	Score float64

	// String length validation
	// +govalid:cel=size(value) > 0
	Name string
}

// Order represents a struct for testing cross-field validation.
type Order struct {
	// Cross-field validation: Price must be less than MaxPrice
	// +govalid:cel=value < this.MaxPrice
	Price float64 `json:"price"`
	
	MaxPrice float64 `json:"max_price"`
	
	// String cross-field validation
	// +govalid:cel=size(value) >= size(this.MinName)
	UserName string `json:"user_name"`
	
	MinName string `json:"min_name"`
}