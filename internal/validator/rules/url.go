// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"
	"github.com/sivchari/govalid/internal/validator"
)

type urlValidator struct {
	pass  *codegen.Pass
	field *ast.Field
}

var _ validator.Validator = (*urlValidator)(nil)

const urlKey = "%s-url"

func (u *urlValidator) Validate() string {
	fieldName := u.FieldName()
	// Generate inline manual URL validation for maximum performance
	return fmt.Sprintf("!isValidURL(t.%s)", fieldName)
}

func (u *urlValidator) FieldName() string {
	return u.field.Names[0].Name
}

func (u *urlValidator) Err() string {
	fieldName := u.FieldName()
	
	var result strings.Builder
	
	// Generate isValidURL function only once
	if !validator.GeneratorMemory["url-function-generated"] {
		validator.GeneratorMemory["url-function-generated"] = true
		result.WriteString(`
	// isValidURL validates URL format manually for maximum performance
	// Validates HTTP/HTTPS URLs with basic structure checking
	isValidURL = func(url string) bool {
		// Check minimum length
		if len(url) < 10 { // "http://a.b" = 10 chars minimum
			return false
		}
		
		// Check protocol
		var rest string
		if len(url) >= 7 && (url[0] == 'h' || url[0] == 'H') {
			lower := make([]byte, 7)
			for i := 0; i < 7; i++ {
				if url[i] >= 'A' && url[i] <= 'Z' {
					lower[i] = url[i] + 32 // convert to lowercase
				} else {
					lower[i] = url[i]
				}
			}
			if string(lower) == "http://" {
				rest = url[7:]
			} else if len(url) >= 8 {
				lower = make([]byte, 8)
				for i := 0; i < 8; i++ {
					if url[i] >= 'A' && url[i] <= 'Z' {
						lower[i] = url[i] + 32
					} else {
						lower[i] = url[i]
					}
				}
				if string(lower) == "https://" {
					rest = url[8:]
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
		
		if len(rest) == 0 {
			return false
		}
		
		// Find path, query, or fragment separator
		pathIndex := -1
		for i, c := range rest {
			if c == '/' || c == '?' || c == '#' {
				pathIndex = i
				break
			}
		}
		
		var hostPort, pathPart string
		if pathIndex == -1 {
			hostPort = rest
			pathPart = ""
		} else {
			hostPort = rest[:pathIndex]
			pathPart = rest[pathIndex:]
		}
		
		// Validate host:port part
		if len(hostPort) == 0 {
			return false
		}
		
		// Check for port (find last colon)
		colonIndex := -1
		for i := len(hostPort) - 1; i >= 0; i-- {
			if hostPort[i] == ':' {
				colonIndex = i
				break
			}
		}
		
		var host, port string
		if colonIndex == -1 {
			host = hostPort
		} else {
			host = hostPort[:colonIndex]
			port = hostPort[colonIndex+1:]
			
			// Validate port number
			if len(port) == 0 || len(port) > 5 {
				return false
			}
			for _, c := range port {
				if c < '0' || c > '9' {
					return false
				}
			}
		}
		
		// Validate hostname
		if len(host) == 0 || len(host) > 253 {
			return false
		}
		
		// Must contain at least one dot
		hasDot := false
		for _, c := range host {
			if c == '.' {
				hasDot = true
				break
			}
		}
		if !hasDot {
			return false
		}
		
		// Split by dots and validate each label
		labels := make([]string, 0, 4)
		start := 0
		for i := 0; i <= len(host); i++ {
			if i == len(host) || host[i] == '.' {
				if i > start {
					label := host[start:i]
					if len(label) == 0 || len(label) > 63 {
						return false
					}
					
					// Can't start or end with hyphen
					if label[0] == '-' || label[len(label)-1] == '-' {
						return false
					}
					
					// Check characters (alphanumeric + hyphen)
					for _, c := range label {
						if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || 
							 (c >= '0' && c <= '9') || c == '-') {
							return false
						}
					}
					
					labels = append(labels, label)
				}
				start = i + 1
			}
		}
		
		if len(labels) < 2 {
			return false
		}
		
		// Basic path/query/fragment validation (if present)
		if pathPart != "" {
			// Must start with /, ?, or #
			if pathPart[0] != '/' && pathPart[0] != '?' && pathPart[0] != '#' {
				return false
			}
			
			// No control chars or spaces
			for _, c := range pathPart {
				if c < 32 || c == ' ' {
					return false
				}
			}
		}
		
		return true
	}`)
	}
	
	if validator.GeneratorMemory[fmt.Sprintf(urlKey, fieldName)] {
		return result.String()
	}

	validator.GeneratorMemory[fmt.Sprintf(urlKey, fieldName)] = true

	result.WriteString(fmt.Sprintf(strings.ReplaceAll(`
	// Err@URLValidation is the error returned when the field is not a valid URL.
	Err@URLValidation = errors.New("field @ must be a valid URL")`, "@", fieldName)))
	
	return result.String()
}

func (u *urlValidator) ErrVariable() string {
	return strings.ReplaceAll("Err@URLValidation", "@", u.FieldName())
}

func (u *urlValidator) Imports() []string {
	return []string{} // No imports needed for manual validation
}

// ValidateURL creates a new urlValidator for string types.
func ValidateURL(pass *codegen.Pass, field *ast.Field, expressions map[string]string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &urlValidator{
		pass:  pass,
		field: field,
	}
}