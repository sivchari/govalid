// Package registry provides a registry system for validators.
package registry

import (
	"fmt"
	"go/ast"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
)

// ValidatorInitializer is an interface for initializing validators.
type ValidatorInitializer interface {
	// Marker returns the marker identifier for the validator (e.g., "govalid:required").
	Marker() string

	// Init initializes the validator factory.
	Init() ValidatorFactory
}

// ValidatorFactory is a function that creates a validator instance.
type ValidatorFactory func(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName string) validator.Validator

// Registry is an interface for managing a collection of validators.
type Registry interface {
	// Markers returns a slice of all registered marker identifiers.
	Markers() []string

	// Validator returns a ValidatorFactory by marker identifier.
	Validator(marker string) (ValidatorFactory, error)

	// Init initializes all validators.
	Init() error
}

// registry is an implementation of the Registry interface.
type registry struct {
	validatorInitializers []ValidatorInitializer
	initializedValidators map[string]ValidatorFactory
}

// NewRegistry creates a new Registry with the provided validator initializers.
func NewRegistry(builders ...Builder) Registry {
	r := &registry{
		initializedValidators: make(map[string]ValidatorFactory),
	}
	for _, builder := range builders {
		builder(r)
	}
	return r
}

// Builder is a function that takes a Registry and registers validators.
type Builder func(registry Registry)

// AddValidators adds a slice of ValidatorInitializers to the registry.
func AddValidators(validators ...ValidatorInitializer) Builder {
	return func(r Registry) {
		reg, ok := r.(*registry)
		if !ok {
			panic("AddValidators: registry is not of type *registry")
		}
		reg.validatorInitializers = append(reg.validatorInitializers, validators...)
	}
}

// Markers returns a slice of all registered marker identifiers.
func (r *registry) Markers() []string {
	markers := make([]string, 0, len(r.initializedValidators))
	for marker := range r.initializedValidators {
		markers = append(markers, marker)
	}
	return markers
}

// Validator returns a ValidatorFactory by marker identifier.
func (r *registry) Validator(marker string) (ValidatorFactory, error) {
	factory, exists := r.initializedValidators[marker]
	if !exists {
		return nil, fmt.Errorf("validator %s not found in registry", marker)
	}
	return factory, nil
}

// Init initializes all validators in the registry.
func (r *registry) Init() error {
	for _, initializer := range r.validatorInitializers {
		marker := initializer.Marker()
		if _, exists := r.initializedValidators[marker]; exists {
			return fmt.Errorf("duplicate validator registration for marker %s", marker)
		}
		r.initializedValidators[marker] = initializer.Init()
	}
	return nil
}

// Global registry instance
var globalRegistry Registry

// Init initializes the global registry with the provided builders.
func Init(builders ...Builder) error {
	globalRegistry = NewRegistry(builders...)
	return globalRegistry.Init()
}

// Markers returns all registered markers from the global registry.
func Markers() []string {
	if globalRegistry == nil {
		return nil
	}
	return globalRegistry.Markers()
}

// Validator retrieves a validator factory from the global registry.
func Validator(marker string) (ValidatorFactory, error) {
	if globalRegistry == nil {
		return nil, fmt.Errorf("registry not initialized")
	}
	return globalRegistry.Validator(marker)
}