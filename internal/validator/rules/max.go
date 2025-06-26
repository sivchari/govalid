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

type maxValidator struct {
	pass     *codegen.Pass
	field    *ast.Field
	maxValue string
}

var _ validator.Validator = (*maxValidator)(nil)

const maxKey = "%s-max"

func (m *maxValidator) Validate() string {
	return fmt.Sprintf("t.%s > %s", m.FieldName(), m.maxValue)
}

func (m *maxValidator) FieldName() string {
	return m.field.Names[0].Name
}

func (m *maxValidator) Err() string {
	if validator.GeneratorMemory[fmt.Sprintf(maxKey, m.FieldName())] {
		return ""
	}

	validator.GeneratorMemory[fmt.Sprintf(maxKey, m.FieldName())] = true

	return fmt.Sprintf(strings.ReplaceAll(`
	// Err@MaxValidation is the error returned when the value of the field is greater than the maximum value.
	Err@MaxValidation = errors.New("field @ must be less than %s")`, "@", m.FieldName()), m.maxValue)
}

func (m *maxValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@MaxValidation", "@", m.FieldName())
}

// ValidateMax creates a new maxValidator if the field type is numeric and the max marker is present.
func ValidateMax(pass *codegen.Pass, field *ast.Field, expressions map[string]string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}

	maxValue, ok := expressions[markers.GoValidMarkerMax]
	if !ok {
		return nil
	}

	return &maxValidator{
		pass:     pass,
		field:    field,
		maxValue: maxValue,
	}
}
