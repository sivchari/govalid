package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// MinItemsInitializer implements ValidatorInitializer for the minitems validator.
type MinItemsInitializer struct{}

// Marker returns the marker identifier for the minitems validator.
func (m MinItemsInitializer) Marker() string {
	return markers.GoValidMarkerMinItems
}

// Init initializes the minitems validator factory.
func (m MinItemsInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateMinItems
}