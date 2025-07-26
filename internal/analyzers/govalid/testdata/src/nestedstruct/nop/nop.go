package nop

//go:generate govalid ./nop.go

// Nop struct should NOT generate any validation code
// because it has no govalid markers at all.
type Nop struct {
	A struct {
		X string
	}
}
