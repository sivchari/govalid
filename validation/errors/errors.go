// Package errors provides structures for handling validation errors.
package errors

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"
)

// ValidationError represents a single validation error.
type ValidationError struct {
	// Path is the path to the field that failed validation, e.g., "Name", "Address.City", "Users[0].Email".
	Path string
	// Type is the type of validation that failed, e.g., "required", "email", "maxlength".
	Type string
	// Value is the actual value that was validated.
	Value any
	// Reason is a human-readable message explaining why the validation failed.
	// It is the default (English) message and is used as the fallback when no
	// localized message is available for the requested language.
	Reason string
	// Messages holds localized messages keyed by language. The {Field}, {Param} and
	// {Path} placeholders are already substituted at code-generation time; the {Value}
	// placeholder, if present, is substituted at validation time from Value.
	// It is only populated when code is generated with the -i18n option.
	Messages map[language.Tag]string
}

// Localize returns a copy of the error whose Reason is replaced by the localized
// message for lang when one is available. If no localized message matches lang,
// the error is returned unchanged (falling back to the default Reason).
func (e ValidationError) Localize(lang language.Tag) ValidationError {
	msg, ok := e.lookup(lang)
	if !ok {
		return e
	}

	if strings.Contains(msg, "{Value}") {
		msg = strings.ReplaceAll(msg, "{Value}", fmt.Sprint(e.Value))
	}

	e.Reason = msg

	return e
}

// lookup finds the best localized message for lang among the registered languages,
// applying language matching (e.g. "ja-JP" matches "ja"). It returns false when no
// message is registered or none matches.
func (e ValidationError) lookup(lang language.Tag) (string, bool) {
	if len(e.Messages) == 0 {
		return "", false
	}

	if msg, ok := e.Messages[lang]; ok {
		return msg, true
	}

	tags := make([]language.Tag, 0, len(e.Messages))
	for tag := range e.Messages {
		tags = append(tags, tag)
	}

	matcher := language.NewMatcher(tags)

	_, idx, conf := matcher.Match(lang)
	if conf == language.No {
		return "", false
	}

	return e.Messages[tags[idx]], true
}

// ValidationErrors is a slice of ValidationError, representing a collection of validation errors.
type ValidationErrors []ValidationError

// Localize replaces each error's Reason with its localized message for lang where
// available, mutating and returning the slice. Errors without a matching localized
// message keep their default Reason.
func (e ValidationErrors) Localize(lang language.Tag) ValidationErrors {
	for i := range e {
		e[i] = e[i].Localize(lang)
	}

	return e
}

// Error implements the error interface for ValidationErrors.
// It returns a string representation of all validation errors, separated by newlines.
func (e ValidationErrors) Error() string {
	buff := strings.Builder{}

	for i := range e {
		buff.WriteString(e[i].Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

// Is implements error matching for ValidationErrors.
// It checks if any of the contained errors match the target.
func (e ValidationErrors) Is(target error) bool {
	for _, err := range e {
		if err.Is(target) {
			return true
		}
	}

	return false
}

// Error implements the error interface for ValidationError.
// It returns a string representation of a single validation error.
func (e ValidationError) Error() string {
	return fmt.Errorf(
		"field %s with value %v has failed validation %s because %s",
		e.Path, e.Value, e.Type, e.Reason,
	).Error()
}

// Is implements error matching for ValidationError.
// It allows errors.Is to work with ValidationError instances.
func (e ValidationError) Is(target error) bool {
	if ve, ok := target.(ValidationError); ok {
		return e.Path == ve.Path && e.Type == ve.Type && e.Reason == ve.Reason
	}

	return false
}
