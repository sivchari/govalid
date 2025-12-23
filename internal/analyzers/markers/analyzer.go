package markers

import (
	"fmt"
	"go/ast"
	"os"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	govaliderrors "github.com/sivchari/govalid/internal/errors"
)

const (
	// Name is the name of the markers analyzer.
	Name = "markers"
	// Doc is the documentation for the markers analyzer.
	Doc = "markers is a helper for generating govalid validation"
)

// Analyzer is the main entry point for the markers analyzer.
// This variable must be initialized by registry package.
var Analyzer *analysis.Analyzer

// analyzer implements the analysis.Analyzer interface for the markers analyzer.
type analyzer struct{}

// newAnalyzer creates a new instance of the markers analyzer.
func newAnalyzer() *analysis.Analyzer {
	a := &analyzer{}
	analyzer := &analysis.Analyzer{
		Name:       Name,
		Doc:        Doc,
		Run:        a.run,
		Requires:   []*analysis.Analyzer{inspect.Analyzer},
		ResultType: reflect.TypeOf(newMarkers()),
		FactTypes: []analysis.Fact{
			(*MarkerFact)(nil),
		},
	}

	return analyzer
}

// run is the main function that runs the markers analyzer.
func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, govaliderrors.ErrCouldNotGetInspector
	}

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}

	results, ok := newMarkers().(*markers)
	if !ok {
		return nil, govaliderrors.ErrCouldNotCreateMarkers
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.GenDecl:
			if n == nil {
				return
			}

			if n.Doc != nil && len(n.Doc.List) > 0 {
				collectTypeMarkers(pass, n, results)
			}

			for _, spec := range n.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				st, ok := ts.Type.(*ast.StructType)
				if !ok {
					return
				}

				collectStructMarkers(pass, st, results)
			}
		default:
		}
	})

	return results, nil
}

// collectTypeMarkers collects markers from a GenDecl node and adds them to the results.
func collectTypeMarkers(pass *analysis.Pass, genDecl *ast.GenDecl, results *markers) {
	if genDecl.Doc == nil || len(genDecl.Doc.List) == 0 {
		return
	}

	for _, doc := range genDecl.Doc.List {
		pos := pass.Fset.Position(doc.Pos())
		markerContent, ok := parseMarkerComment(doc.Text, pos.Filename, pos.Line)

		if !ok {
			continue
		}

		identifier, expressions := extractMarker(markerContent)
		marker := Marker{
			Identifier:  identifier,
			Expressions: expressions,
		}

		for _, spec := range genDecl.Specs {
			if ts, ok := spec.(*ast.TypeSpec); ok {
				results.insertTypeMarker(ts, marker)

				if obj, ok := pass.TypesInfo.Defs[ts.Name]; ok {
					pass.ExportObjectFact(obj, &MarkerFact{
						Identifier:  identifier,
						Expressions: expressions,
					})
				}
			}
		}
	}
}

// collectStructMarkers collects markers from a TypeSpec node and adds them to the results.
func collectStructMarkers(pass *analysis.Pass, s *ast.StructType, results *markers) {
	if s == nil || s.Fields == nil || len(s.Fields.List) == 0 {
		return
	}

	for _, field := range s.Fields.List {
		fieldMarkers(pass, field, results)

		structType, ok := field.Type.(*ast.StructType)
		if !ok {
			continue
		}

		collectStructMarkers(pass, structType, results)
	}
}

// fieldMarkers extracts markers from a struct field and adds them to the results.
func fieldMarkers(pass *analysis.Pass, field *ast.Field, results *markers) {
	if field == nil || field.Doc == nil || len(field.Doc.List) == 0 {
		return
	}

	for _, doc := range field.Doc.List {
		pos := pass.Fset.Position(doc.Pos())
		markerContent, ok := parseMarkerComment(doc.Text, pos.Filename, pos.Line)

		if !ok {
			continue
		}

		identifier, expressions := extractMarker(markerContent)
		marker := Marker{
			Identifier:  identifier,
			Expressions: expressions,
		}
		results.insertFieldMarker(field, marker)

		if obj, ok := pass.TypesInfo.Defs[field.Names[0]]; ok {
			pass.ExportObjectFact(obj, &MarkerFact{
				Identifier:  identifier,
				Expressions: expressions,
			})
		}
	}
}

// parseMarkerComment parses a comment and extracts the marker content.
// It supports both old format "// +govalid:" and new format "//govalid:".
// Returns the marker content (e.g., "govalid:required") and true if valid, or empty string and false if not a marker.
func parseMarkerComment(text, filename string, line int) (string, bool) {
	// New format: //govalid:xxx (recommended, go doc directive style)
	if strings.HasPrefix(text, "//govalid:") {
		return strings.TrimPrefix(text, "//"), true
	}
	// Old format: // +govalid:xxx (deprecated, kept for backward compatibility)
	if strings.HasPrefix(text, "// +govalid:") {
		fmt.Fprintf(os.Stderr, "\033[33mWARNING\033[0m: %s:%d: "+`deprecated marker format "// +govalid:"; use "//govalid:" instead (run 'govalid migrate' to fix)`+"\n", filename, line)

		return strings.TrimPrefix(text, "// +"), true
	}

	return "", false
}

// extractMarker extracts the identifier and expressions from a marker content string.
// It returns the identifier and a map of expressions if applicable.
// If the content does not contain an identifier or expressions, it returns an empty string and nil.
func extractMarker(content string) (string, map[string]string) {
	if strings.Count(content, "=") == 0 {
		return content, nil
	}

	// Split on the first = only, allowing expressions to contain = characters
	splits := strings.SplitN(content, "=", 2)
	if len(splits) != 2 {
		return "", nil
	}

	expressions := map[string]string{}
	expressions[splits[0]] = splits[1]

	return splits[0], expressions
}
