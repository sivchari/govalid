package expr

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

// Render converts an ast.Expr to its Go source code string representation.
// This is used at the template boundary to embed AST expressions into generated code.
func Render(e ast.Expr) string {
	var buf bytes.Buffer
	fset := token.NewFileSet()

	if err := format.Node(&buf, fset, e); err != nil {
		return ""
	}

	return buf.String()
}

// RenderStmt converts an ast.Stmt to its Go source code string representation.
func RenderStmt(s ast.Stmt) string {
	var buf bytes.Buffer
	fset := token.NewFileSet()

	if err := format.Node(&buf, fset, s); err != nil {
		return ""
	}

	return buf.String()
}

// Parse converts a Go expression string into an ast.Expr node.
// This is primarily used by the CEL validator to convert its string output to AST,
// and by the required validator for zero-value expressions.
func Parse(goExpr string) (ast.Expr, error) {
	return parser.ParseExpr(goExpr)
}
