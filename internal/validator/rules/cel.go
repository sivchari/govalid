// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/google/cel-go/cel"
	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

type celValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	expression string
}

var _ validator.Validator = (*celValidator)(nil)

const (
	celKey            = "%s-cel"
	trueFallback      = "true"
	placeholderList   = "[]interface{}{}"
	placeholderStruct = "struct{}{}"
)

func (c *celValidator) Validate() string {
	fieldName := c.FieldName()
	// Convert CEL expression to Go expression at generation time
	goExpr, err := c.convertCELToGo(c.expression, fieldName)
	if err != nil {
		// Fallback to comment if conversion fails
		return fmt.Sprintf("true /* CEL conversion failed: %v */", err)
	}
	// Return the converted Go expression wrapped in negation for validation
	return fmt.Sprintf("!(%s)", goExpr)
}

func (c *celValidator) FieldName() string {
	return c.field.Names[0].Name
}

func (c *celValidator) Err() string {
	fieldName := c.FieldName()

	var result strings.Builder

	// Generate error variable only once per field
	if validator.GeneratorMemory[fmt.Sprintf(celKey, fieldName)] {
		return result.String()
	}

	validator.GeneratorMemory[fmt.Sprintf(celKey, fieldName)] = true

	// Generate error variable with the CEL expression included for debugging
	result.WriteString(strings.ReplaceAll(`
	// Err@CELValidation is the error returned when the CEL expression evaluation fails.
	Err@CELValidation = errors.New("field @ failed CEL validation: EXPRESSION")`, "@", fieldName))

	// Replace EXPRESSION placeholder with the actual CEL expression
	errorString := result.String()
	errorString = strings.ReplaceAll(errorString, "EXPRESSION", c.expression)

	return errorString
}

func (c *celValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@CELValidation", "@", c.FieldName())
}

func (c *celValidator) Imports() []string {
	imports := []string{}

	// Add imports based on the CEL expression content
	if strings.Contains(c.expression, "contains(") {
		imports = append(imports, "strings")
	}

	if strings.Contains(c.expression, "matches(") {
		imports = append(imports, "regexp")
	}

	return imports
}

// ValidateCEL creates a new celValidator for fields with CEL marker.
// This validator supports all field types since CEL can handle various data types.
func ValidateCEL(pass *codegen.Pass, field *ast.Field, expressions map[string]string) validator.Validator {
	celExpression, ok := expressions[markers.GoValidMarkerCEL]
	if !ok {
		return nil
	}

	// CEL expressions must not be empty
	if strings.TrimSpace(celExpression) == "" {
		return nil
	}

	return &celValidator{
		pass:       pass,
		field:      field,
		expression: celExpression,
	}
}

// convertCELToGo converts a CEL expression to equivalent Go code.
func (c *celValidator) convertCELToGo(celExpr, fieldName string) (string, error) {
	// Create a CEL environment to parse the expression
	env, err := cel.NewEnv(
		cel.StdLib(),
		cel.Variable("value", cel.DynType),
		cel.Variable("this", cel.DynType),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create CEL environment: %w", err)
	}

	// Parse the CEL expression
	ast, issues := env.Compile(celExpr)
	if issues != nil && issues.Err() != nil {
		return "", fmt.Errorf("failed to compile CEL expression: %w", issues.Err())
	}

	// Convert CEL AST to Go expression string
	// Use the parsed AST directly to avoid deprecated methods
	//nolint:staticcheck // ast.Expr() is deprecated but still functional
	goExpr := c.convertASTToGo(ast.Expr(), fieldName)

	return goExpr, nil
}

// convertASTToGo recursively converts CEL AST nodes to Go expression strings.
func (c *celValidator) convertASTToGo(expr *exprpb.Expr, fieldName string) string {
	switch expr.ExprKind.(type) {
	case *exprpb.Expr_IdentExpr:
		return c.convertIdentExpr(expr.GetIdentExpr(), fieldName)
	case *exprpb.Expr_SelectExpr:
		return c.convertSelectExpr(expr.GetSelectExpr(), fieldName)
	case *exprpb.Expr_CallExpr:
		return c.convertCallToGo(expr.GetCallExpr(), fieldName)
	case *exprpb.Expr_ConstExpr:
		return c.convertConstToGo(expr.GetConstExpr())
	case *exprpb.Expr_ListExpr:
		return placeholderList // placeholder
	case *exprpb.Expr_StructExpr:
		return placeholderStruct // placeholder
	case *exprpb.Expr_ComprehensionExpr:
		return trueFallback // placeholder
	default:
		return trueFallback // fallback
	}
}

// convertIdentExpr converts identifier expressions.
func (c *celValidator) convertIdentExpr(ident *exprpb.Expr_Ident, fieldName string) string {
	switch ident.Name {
	case "value":
		return fmt.Sprintf("t.%s", fieldName)
	case "this":
		return "t"
	default:
		return ident.Name
	}
}

// convertSelectExpr converts select expressions (field access).
func (c *celValidator) convertSelectExpr(selectExpr *exprpb.Expr_Select, fieldName string) string {
	operand := c.convertASTToGo(selectExpr.Operand, fieldName)

	if operand == "t" {
		return fmt.Sprintf("t.%s", selectExpr.Field)
	}

	return fmt.Sprintf("%s.%s", operand, selectExpr.Field)
}

// convertCallToGo converts CEL function calls to Go expressions.
func (c *celValidator) convertCallToGo(callExpr *exprpb.Expr_Call, fieldName string) string {
	function := callExpr.Function
	args := callExpr.Args

	// Try operators first
	if result := c.convertOperator(function, args, fieldName); result != "" {
		return result
	}

	// Try built-in functions
	if result := c.convertBuiltinFunction(function, args, fieldName); result != "" {
		return result
	}

	// Fallback for unknown functions
	return trueFallback
}

// convertOperator converts CEL operators to Go operators.
func (c *celValidator) convertOperator(function string, args []*exprpb.Expr, fieldName string) string {
	if len(args) != 2 {
		return ""
	}

	left := c.convertASTToGo(args[0], fieldName)
	right := c.convertASTToGo(args[1], fieldName)

	// Try logical operators first
	if result := c.convertLogicalOperator(function, left, right); result != "" {
		return result
	}

	// Try comparison operators
	if result := c.convertComparisonOperator(function, left, right); result != "" {
		return result
	}

	// Try arithmetic operators
	return c.convertArithmeticOperator(function, left, right)
}

// convertLogicalOperator converts logical operators.
func (c *celValidator) convertLogicalOperator(function, left, right string) string {
	switch function {
	case "_&&_":
		return fmt.Sprintf("(%s) && (%s)", left, right)
	case "_||_":
		return fmt.Sprintf("(%s) || (%s)", left, right)
	default:
		return ""
	}
}

// convertComparisonOperator converts comparison operators.
func (c *celValidator) convertComparisonOperator(function, left, right string) string {
	switch function {
	case "_>_":
		return fmt.Sprintf("%s > %s", left, right)
	case "_>=_":
		return fmt.Sprintf("%s >= %s", left, right)
	case "_<_":
		return fmt.Sprintf("%s < %s", left, right)
	case "_<=_":
		return fmt.Sprintf("%s <= %s", left, right)
	case "_==_":
		return fmt.Sprintf("%s == %s", left, right)
	case "_!=_":
		return fmt.Sprintf("%s != %s", left, right)
	default:
		return ""
	}
}

// convertArithmeticOperator converts arithmetic operators.
func (c *celValidator) convertArithmeticOperator(function, left, right string) string {
	switch function {
	case "_+_":
		return fmt.Sprintf("%s + %s", left, right)
	case "_-_":
		return fmt.Sprintf("%s - %s", left, right)
	case "_*_":
		return fmt.Sprintf("%s * %s", left, right)
	case "_/_":
		return fmt.Sprintf("%s / %s", left, right)
	default:
		return ""
	}
}

// convertBuiltinFunction converts CEL built-in functions to Go equivalents.
func (c *celValidator) convertBuiltinFunction(function string, args []*exprpb.Expr, fieldName string) string {
	switch function {
	case "size":
		if len(args) == 1 {
			arg := c.convertASTToGo(args[0], fieldName)

			return fmt.Sprintf("len(%s)", arg)
		}
	case "contains":
		if len(args) == 2 {
			str := c.convertASTToGo(args[0], fieldName)
			substr := c.convertASTToGo(args[1], fieldName)

			return fmt.Sprintf("strings.Contains(%s, %s)", str, substr)
		}
	case "matches":
		if len(args) == 2 {
			str := c.convertASTToGo(args[0], fieldName)
			pattern := c.convertASTToGo(args[1], fieldName)

			return fmt.Sprintf("regexp.MustCompile(%s).MatchString(%s)", pattern, str)
		}
	}

	return ""
}

// convertConstToGo converts CEL constants to Go literals.
func (c *celValidator) convertConstToGo(constExpr *exprpb.Constant) string {
	switch constExpr.ConstantKind.(type) {
	case *exprpb.Constant_BoolValue:
		if constExpr.GetBoolValue() {
			return "true"
		}

		return "false"
	case *exprpb.Constant_Int64Value:
		return fmt.Sprintf("%d", constExpr.GetInt64Value())
	case *exprpb.Constant_Uint64Value:
		return fmt.Sprintf("%d", constExpr.GetUint64Value())
	case *exprpb.Constant_DoubleValue:
		return fmt.Sprintf("%g", constExpr.GetDoubleValue())
	case *exprpb.Constant_StringValue:
		return fmt.Sprintf("%q", constExpr.GetStringValue())
	case *exprpb.Constant_BytesValue:
		return fmt.Sprintf("%q", constExpr.GetBytesValue())
	default:
		return "nil"
	}
}
