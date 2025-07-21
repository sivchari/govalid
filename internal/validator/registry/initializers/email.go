package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// EmailInitializer implements ValidatorInitializer for the email validator.
type EmailInitializer struct{}

// Marker returns the marker identifier for the email validator.
func (e EmailInitializer) Marker() string {
	return markers.GoValidMarkerEmail
}

// Init initializes the email validator factory.
func (e EmailInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateEmail
}