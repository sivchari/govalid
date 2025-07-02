package maxitems

//go:generate govalid ./maxitems.go

type MaxItems struct {
	// +govalid:maxitems=5
	Slice []string

	// +govalid:maxitems=3
	Array [10]int

	Struct struct {
		// +govalid:maxitems=2
		Items []int
	}
}