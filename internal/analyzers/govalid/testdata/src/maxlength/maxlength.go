package maxlength

//go:generate govalid ./maxlength.go

type MaxLength struct {
	// +govalid:maxlength=10
	String string

	Struct struct {
		// +govalid:maxlength=20
		Name string
	}
}