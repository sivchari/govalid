package errors

type ValidationError struct {
	Path   string // "Name", "Address.City", "Users[0].Email"
	Type   string // "required", "email", "maxlength"
	Value  any    // The actual value being validated
	Reason string // "field is required"
}

type ValidationErrors []ValidationError

// Implement error interface
func (e ValidationErrors) Error() string {
	return ""
}
