/*
Copyright 2025 sivchari.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package markers

import (
	"go/ast"
)

type MarkerSet map[string]Marker

func newMarkerSet() MarkerSet {
	return make(MarkerSet)
}

type Marker struct {
	Identifier  string
	Expressions map[string]string
}

type Markers interface {
	// FieldMarkers returns markers for struct fields.
	FieldMarkers(*ast.Field) MarkerSet
}

func newMarkers() Markers {
	return &markers{
		fieldMarkers: make(map[*ast.Field]MarkerSet),
	}
}

type markers struct {
	fieldMarkers map[*ast.Field]MarkerSet
}

func (m *markers) FieldMarkers(field *ast.Field) MarkerSet {
	return m.fieldMarkers[field]
}

func (m *markers) insertFieldMarkers(field *ast.Field, markers MarkerSet) {
	if len(markers) == 0 {
		return
	}
	m.fieldMarkers[field] = markers
}
