package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// MaxitemsInitializer implements ValidatorInitializer for the maxitems validator.
type MaxitemsInitializer struct{}

// Marker returns the marker identifier for the maxitems validator.
func (m MaxitemsInitializer) Marker() string {
	return markers.GoValidMarkerMaxitems
}

// Init initializes the maxitems validator factory.
func (m MaxitemsInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateMaxItems
}