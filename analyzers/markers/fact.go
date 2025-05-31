/*
Copyright 2025 sivchari.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package markers

import (
	"fmt"
	"strings"
)

type MarkerFact struct {
	Identifier  string
	Expressions map[string]string
}

func (mf *MarkerFact) AFact() {}

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
