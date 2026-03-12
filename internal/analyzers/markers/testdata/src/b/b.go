package b

// Test that old format (// +govalid:) reports an error diagnostic and is skipped.

type OldFormatMarkers struct {
	// +govalid:required // want `deprecated marker format`
	Required string

	// +govalid:lt=10 // want `deprecated marker format`
	Min int

	// +govalid:required // want `deprecated marker format`
	// +govalid:lt=10 // want `deprecated marker format`
	RequiredAndMin string
}

// +govalid:required // want `deprecated marker format`
type TypeLevelOldFormat struct {
	String string

	// +govalid:required // want `deprecated marker format`
	RequiredString int

	// +govalid:email // want `deprecated marker format`
	Email string
}
