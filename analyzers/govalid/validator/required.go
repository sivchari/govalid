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
	"go/ast"
	"strings"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/validator"
)

type requiredValidator struct {
	pass  *codegen.Pass
	field *ast.Field
}

var _ Validator = (*requiredValidator)(nil)

func (r *requiredValidator) Validate() string {
	typ := r.pass.TypesInfo.TypeOf(r.field.Type)

	return validator.ZeroExpr(r.FieldName(), typ)
}

func (r *requiredValidator) FieldName() string {
	return r.field.Names[0].Name
}

func (r *requiredValidator) Err() string {
	if generatorMemory[r.FieldName()] {
		return ""
	}

	generatorMemory[r.FieldName()] = true

	return strings.ReplaceAll(`
	// Err @ is returned when the @ is required but not provided.
	Err@Required = errors.New("field @ is required")`, "@", r.FieldName())
}

func (r *requiredValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@Required", "@", r.FieldName())
}

// ValidateRequired creates a new required validator for the given field.
func ValidateRequired(pass *codegen.Pass, field *ast.Field) Validator {
	generatorMemory[field.Names[0].Name] = false

	return &requiredValidator{
		pass:  pass,
		field: field,
	}
}
