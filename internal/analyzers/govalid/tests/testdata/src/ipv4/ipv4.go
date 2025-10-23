package ipv4

//go:generate govalid ./ipv4.go

type IPv4 struct {
	// +govalid:ipv4
	Value string

	// +govalid:ipv4
	Struct struct {
		Value string
	}
}
