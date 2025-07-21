package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// MaxLengthInitializer implements ValidatorInitializer for the maxlength validator.
type MaxLengthInitializer struct{}

// Marker returns the marker identifier for the maxlength validator.
func (m MaxLengthInitializer) Marker() string {
	return markers.GoValidMarkerMaxLength
}

// Init initializes the maxlength validator factory.
func (m MaxLengthInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateMaxLength
}