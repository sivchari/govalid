// Package generate provides functions for discovering and generating validator registry files.
package generate

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"sort"
	"strings"
)

// ValidatorInfo contains information about a discovered validator.
type ValidatorInfo struct {
	MarkerName   string // e.g., "required", "maxlength"
	FunctionName string // e.g., "ValidateRequired"
}

// DiscoverValidators finds all validator files in the rules directory and extracts their information.
func DiscoverValidators(rulesDir string) ([]ValidatorInfo, error) {
	files, err := filepath.Glob(filepath.Join(rulesDir, "*.go"))
	if err != nil {
		return nil, fmt.Errorf("failed to glob rules directory: %w", err)
	}

	var validators []ValidatorInfo

	for _, file := range files {
		if strings.Contains(file, "_test.go") || strings.Contains(file, "validatorhelper") {
			continue
		}

		info, err := analyzeValidatorFile(file)
		if err != nil {
			return nil, fmt.Errorf("failed to analyze file %s: %w", file, err)
		}

		if info != nil {
			validators = append(validators, *info)
		}
	}

	// Sort for consistent output
	sort.Slice(validators, func(i, j int) bool {
		return validators[i].MarkerName < validators[j].MarkerName
	})

	return validators, nil
}

// analyzeValidatorFile parses a Go file and extracts validator information.
func analyzeValidatorFile(filepath string) (*ValidatorInfo, error) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, filepath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file %s: %w", filepath, err)
	}

	var validatorType string

	var validateFunc string

	// Find validator struct and Validate function
	ast.Inspect(node, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.TypeSpec:
			if strings.HasSuffix(node.Name.Name, "Validator") {
				validatorType = node.Name.Name
			}
		case *ast.FuncDecl:
			if node.Name.IsExported() && strings.HasPrefix(node.Name.Name, "Validate") {
				// Check if it returns validator.Validator
				if node.Type.Results != nil && len(node.Type.Results.List) == 1 {
					validateFunc = node.Name.Name
				}
			}
		}

		return true
	})

	if validatorType == "" || validateFunc == "" {
		return nil, nil //nolint:nilnil // No valid validator found
	}

	// Extract marker name from validator type
	markerName := strings.TrimSuffix(validatorType, "Validator")
	markerName = strings.ToLower(markerName)

	return &ValidatorInfo{
		MarkerName:   markerName,
		FunctionName: validateFunc,
	}, nil
}
