//go:generate govalid ./numeric.go

package numeric

type Numeric struct {
	// +govalid:numeric
	Number string
}
