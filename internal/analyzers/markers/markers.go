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

	// TypeMarkers returns markers for struct types.
	TypeMarkers(*ast.TypeSpec) MarkerSet
}

// newMarkers creates a new instance of Markers, initializing the internal map for field markers.
func newMarkers() Markers {
	return &markers{
		fieldMarkers: make(map[*ast.Field]MarkerSet),
		typeMarkers:  make(map[*ast.TypeSpec]MarkerSet),
	}
}

// markers is an implementation of the Markers interface.
type markers struct {
	fieldMarkers map[*ast.Field]MarkerSet
	typeMarkers  map[*ast.TypeSpec]MarkerSet
}

// FieldMarkers retrieves the markers for a given struct field.
func (m *markers) FieldMarkers(field *ast.Field) MarkerSet {
	return m.fieldMarkers[field]
}

// TypeMarkers retrieves the markers for a given struct type.
func (m *markers) TypeMarkers(ts *ast.TypeSpec) MarkerSet {
	return m.typeMarkers[ts]
}

// insertFieldMarker adds a marker to a specific struct field.
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

// insertTypeMarkers adds a set of markers to a struct type.
func (m *markers) insertTypeMarker(ts *ast.TypeSpec, marker Marker) {
	if existing, ok := m.typeMarkers[ts]; ok {
		existing.Add(marker)
		m.typeMarkers[ts] = existing

		return
	}

	ms := newMarkerSet()
	ms.Add(marker)
	m.typeMarkers[ts] = ms
}
