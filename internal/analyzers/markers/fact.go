package markers

import (
	"fmt"
	"strings"
)

// MarkerFact is a struct that represents a marker fact.
type MarkerFact struct {
	Identifier  string
	Expressions map[string]string
}

// AFact is a method that satisfies the Fact interface.
func (mf *MarkerFact) AFact() {}

// String returns a string representation of the MarkerFact.
func (mf *MarkerFact) String() string {
	if mf == nil {
		return "<nil>"
	}

	var expressionsString string
	if len(mf.Expressions) == 0 {
		expressionsString = "no expressions"
	} else {
		elms := make([]string, 0, len(mf.Expressions))
		for key, value := range mf.Expressions {
			elms = append(elms, fmt.Sprintf("%s: %s", key, value))
		}

		expressionsString = strings.Join(elms, ", ")
	}

	quotedIdentifier := fmt.Sprintf("%q", mf.Identifier)

	return fmt.Sprintf("Identifier: %s, Expressions: {%s}", quotedIdentifier, expressionsString)
}
