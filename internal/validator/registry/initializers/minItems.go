package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// MinitemsInitializer implements ValidatorInitializer for the minitems validator.
type MinitemsInitializer struct{}

// Marker returns the marker identifier for the minitems validator.
func (m MinitemsInitializer) Marker() string {
	return markers.GoValidMarkerMinitems
}

// Init initializes the minitems validator factory.
func (m MinitemsInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateMinItems
}