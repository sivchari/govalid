// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"

	"github.com/google/cel-go/cel"
	"github.com/gostaticanalysis/codegen"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"

	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator"
)

type celValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	expression string
	structName string
}

var _ validator.Validator = (*celValidator)(nil)

const (
	celKey            = "%s-cel"
	trueFallback      = "true"
	placeholderList   = "[]interface{}{}"
	placeholderStruct = "struct{}{}"
	ternaryOperator   = "_?_:_"
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
	key := fmt.Sprintf(celKey, c.structName+fieldName)
	if validator.GeneratorMemory[key] {
		return result.String()
	}

	validator.GeneratorMemory[key] = true

	result.WriteString(
		strings.ReplaceAll(`
			// Err@CELValidation is the error returned when the CEL expression evaluation fails.
			Err@CELValidation = govaliderrors.ValidationError{Reason:"field @ failed CEL validation: EXPRESSION",Path:"PATH"}
			`,
			"@",
			c.structName+c.FieldName(),
		),
	)

	// Replace EXPRESSION placeholder with the actual CEL expression
	errorString := result.String()
	errorString = strings.ReplaceAll(errorString, "EXPRESSION", c.expression)

	return strings.ReplaceAll(result.String(), "PATH", fmt.Sprintf("%s.%s", c.structName, c.FieldName()))
}

func (c *celValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@CELValidation", "@", c.structName+c.FieldName())
}

func (c *celValidator) Imports() []string {
	imports := []string{}

	// Add imports based on the CEL expression content
	if c.needsStringsImport() {
		imports = append(imports, "strings")
	}

	if strings.Contains(c.expression, "matches(") {
		imports = append(imports, "regexp")
	}

	if c.needsStrconvImport() {
		imports = append(imports, "strconv")
	}

	if c.needsFmtImport() {
		imports = append(imports, "fmt")
	}

	if c.needsTimeImport() {
		imports = append(imports, "time")
	}

	// Add slices import for optimized contains operations
	if strings.Contains(c.expression, " in ") {
		imports = append(imports, "slices")
	}

	return imports
}

func (c *celValidator) needsStringsImport() bool {
	return strings.Contains(c.expression, "contains(") ||
		strings.Contains(c.expression, "startsWith(") ||
		strings.Contains(c.expression, "endsWith(")
}

func (c *celValidator) needsStrconvImport() bool {
	return strings.Contains(c.expression, "int(") ||
		strings.Contains(c.expression, "double(")
}

func (c *celValidator) needsFmtImport() bool {
	return strings.Contains(c.expression, "string(") ||
		strings.Contains(c.expression, "double(")
}

func (c *celValidator) needsTimeImport() bool {
	return strings.Contains(c.expression, "timestamp(") ||
		strings.Contains(c.expression, "duration(")
}

// ValidateCEL creates a new celValidator for fields with CEL marker.
// This validator supports all field types since CEL can handle various data types.
func ValidateCEL(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName string) validator.Validator {
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
		structName: structName,
	}
}

// convertCELToGo converts a CEL expression to equivalent Go code.
func (c *celValidator) convertCELToGo(celExpr, fieldName string) (string, error) {
	// Pre-validate that this is a standard CEL expression
	if err := c.validateStandardCEL(celExpr); err != nil {
		return "", fmt.Errorf("non-standard CEL expression: %w", err)
	}

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
		return c.convertListExpr(expr.GetListExpr(), fieldName)
	case *exprpb.Expr_StructExpr:
		return placeholderStruct // placeholder
	case *exprpb.Expr_ComprehensionExpr:
		return c.convertComprehensionExpr(expr.GetComprehensionExpr(), fieldName)
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

	// Handle method calls with a target (e.g., value.startsWith('prefix_'))
	if callExpr.Target != nil {
		return c.convertMethodCall(function, callExpr.Target, args, fieldName)
	}

	// Try operators first
	if result := c.convertOperator(function, args, fieldName); result != "" {
		return result
	}

	// Try built-in functions
	if result := c.convertBuiltinFunction(function, fieldName, args); result != "" {
		return result
	}

	// Fallback for unknown functions
	return trueFallback
}

// convertOperator converts CEL operators to Go operators.
func (c *celValidator) convertOperator(function string, args []*exprpb.Expr, fieldName string) string {
	// Handle ternary operator (condition ? true_value : false_value)
	if function == ternaryOperator && len(args) == 3 {
		return c.convertTernaryOperator(args, fieldName)
	}

	// Handle 'in' operator
	if function == "@in" && len(args) == 2 {
		return c.convertInOperator(args, fieldName)
	}

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

func (c *celValidator) convertTernaryOperator(args []*exprpb.Expr, fieldName string) string {
	condition := c.convertASTToGo(args[0], fieldName)
	trueValue := c.convertASTToGo(args[1], fieldName)
	falseValue := c.convertASTToGo(args[2], fieldName)
	// Return the ternary result - it will be used in comparison context
	// Since this is used in validation context, we need to handle it properly
	return fmt.Sprintf("func() int { if %s { return %s }; return %s }() > 0", condition, trueValue, falseValue)
}

func (c *celValidator) convertInOperator(args []*exprpb.Expr, fieldName string) string {
	element := c.convertASTToGo(args[0], fieldName)
	collection := c.convertASTToGo(args[1], fieldName)

	// Optimize for string literal slices using slices.Contains
	if strings.HasPrefix(collection, "[]interface{}{") && strings.HasSuffix(collection, "}") {
		if optimized := c.optimizeStringSliceContains(collection, element); optimized != "" {
			return optimized
		}
	}

	// Optimize for field contains literal string: field contains "literal"
	if strings.HasPrefix(collection, "t.") && strings.HasPrefix(element, "\"") && strings.HasSuffix(element, "\"") {
		return fmt.Sprintf("slices.Contains(%s, %s)", collection, element)
	}

	// Generate a contains check for slices
	return fmt.Sprintf("func() bool { for _, item := range %s { if item == %s { return true } }; return false }()", collection, element)
}

func (c *celValidator) optimizeStringSliceContains(collection, element string) string {
	// Extract elements from []interface{}{"a", "b", "c"}
	content := strings.TrimPrefix(collection, "[]interface{}{")
	content = strings.TrimSuffix(content, "}")

	// Check if all elements are string literals
	if !strings.Contains(content, "\"") {
		return ""
	}

	// Convert to string slice for optimization
	stringSlice := fmt.Sprintf("[]string{%s}", content)

	// If element is fmt.Sprintf("%v", field), optimize for string fields
	if strings.Contains(element, "fmt.Sprintf(\"%%v\", ") {
		if fieldName := c.extractFieldFromFmtSprintf(element); fieldName != "" {
			return fmt.Sprintf("slices.Contains(%s, %s)", stringSlice, fieldName)
		}
	}

	return fmt.Sprintf("slices.Contains(%s, %s)", stringSlice, element)
}

func (c *celValidator) extractFieldFromFmtSprintf(element string) string {
	// Extract the field name from fmt.Sprintf("%v", t.Field)
	fieldStart := strings.Index(element, "t.")
	if fieldStart == -1 {
		return ""
	}

	fieldEnd := strings.Index(element[fieldStart:], ")")
	if fieldEnd == -1 {
		return ""
	}

	return element[fieldStart : fieldStart+fieldEnd]
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
func (c *celValidator) convertBuiltinFunction(function, fieldName string, args []*exprpb.Expr) string {
	switch function {
	case "size":
		return c.convertSizeFunction(fieldName, args)
	case "contains":
		return c.convertContainsFunction(fieldName, args)
	case "matches":
		return c.convertMatchesFunction(fieldName, args)
	case "startsWith":
		return c.convertStartsWithFunction(fieldName, args)
	case "endsWith":
		return c.convertEndsWithFunction(fieldName, args)
	case "int":
		return c.convertIntFunction(fieldName, args)
	case "string":
		return c.convertStringFunction(fieldName, args)
	case "double":
		return c.convertDoubleFunction(fieldName, args)
	case "timestamp":
		return c.convertTimestampFunction(fieldName, args)
	case "duration":
		return c.convertDurationFunction(fieldName, args)
	}

	return ""
}

func (c *celValidator) convertSizeFunction(fieldName string, args []*exprpb.Expr) string {
	if len(args) != 1 {
		return ""
	}

	arg := c.convertASTToGo(args[0], fieldName)

	return fmt.Sprintf("len(%s)", arg)
}

func (c *celValidator) convertContainsFunction(fieldName string, args []*exprpb.Expr) string {
	if len(args) != 2 {
		return ""
	}

	str := c.convertASTToGo(args[0], fieldName)
	substr := c.convertASTToGo(args[1], fieldName)

	return fmt.Sprintf("strings.Contains(%s, %s)", str, substr)
}

func (c *celValidator) convertMatchesFunction(fieldName string, args []*exprpb.Expr) string {
	if len(args) != 2 {
		return ""
	}

	str := c.convertASTToGo(args[0], fieldName)
	pattern := c.convertASTToGo(args[1], fieldName)

	return fmt.Sprintf("regexp.MustCompile(%s).MatchString(%s)", pattern, str)
}

func (c *celValidator) convertStartsWithFunction(fieldName string, args []*exprpb.Expr) string {
	if len(args) != 2 {
		return ""
	}

	str := c.convertASTToGo(args[0], fieldName)
	prefix := c.convertASTToGo(args[1], fieldName)

	return fmt.Sprintf("strings.HasPrefix(%s, %s)", str, prefix)
}

func (c *celValidator) convertEndsWithFunction(fieldName string, args []*exprpb.Expr) string {
	if len(args) != 2 {
		return ""
	}

	str := c.convertASTToGo(args[0], fieldName)
	suffix := c.convertASTToGo(args[1], fieldName)

	return fmt.Sprintf("strings.HasSuffix(%s, %s)", str, suffix)
}

func (c *celValidator) convertIntFunction(fieldName string, args []*exprpb.Expr) string {
	if len(args) != 1 {
		return ""
	}

	arg := c.convertASTToGo(args[0], fieldName)

	return fmt.Sprintf("func() int { v, err := strconv.Atoi(%s); if err != nil { return 0 }; return v }()", arg)
}

func (c *celValidator) convertStringFunction(fieldName string, args []*exprpb.Expr) string {
	if len(args) != 1 {
		return ""
	}

	arg := c.convertASTToGo(args[0], fieldName)

	return fmt.Sprintf("fmt.Sprintf(\"%%v\", %s)", arg)
}

func (c *celValidator) convertDoubleFunction(fieldName string, args []*exprpb.Expr) string {
	if len(args) != 1 {
		return ""
	}

	arg := c.convertASTToGo(args[0], fieldName)

	return fmt.Sprintf("func() float64 { v, err := strconv.ParseFloat(fmt.Sprintf(\"%%v\", %s), 64); if err != nil { return 0.0 }; return v }()", arg)
}

func (c *celValidator) convertTimestampFunction(fieldName string, args []*exprpb.Expr) string {
	if len(args) != 1 {
		return ""
	}

	arg := c.convertASTToGo(args[0], fieldName)

	return fmt.Sprintf("func() time.Time { t, err := time.Parse(time.RFC3339, %s); if err != nil { return time.Time{} }; return t }()", arg)
}

func (c *celValidator) convertDurationFunction(fieldName string, args []*exprpb.Expr) string {
	if len(args) != 1 {
		return ""
	}

	arg := c.convertASTToGo(args[0], fieldName)

	return fmt.Sprintf("func() time.Duration { d, err := time.ParseDuration(%s); if err != nil { return 0 }; return d }()", arg)
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

// convertMethodCall converts method calls with a target (e.g., value.startsWith('prefix_')).
func (c *celValidator) convertMethodCall(method string, target *exprpb.Expr, args []*exprpb.Expr, fieldName string) string {
	targetStr := c.convertASTToGo(target, fieldName)

	switch method {
	case "startsWith":
		if len(args) == 1 {
			prefix := c.convertASTToGo(args[0], fieldName)

			return fmt.Sprintf("strings.HasPrefix(%s, %s)", targetStr, prefix)
		}
	case "endsWith":
		if len(args) == 1 {
			suffix := c.convertASTToGo(args[0], fieldName)

			return fmt.Sprintf("strings.HasSuffix(%s, %s)", targetStr, suffix)
		}
	case "contains":
		if len(args) == 1 {
			substr := c.convertASTToGo(args[0], fieldName)

			return fmt.Sprintf("strings.Contains(%s, %s)", targetStr, substr)
		}
	case "matches":
		if len(args) == 1 {
			pattern := c.convertASTToGo(args[0], fieldName)

			return fmt.Sprintf("regexp.MustCompile(%s).MatchString(%s)", pattern, targetStr)
		}
	}

	// Fallback for unknown method calls
	return trueFallback
}

// convertListExpr converts list expressions like ['a', 'b', 'c'].
func (c *celValidator) convertListExpr(listExpr *exprpb.Expr_CreateList, fieldName string) string {
	elements := make([]string, len(listExpr.Elements))
	for i, element := range listExpr.Elements {
		elements[i] = c.convertASTToGo(element, fieldName)
	}

	return fmt.Sprintf("[]interface{}{%s}", strings.Join(elements, ", "))
}

// validateStandardCEL validates that the expression uses only standard CEL features.
func (c *celValidator) validateStandardCEL(celExpr string) error {
	// List of non-standard patterns to reject
	nonStandardPatterns := []string{
		// String methods not in CEL standard
		".split(", ".trim(", ".replace(", ".substring(",
		".toLowerCase(", ".toUpperCase(",

		// Math functions not in CEL standard
		"math.abs(", "math.min(", "math.max(", "math.floor(", "math.ceil(",

		// Advanced syntax not in CEL standard
		"?.", "?:", "try(", "..", "range(",

		// List/Map methods not in CEL standard
		".reverse(", ".sort(", ".unique(", ".keys(", ".values(",

		// String interpolation syntax not in standard CEL
		"${",
	}

	for _, pattern := range nonStandardPatterns {
		if strings.Contains(celExpr, pattern) {
			return fmt.Errorf("expression contains non-standard CEL feature: %s", pattern)
		}
	}

	// Check for slice notation like [1:3] which is non-standard
	if strings.Contains(celExpr, ":") && strings.Contains(celExpr, "[") && strings.Contains(celExpr, "]") {
		// Look for pattern like [number:number] which indicates slice notation
		slicePattern := regexp.MustCompile(`\[\s*\d+\s*:\s*\d*\s*\]`)
		if slicePattern.MatchString(celExpr) {
			return fmt.Errorf("expression contains non-standard slice notation")
		}
	}

	return nil
}

// convertComprehensionExpr converts CEL list comprehension expressions to Go equivalents.
func (c *celValidator) convertComprehensionExpr(comprExpr *exprpb.Expr_Comprehension, fieldName string) string {
	// Get the iteration variable and collection
	iterVar := comprExpr.IterVar
	iterRange := c.convertASTToGo(comprExpr.IterRange, fieldName)

	// Determine the type of comprehension based on the structure
	return c.generateComprehensionGo(comprExpr, iterVar, iterRange, fieldName)
}

// generateComprehensionGo generates Go code for different types of comprehensions.
func (c *celValidator) generateComprehensionGo(comprExpr *exprpb.Expr_Comprehension, iterVar, iterRange, fieldName string) string {
	// Get the loop condition and step to determine comprehension type
	accuInit := c.convertASTToGo(comprExpr.AccuInit, fieldName)
	loopStep := c.convertASTToGo(comprExpr.LoopStep, fieldName)

	// Determine comprehension type based on initial accumulator
	switch {
	case accuInit == "true":
		// items.all(item, condition) - starts with true, AND operation
		return c.generateAllComprehension(comprExpr, iterVar, iterRange, fieldName)
	case accuInit == "false":
		// items.exists(item, condition) - starts with false, OR operation
		return c.generateExistsComprehension(comprExpr, iterVar, iterRange, fieldName)
	case strings.Contains(accuInit, "[]interface{}"):
		// items.filter(item, condition) or items.map(item, transform) - starts with empty list
		if strings.Contains(loopStep, "func() int { if") {
			return c.generateFilterComprehension(comprExpr, iterVar, iterRange, fieldName)
		}

		return c.generateMapComprehension(comprExpr, iterVar, iterRange, fieldName)
	default:
		// Could be exists_one or other complex comprehension
		return c.generateExistsOneComprehension(comprExpr, iterVar, iterRange, fieldName)
	}
}

// generateAllComprehension generates Go code for items.all(item, condition).
func (c *celValidator) generateAllComprehension(comprExpr *exprpb.Expr_Comprehension, iterVar, iterRange, fieldName string) string {
	// Extract the actual condition from loop step (it's the right side of &&)
	condition := c.extractConditionFromAndStep(comprExpr.LoopStep, fieldName)

	return fmt.Sprintf("func() bool { for _, %s := range %s { if !(%s) { return false } }; return true }()",
		iterVar, iterRange, condition)
}

// generateExistsComprehension generates Go code for items.exists(item, condition).
func (c *celValidator) generateExistsComprehension(comprExpr *exprpb.Expr_Comprehension, iterVar, iterRange, fieldName string) string {
	// Extract the actual condition from loop step (it's the right side of ||)
	condition := c.extractConditionFromOrStep(comprExpr.LoopStep, fieldName)

	return fmt.Sprintf("func() bool { for _, %s := range %s { if %s { return true } }; return false }()",
		iterVar, iterRange, condition)
}

// generateFilterComprehension generates Go code for items.filter(item, condition).
func (c *celValidator) generateFilterComprehension(comprExpr *exprpb.Expr_Comprehension, iterVar, iterRange, fieldName string) string {
	// Extract the condition from the ternary operation in loop step
	condition := c.extractConditionFromTernary(comprExpr.LoopStep, iterVar, fieldName)

	return fmt.Sprintf("func() []interface{} { var result []interface{}; for _, %s := range %s { if %s { result = append(result, %s) } }; return result }()",
		iterVar, iterRange, condition, iterVar)
}

// generateMapComprehension generates Go code for items.map(item, transform).
func (c *celValidator) generateMapComprehension(comprExpr *exprpb.Expr_Comprehension, iterVar, iterRange, fieldName string) string {
	// Extract the transformation from loop step
	transform := c.extractTransformFromLoopStep(comprExpr.LoopStep, iterVar, fieldName)

	return fmt.Sprintf("func() []interface{} { var result []interface{}; for _, %s := range %s { result = append(result, %s) }; return result }()",
		iterVar, iterRange, transform)
}

// generateExistsOneComprehension generates Go code for items.exists_one(item, condition).
func (c *celValidator) generateExistsOneComprehension(comprExpr *exprpb.Expr_Comprehension, iterVar, iterRange, fieldName string) string {
	// For exists_one, the condition is embedded in a ternary within loop step
	condition := c.extractConditionFromExistsOneTernary(comprExpr.LoopStep, iterVar, fieldName)

	return fmt.Sprintf("func() bool { count := 0; for _, %s := range %s { if %s { count++; if count > 1 { return false } } }; return count == 1 }()",
		iterVar, iterRange, condition)
}

func (c *celValidator) extractConditionFromTernary(loopStep *exprpb.Expr, iterVar, fieldName string) string {
	// Extract condition from ternary operation used in filter
	if callExpr := loopStep.GetCallExpr(); callExpr != nil {
		function := callExpr.Function
		args := callExpr.Args

		// For ternary operations _?_:_, the first arg is the condition
		if function == ternaryOperator && len(args) >= 1 {
			return c.convertASTToGo(args[0], fieldName)
		}
	}

	// For function literal with ternary, parse the structure differently
	loopStepStr := c.convertASTToGo(loopStep, fieldName)
	if strings.Contains(loopStepStr, "if ") && strings.Contains(loopStepStr, " { return") {
		// Extract condition from "func() int { if condition { return ... }; return @result }()"
		start := strings.Index(loopStepStr, "if ") + 3
		end := strings.Index(loopStepStr[start:], " { return")

		if start < len(loopStepStr) && end > 0 {
			return strings.TrimSpace(loopStepStr[start : start+end])
		}
	}

	// Fallback condition
	return fmt.Sprintf("%s != nil", iterVar)
}

func (c *celValidator) extractTransformFromLoopStep(loopStep *exprpb.Expr, iterVar, fieldName string) string {
	// Extract transformation from loop step for map operations
	if callExpr := loopStep.GetCallExpr(); callExpr != nil {
		function := callExpr.Function
		args := callExpr.Args

		// For list append operations in map, look for the value being appended
		if function == "_+_" && len(args) >= 2 {
			// Usually the second argument is the transform
			if listExpr := args[1].GetListExpr(); listExpr != nil && len(listExpr.Elements) > 0 {
				return c.convertASTToGo(listExpr.Elements[0], fieldName)
			}
		}
	}

	// Parse from string representation for map operations like "@result + []interface{}{len(item)}"
	loopStepStr := c.convertASTToGo(loopStep, fieldName)
	if strings.Contains(loopStepStr, "+ []interface{}{") {
		start := strings.Index(loopStepStr, "+ []interface{}{") + len("+ []interface{}{")
		end := strings.LastIndex(loopStepStr, "}")

		if start < len(loopStepStr) && end > start {
			return strings.TrimSpace(loopStepStr[start:end])
		}
	}

	// Fallback transformation - convert item to itself
	return iterVar
}

// extractConditionFromAndStep extracts the condition from the right side of && operation.
func (c *celValidator) extractConditionFromAndStep(loopStep *exprpb.Expr, fieldName string) string {
	if callExpr := loopStep.GetCallExpr(); callExpr != nil {
		function := callExpr.Function
		args := callExpr.Args

		// For && operation, we want the right side (the actual condition)
		if function == "_&&_" && len(args) >= 2 {
			return c.convertASTToGo(args[1], fieldName)
		}
	}

	// Fallback - return the whole expression
	return c.convertASTToGo(loopStep, fieldName)
}

// extractConditionFromOrStep extracts the condition from the right side of || operation.
func (c *celValidator) extractConditionFromOrStep(loopStep *exprpb.Expr, fieldName string) string {
	if callExpr := loopStep.GetCallExpr(); callExpr != nil {
		function := callExpr.Function
		args := callExpr.Args

		// For || operation, we want the right side (the actual condition)
		if function == "_||_" && len(args) >= 2 {
			return c.convertASTToGo(args[1], fieldName)
		}
	}

	// Fallback - return the whole expression
	return c.convertASTToGo(loopStep, fieldName)
}

// extractConditionFromExistsOneTernary extracts condition from ternary in exists_one.
func (c *celValidator) extractConditionFromExistsOneTernary(loopStep *exprpb.Expr, iterVar, fieldName string) string {
	if callExpr := loopStep.GetCallExpr(); callExpr != nil {
		function := callExpr.Function
		args := callExpr.Args

		// For ternary operation, the condition is the first argument
		if function == ternaryOperator && len(args) >= 1 {
			return c.convertASTToGo(args[0], fieldName)
		}
	}

	// Fallback
	return fmt.Sprintf("%s != nil", iterVar)
}
