package length

//go:generate govalid ./length.go

type Length struct {
	// +govalid:length=7
	String string

	Struct struct {
		// +govalid:length=10
		Name string
	}
}