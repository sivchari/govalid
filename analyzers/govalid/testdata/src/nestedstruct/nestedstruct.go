package nestedstruct

//go:generate govalid ./nestedstruct.go

// NoMarkersSimpleNested - This struct should NOT generate any validation code
// because it has no govalid markers at all.
type NoMarkersSimpleNested struct {
	A struct {
		X string
	}
}

// PartialMarkersOuter - This struct should generate validation code only for
// the field with marker, not for the nested struct without markers.
type PartialMarkersOuter struct {
	// +govalid:required
	Name string

	NestedWithoutMarkers struct {
		X string
		Y int
	}
}

// NestedMarkerInside - This struct should generate validation code for
// the nested field with marker inside the anonymous struct.
type NestedMarkerInside struct {
	A struct {
		// +govalid:required
		X string
	}
}
