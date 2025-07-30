package gte

//go:generate govalid ./gte.go

type GTE struct {
	// +govalid:gte=18
	Age int

	// +govalid:gte=0
	Score float64

	Struct struct {
		// +govalid:gte=100
		Value int
	}
}