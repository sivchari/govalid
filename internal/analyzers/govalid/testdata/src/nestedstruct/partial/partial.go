package partial

//go:generate govalid ./partial.go

// Partial struct should generate validation code only for
// the field with marker, not for the nested struct without markers.
type Partial struct {
	// +govalid:required
	Name string

	NestedWithoutMarkers struct {
		X string
		Y int
	}
}
