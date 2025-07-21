package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// MaxItemsInitializer implements ValidatorInitializer for the maxitems validator.
type MaxItemsInitializer struct{}

// Marker returns the marker identifier for the maxitems validator.
func (m MaxItemsInitializer) Marker() string {
	return markers.GoValidMarkerMaxItems
}

// Init initializes the maxitems validator factory.
func (m MaxItemsInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateMaxItems
}