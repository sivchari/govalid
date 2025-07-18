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
	// Values should be comma-separated (e.g., admin,user,guest).
	GoValidMarkerEnum = "govalid:enum"

	// GoValidMarkerEmail is the marker for checking if a string is a valid email address in govalid.
	// This marker is only available for string types and follows HTML5 email validation specification.
	GoValidMarkerEmail = "govalid:email"

	// GoValidMarkerUUID is the marker for checking if a string is a valid UUID in govalid.
	// This marker is only available for string types and follows RFC 4122 UUID specification.
	GoValidMarkerUUID = "govalid:uuid"

	// GoValidMarkerURL is the marker for checking if a string is a valid URL in govalid.
	// This marker is only available for string types and follows RFC 3986 URL specification.
	GoValidMarkerURL = "govalid:url"

	// GoValidMarkerCEL is the marker for CEL (Common Expression Language) validation in govalid.
	// This marker supports complex validation expressions with access to 'value' (current field) and 'this' (struct instance).
	// Available for all types and enables cross-field validation and complex business rules.
	GoValidMarkerCEL = "govalid:cel"

	// GoValidMarkerLength is the marker for checking if a collection's length is equal to a specified length in govalid.
	// This marker is only available for string types and enhances the maxlength and minlength markers.
	GoValidMarkerLength = "govalid:length"
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
	GoValidMarkerEmail:     {},
	GoValidMarkerUUID:      {},
	GoValidMarkerURL:       {},
	GoValidMarkerCEL:       {},
	GoValidMarkerLength:    {},
}
