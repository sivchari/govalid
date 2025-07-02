package lte

//go:generate govalid ./lte.go

type LTE struct {
	// +govalid:lte=100
	Age int

	// +govalid:lte=10.5
	Score float64

	Struct struct {
		// +govalid:lte=50
		Value int
	}
}