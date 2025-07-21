package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// MaxlengthInitializer implements ValidatorInitializer for the maxlength validator.
type MaxlengthInitializer struct{}

// Marker returns the marker identifier for the maxlength validator.
func (m MaxlengthInitializer) Marker() string {
	return markers.GoValidMarkerMaxlength
}

// Init initializes the maxlength validator factory.
func (m MaxlengthInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateMaxLength
}