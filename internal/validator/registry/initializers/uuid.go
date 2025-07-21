package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// UUIDInitializer implements ValidatorInitializer for the uuid validator.
type UUIDInitializer struct{}

// Marker returns the marker identifier for the uuid validator.
func (u UUIDInitializer) Marker() string {
	return markers.GoValidMarkerUUID
}

// Init initializes the uuid validator factory.
func (u UUIDInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateUUID
}