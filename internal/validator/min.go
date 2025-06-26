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

// Package validator implements rules for validating fields.
package validator

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/markers"
)

type minValidator struct {
	pass     *codegen.Pass
	field    *ast.Field
	minValue string
}

var _ Validator = (*minValidator)(nil)

const minKey = "%s-min"

func (m *minValidator) Validate() string {
	return fmt.Sprintf("t.%s < %s", m.FieldName(), m.minValue)
}

func (m *minValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *minValidator) Err() string {
	if generatorMemory[fmt.Sprintf(minKey, m.FieldName())] {
		return ""
	}

	generatorMemory[fmt.Sprintf(minKey, m.FieldName())] = true

	return strings.ReplaceAll(`
	// Err@Min is the error returned when the value of the field is less than the minimum value.
	Err@Min = errors.New("field @ must be greater than or equal to the minimum value")`, "@", m.FieldName())
}

func (m *minValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@Min", "@", m.FieldName())
}

func ValidateMin(pass *codegen.Pass, field *ast.Field, expressions map[string]string) Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}
	minValue, ok := expressions[markers.GoValidMarkerMin]

	return &minValidator{
		pass:     pass,
		field:    field,
		minValue: minValue,
	}
}
