package lt

//go:generate govalid ./lt.go

type LT struct {
	// +govalid:lt=1
	Int int

	// +govalid:lt=1
	Int8 int8

	// +govalid:lt=1
	Int16 int16

	// +govalid:lt=1
	Int32 int32

	// +govalid:lt=1
	Int64 int64

	// +govalid:lt=1
	Float32 float32

	// +govalid:lt=1
	Float64 float64

	// +govalid:lt=1
	Uint uint

	// +govalid:lt=1
	Uint8 uint8

	// +govalid:lt=1
	Uint16 uint16

	// +govalid:lt=1
	Uint32 uint32

	// +govalid:lt=1
	Uint64 uint64

	// +govalid:lt=1
	Uintptr uintptr

	// +govalid:lt=1
	Complex64 complex64

	// +govalid:lt=1
	Complex128 complex128

	// +govalid:lt=1
	String string

	Struct struct {
		// +govalid:lt=1
		Int int
	}
}
