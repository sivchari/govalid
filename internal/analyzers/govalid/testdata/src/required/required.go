package required

//go:generate govalid ./required.go

type Required struct {
	// +govalid:required
	String string

	// +govalid:required
	Int int

	// +govalid:required
	Array [1]string

	// +govalid:required
	Slice []string

	// +govalid:required
	Map map[string]string

	// +govalid:required
	Interface interface{}

	// +govalid:required
	Any any

	// +govalid:required
	Pointer *string

	// +govalid:required
	EmptyStruct struct{}

	// +govalid:required
	EntireRequiredStruct struct {
		EntireRequiredStructName string
	}

	PartialStruct struct {
		// +govalid:required
		PartialStructString string

		Int int
	}

	NestedStruct struct {
		// +govalid:required
		Nested2 struct {
			Nested2String string
		}
	}

	OtherNestedStruct struct {
		// +govalid:required
		Nested2 struct {
			Nested2String string
		}
	}

	// +govalid:required
	Channel chan string

	// +govalid:required
	Func func(string) string

	// +govalid:required
	Named Named
}

type Named string
