package markers

import (
	"go/ast"
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
		(*ast.TypeSpec)(nil),
	}

	results, ok := newMarkers().(*markers)
	if !ok {
		return nil, govaliderrors.ErrCouldNotCreateMarkers
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.TypeSpec:
			if n == nil {
				return
			}

			st, ok := n.Type.(*ast.StructType)
			if !ok {
				return
			}

			collectStructMarkers(pass, st, results)
		default:
		}
	})

	return results, nil
}

// collectTypeSpecMarkers collects markers from a TypeSpec node and adds them to the results.
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
		if !strings.HasPrefix(doc.Text, "// +") {
			continue
		}

		markerContent := strings.TrimPrefix(doc.Text, "// +")

		identifier, expressions := extractMarker(markerContent)
		mset := newMarkerSet()
		mset[identifier] = Marker{
			Identifier:  identifier,
			Expressions: expressions,
		}
		results.insertFieldMarkers(field, mset)

		if obj, ok := pass.TypesInfo.Defs[field.Names[0]]; ok {
			pass.ExportObjectFact(obj, &MarkerFact{
				Identifier:  identifier,
				Expressions: expressions,
			})
		}
	}
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
