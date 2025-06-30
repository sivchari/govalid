// Package errors defines all errors used in the govalid package.
package errors

import "errors"

var (
	// ErrCouldNotGetInspector is returned when the inspector could not be retrieved.
	ErrCouldNotGetInspector = errors.New("could not get inspector")

	// ErrCouldNotMarkersInspector is returned when the markers inspector could not be retrieved.
	ErrCouldNotMarkersInspector = errors.New("could not get markers inspector")

	// ErrCouldNotCreateMarkers is returned when the markers could not be created.
	ErrCouldNotCreateMarkers = errors.New("could not create markers")
)
