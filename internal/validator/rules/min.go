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

// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator"
)

type minValidator struct {
	pass     *codegen.Pass
	field    *ast.Field
	minValue string
}

var _ validator.Validator = (*minValidator)(nil)

const minKey = "%s-min"

func (m *minValidator) Validate() string {
	return fmt.Sprintf("t.%s < %s", m.FieldName(), m.minValue)
}

func (m *minValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *minValidator) Err() string {
	if validator.GeneratorMemory[fmt.Sprintf(minKey, m.FieldName())] {
		return ""
	}

	validator.GeneratorMemory[fmt.Sprintf(minKey, m.FieldName())] = true

	return fmt.Sprintf(strings.ReplaceAll(`
	// Err@MinValidation is the error returned when the value of the field is less than the minimum value.
	Err@MinValidation = errors.New("field @ must be greater than %s")`, "@", m.FieldName()), m.minValue)
}

func (m *minValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@MinValidation", "@", m.FieldName())
}

// ValidateMin creates a new minValidator if the field type is numeric and the min marker is present.
func ValidateMin(pass *codegen.Pass, field *ast.Field, expressions map[string]string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}

	minValue, ok := expressions[markers.GoValidMarkerMin]
	if !ok {
		return nil
	}

	return &minValidator{
		pass:     pass,
		field:    field,
		minValue: minValue,
	}
}
