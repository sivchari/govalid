// Package basic provides a basic example of using govalid for struct validation.
package basic

//go:generate govalid .

// User represents a user with validation rules.
type User struct {
	//govalid:required
	//govalid:minlength=1
	//govalid:maxlength=100
	Name string

	//govalid:required
	//govalid:email
	Email string

	//govalid:gte=0
	//govalid:lte=150
	Age int
}
