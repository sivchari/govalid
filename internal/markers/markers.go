// Package markers defines all markers used in the govalid package.
package markers

var (
	// GoValidMarkerRequired is the marker for required fields in govalid.
	GoValidMarkerRequired = "govalid:required"

	// GoValidMarkerLT is the marker fro checking if a value is less than a specified maximum in govalid.
	// This marker is only available for numeric types.
	GoValidMarkerLT = "govalid:lt"

	// GoValidMarkerGT is the marker for checking if a value is greater than a specified minimum in govalid.
	// This marker is only available for numeric types.
	GoValidMarkerGT = "govalid:gt"
)

// GoValidMarkers is a map of valid govalid markers.
var GoValidMarkers = map[string]struct{}{
	GoValidMarkerRequired: {},
	GoValidMarkerLT:       {},
	GoValidMarkerGT:       {},
}
