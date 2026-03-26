// Package expr provides helper functions for constructing Go AST expression nodes.
// These helpers reduce the verbosity of building go/ast nodes directly,
// which typically requires 10-15 lines per expression.
package expr

import (
	"fmt"
	"go/ast"
	"go/token"
)

// Field creates a selector expression representing receiver.name (e.g., t.FieldName).
func Field(receiver, name string) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X:   ast.NewIdent(receiver),
		Sel: ast.NewIdent(name),
	}
}

// Ident creates an identifier expression (e.g., a variable name).
func Ident(name string) *ast.Ident {
	return ast.NewIdent(name)
}

// IntLit creates an integer literal expression from a string value (e.g., "42").
func IntLit(value string) *ast.BasicLit {
	return &ast.BasicLit{Kind: token.INT, Value: value}
}

// StringLit creates a double-quoted string literal expression (e.g., "hello" becomes `"hello"`).
func StringLit(value string) *ast.BasicLit {
	return &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", value)}
}

// Nil creates a nil identifier expression.
func Nil() *ast.Ident {
	return ast.NewIdent("nil")
}

// Not creates a unary NOT expression (!x).
func Not(x ast.Expr) *ast.UnaryExpr {
	return &ast.UnaryExpr{Op: token.NOT, X: x}
}

// Paren creates a parenthesized expression ((x)).
func Paren(x ast.Expr) *ast.ParenExpr {
	return &ast.ParenExpr{X: x}
}

// Binary creates a binary expression (left op right).
// This is the foundation for comparison and logical operator helpers.
func Binary(left ast.Expr, op token.Token, right ast.Expr) *ast.BinaryExpr {
	return &ast.BinaryExpr{X: left, Op: op, Y: right}
}

// GT creates a greater-than comparison (left > right).
func GT(left, right ast.Expr) *ast.BinaryExpr {
	return Binary(left, token.GTR, right)
}

// GTE creates a greater-than-or-equal comparison (left >= right).
func GTE(left, right ast.Expr) *ast.BinaryExpr {
	return Binary(left, token.GEQ, right)
}

// LT creates a less-than comparison (left < right).
func LT(left, right ast.Expr) *ast.BinaryExpr {
	return Binary(left, token.LSS, right)
}

// LTE creates a less-than-or-equal comparison (left <= right).
func LTE(left, right ast.Expr) *ast.BinaryExpr {
	return Binary(left, token.LEQ, right)
}

// Eq creates an equality comparison (left == right).
func Eq(left, right ast.Expr) *ast.BinaryExpr {
	return Binary(left, token.EQL, right)
}

// NEq creates an inequality comparison (left != right).
func NEq(left, right ast.Expr) *ast.BinaryExpr {
	return Binary(left, token.NEQ, right)
}

// And creates a logical AND expression (left && right).
func And(left, right ast.Expr) *ast.BinaryExpr {
	return Binary(left, token.LAND, right)
}

// Or creates a logical OR expression (left || right).
func Or(left, right ast.Expr) *ast.BinaryExpr {
	return Binary(left, token.LOR, right)
}

// AndAll chains multiple expressions with logical AND (&&), left-associatively.
// Returns nil if no expressions are provided.
func AndAll(exprs ...ast.Expr) ast.Expr {
	if len(exprs) == 0 {
		return nil
	}

	result := exprs[0]
	for _, e := range exprs[1:] {
		result = And(result, e)
	}

	return result
}

// Call creates a function call expression.
// If pkg is empty, it creates a simple call like len(args...).
// If pkg is non-empty, it creates a qualified call like pkg.fn(args...).
func Call(pkg, fn string, args ...ast.Expr) *ast.CallExpr {
	var funcExpr ast.Expr
	if pkg == "" {
		funcExpr = ast.NewIdent(fn)
	} else {
		funcExpr = &ast.SelectorExpr{
			X:   ast.NewIdent(pkg),
			Sel: ast.NewIdent(fn),
		}
	}

	return &ast.CallExpr{Fun: funcExpr, Args: args}
}

// MethodCall creates a method call expression like receiver.method(args...).
func MethodCall(receiver ast.Expr, method string, args ...ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: receiver, Sel: ast.NewIdent(method)},
		Args: args,
	}
}

// ShortVarDecl creates a short variable declaration statement (name := value).
func ShortVarDecl(name string, value ast.Expr) *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent(name)},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{value},
	}
}
