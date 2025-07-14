package main

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/validator/rules"
	"github.com/sivchari/govalid/internal/markers"
)

func testFilterMap(name, expr, fieldType string) {
	fmt.Printf("Testing: %s\n", name)
	fmt.Printf("Expression: %s\n", expr)
	
	field := &ast.Field{
		Names: []*ast.Ident{{Name: "TestField"}},
		Type:  &ast.Ident{Name: fieldType},
	}
	
	exprMap := map[string]string{
		markers.GoValidMarkerCEL: expr,
	}
	
	validator := rules.ValidateCEL(&codegen.Pass{}, field, exprMap)
	if validator == nil {
		fmt.Printf("❌ ValidateCEL returned nil\n")
		return
	}
	
	result := validator.Validate()
	if strings.Contains(result, "CEL conversion failed") {
		fmt.Printf("❌ FAILED: %s\n", result)
	} else if result == "!(true)" {
		fmt.Printf("⚠️  FALLBACK: Not implemented\n")
	} else {
		fmt.Printf("✅ SUCCESS: %s\n", result)
	}
	fmt.Println()
}

func main() {
	fmt.Printf("Testing CEL Filter and Map Comprehensions:\n")
	fmt.Printf("==========================================\n\n")

	// Test filter and map operations
	testFilterMap("Simple filter", "value.filter(item, item.startsWith('prefix'))", "[]string")
	testFilterMap("Simple map", "value.map(item, size(item))", "[]string")
	
	// Test filter with size check
	testFilterMap("Filter with size", "size(value.filter(item, item.startsWith('prefix'))) > 0", "[]string")
	testFilterMap("Map with size", "size(value.map(item, size(item))) == size(value)", "[]string")
}