// Package validatorhelper provides helper functions for the govalid validator.
package validatorhelper

import (
	"go/types"
)

// Zero returns the zero value for the given type.
func Zero(typ types.Type) string {
	switch t := typ.(type) {
	case *types.Basic:
		return zeroOfBasic(t)
	case *types.Pointer, *types.Interface, *types.Signature:
		return "nil"
	case *types.Alias, *types.Named:
		if underlying := types.Unalias(t).Underlying(); underlying != nil {
			return Zero(underlying)
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
