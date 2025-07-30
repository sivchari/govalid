package minlength

//go:generate govalid ./minlength.go

type MinLength struct {
	// +govalid:minlength=5
	String string

	Struct struct {
		// +govalid:minlength=3
		Name string
	}
}