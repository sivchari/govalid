package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// MinLengthInitializer implements ValidatorInitializer for the minlength validator.
type MinLengthInitializer struct{}

// Marker returns the marker identifier for the minlength validator.
func (m MinLengthInitializer) Marker() string {
	return markers.GoValidMarkerMinLength
}

// Init initializes the minlength validator factory.
func (m MinLengthInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateMinLength
}