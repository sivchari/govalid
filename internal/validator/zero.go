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

// Package validator implements the required validator for govalid.
package validator

import (
	"go/types"
)

// zero returns the zero value for the given type.
func zero(typ types.Type) string {
	switch t := typ.(type) {
	case *types.Basic:
		return zeroOfBasic(t)
	case *types.Pointer, *types.Interface, *types.Signature:
		return "nil"
	case *types.Alias, *types.Named:
		if underlying := types.Unalias(t).Underlying(); underlying != nil {
			return zero(underlying)
		}
		return ""
	default:
		return ""
	}
}

func zeroOfBasic(typ *types.Basic) string {
	switch typ.Kind() {
	case types.Bool, types.UntypedBool:
		return "false"
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64, types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64, types.Uintptr, types.UntypedInt, types.UntypedRune:
		return "0"
	case types.Float32, types.Float64, types.UntypedFloat:
		return "0.0"
	case types.Complex64, types.Complex128, types.UntypedComplex:
		return "0.0i"
	case types.String, types.UntypedString:
		return `""`
	case types.UnsafePointer, types.UntypedNil:
		return "nil"
	case types.Invalid:
		return ""
	default:
		return ""
	}
}
