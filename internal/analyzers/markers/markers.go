package markers

import (
	"go/ast"
)

// Marker represents a single marker with an identifier and associated expressions.
type Marker struct {
	Identifier  string
	Expressions map[string]string
}

// MarkerSet is an ordered collection of markers that preserves definition order.
type MarkerSet []Marker

// newMarkerSet creates a new empty MarkerSet.
func newMarkerSet() MarkerSet {
	return MarkerSet{}
}

// Add adds a marker to the set if it doesn't already exist (by identifier).
func (ms *MarkerSet) Add(marker Marker) {
	// Check if marker with same identifier already exists
	for i, existing := range *ms {
		if existing.Identifier == marker.Identifier {
			// Update existing marker
			(*ms)[i] = marker

			return
		}
	}
	// Add new marker
	*ms = append(*ms, marker)
}

// Markers is an interface that provides methods to retrieve markers for struct fields.
type Markers interface {
	// FieldMarkers returns markers for struct fields.
	FieldMarkers(*ast.Field) MarkerSet
}

// newMarkers creates a new instance of Markers, initializing the internal map for field markers.
func newMarkers() Markers {
	return &markers{
		fieldMarkers: make(map[*ast.Field]MarkerSet),
	}
}

// markers is an implementation of the Markers interface.
type markers struct {
	fieldMarkers map[*ast.Field]MarkerSet
}

// FieldMarkers retrieves the markers for a given struct field.
func (m *markers) FieldMarkers(field *ast.Field) MarkerSet {
	return m.fieldMarkers[field]
}

// insertFieldMarkers adds a marker to a specific struct field.
func (m *markers) insertFieldMarker(field *ast.Field, marker Marker) {
	if existing, ok := m.fieldMarkers[field]; ok {
		existing.Add(marker)
		m.fieldMarkers[field] = existing

		return
	}

	ms := newMarkerSet()
	ms.Add(marker)
	m.fieldMarkers[field] = ms
}
