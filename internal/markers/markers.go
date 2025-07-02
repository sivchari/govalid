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

	// GoValidMarkerMaxLength is the marker for checking if a string's length is less than or equal to a specified maximum in govalid.
	// This marker is only available for string types.
	GoValidMarkerMaxLength = "govalid:maxlength"

	// GoValidMarkerMinLength is the marker for checking if a string's length is greater than or equal to a specified minimum in govalid.
	// This marker is only available for string types.
	GoValidMarkerMinLength = "govalid:minlength"

	// GoValidMarkerMaxItems is the marker for checking if a collection's length is less than or equal to a specified maximum in govalid.
	// This marker is available for slice, array, map, and channel types.
	GoValidMarkerMaxItems = "govalid:maxitems"

	// GoValidMarkerMinItems is the marker for checking if a collection's length is greater than or equal to a specified minimum in govalid.
	// This marker is available for slice, array, map, and channel types.
	GoValidMarkerMinItems = "govalid:minitems"

	// GoValidMarkerGTE is the marker for checking if a value is greater than or equal to a specified minimum in govalid.
	// This marker is only available for numeric types.
	GoValidMarkerGTE = "govalid:gte"

	// GoValidMarkerLTE is the marker for checking if a value is less than or equal to a specified maximum in govalid.
	// This marker is only available for numeric types.
	GoValidMarkerLTE = "govalid:lte"

	// GoValidMarkerEnum is the marker for checking if a field value is within a specified set of allowed values in govalid.
	// This marker is available for string, numeric, and custom types with comparable values.
	GoValidMarkerEnum = "govalid:enum"
)

// GoValidMarkers is a map of valid govalid markers.
var GoValidMarkers = map[string]struct{}{
	GoValidMarkerRequired:  {},
	GoValidMarkerLT:        {},
	GoValidMarkerGT:        {},
	GoValidMarkerMaxLength: {},
	GoValidMarkerMaxItems:  {},
	GoValidMarkerMinItems:  {},
	GoValidMarkerMinLength: {},
	GoValidMarkerGTE:       {},
	GoValidMarkerLTE:       {},
	GoValidMarkerEnum:      {},
}
