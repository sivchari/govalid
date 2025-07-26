package inside

//go:generate govalid ./inside.go

// Inside struct should generate validation code for
// the nested field with marker inside the anonymous struct.
type Inside struct {
	A struct {
		// +govalid:required
		X string
	}
}
