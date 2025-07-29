package maxitems

//go:generate govalid ./maxitems.go

type MaxItems struct {
	// +govalid:maxitems=5
	Slice []string

	// +govalid:maxitems=3
	Array [10]int

	// +govalid:maxitems=4
	MapField map[string]int

	// +govalid:maxitems=2
	ChanField chan string

	Struct struct {
		// +govalid:maxitems=2
		Items []int
	}
}