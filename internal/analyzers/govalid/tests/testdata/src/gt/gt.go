package gt

//go:generate govalid ./gt.go

type GT struct {
	// +govalid:gt=1
	Int int

	// +govalid:gt=1
	Int8 int8

	// +govalid:gt=1
	Int16 int16

	// +govalid:gt=1
	Int32 int32

	// +govalid:gt=1
	Int64 int64

	// +govalid:gt=1
	Float32 float32

	// +govalid:gt=1
	Float64 float64

	// +govalid:gt=1
	Uint uint

	// +govalid:gt=1
	Uint8 uint8

	// +govalid:gt=1
	Uint16 uint16

	// +govalid:gt=1
	Uint32 uint32

	// +govalid:gt=1
	Uint64 uint64

	// +govalid:gt=1
	Uintptr uintptr

	// +govalid:gt=1
	Complex64 complex64

	// +govalid:gt=1
	Complex128 complex128

	// +govalid:gt=1
	String string

	Struct struct {
		// +govalid:gt=1
		Int int
	}
}
