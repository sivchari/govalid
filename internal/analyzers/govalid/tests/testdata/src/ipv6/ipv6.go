package ipv6

//go:generate govalid ./ipv6.go

type IPv6 struct {
	// +govalid:ipv6
	Value string

	// +govalid:ipv6
	Struct struct {
		Value string
	}
}
