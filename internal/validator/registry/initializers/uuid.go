package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// UuidInitializer implements ValidatorInitializer for the uuid validator.
type UuidInitializer struct{}

// Marker returns the marker identifier for the uuid validator.
func (u UuidInitializer) Marker() string {
	return markers.GoValidMarkerUuid
}

// Init initializes the uuid validator factory.
func (u UuidInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateUUID
}