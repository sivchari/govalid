package validator

import (
	"fmt"
	"go/types"
)

func Zero(typ types.Type) string {
	switch t := typ.(type) {
	case *types.Basic:
		switch t.Kind() {
		case types.Bool:
			return "false"
		case types.Int, types.Int8, types.Int16, types.Int32, types.Int64:
			return "0"
		case types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
			return "0"
		case types.Float32, types.Float64:
			return "0.0"
		case types.Complex64, types.Complex128:
			return "0.0i"
		case types.String:
			return `""`
		default:
			return ""
		}
	case *types.Pointer, *types.Interface:
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

func ZeroExpr(name string, typ types.Type) string {
	zero := Zero(typ)
	if zero == "" {
		switch typ.(type) {
		case *types.Slice, *types.Array, *types.Map, *types.Chan:
			return fmt.Sprintf("len(t.%s) == 0", name)
		default:
			return ""
		}
	}
	return fmt.Sprintf("t.%s == %s", name, zero)
}
