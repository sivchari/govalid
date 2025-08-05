package minitems

//go:generate govalid ./minitems.go

type MinItems struct {
	// +govalid:minitems=2
	Slice []string

	// +govalid:minitems=3
	Array [10]int

	// +govalid:minitems=1
	MapField map[string]int

	// +govalid:minitems=2
	ChanField chan string

	Struct struct {
		// +govalid:minitems=1
		Items []int
	}
}