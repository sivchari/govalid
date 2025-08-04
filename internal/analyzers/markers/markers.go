package markers

import (
	"go/ast"
	"maps"
)

// Marker represents a single marker with an identifier and associated expressions.
type Marker struct {
	Identifier  string
	Expressions map[string]string
}

// MarkerSet is a collection of markers identified by their string identifiers.
type MarkerSet map[string]Marker

// newMarkerSet creates a new empty MarkerSet.
func newMarkerSet() MarkerSet {
	return make(MarkerSet)
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

// insertFieldMarkers adds a set of markers to a specific struct field.
func (m *markers) insertFieldMarkers(field *ast.Field, markers MarkerSet) {
	if len(markers) == 0 {
		return
	}

	if existing, ok := m.fieldMarkers[field]; ok {
		maps.Copy(existing, markers)
		return
	}
	m.fieldMarkers[field] = markers
}
