/*
Copyright 2025 sivchari.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package markers defines all markers used in the govalid package.
package markers

var (
	// GoValidMarkerRequired is the marker for required fields in govalid.
	GoValidMarkerRequired = "govalid:required"

	// GoValidMarkerMin is the marker for minimum value in govalid.
	// This marker is only available for numeric types.
	GoValidMarkerMin = "govalid:min"

	// GoValidMarkerMax is the marker for maximum value in govalid.
	// This marker is only available for numeric types.
	GoValidMarkerMax = "govalid:max"
)

// GoValidMarkers is a map of valid govalid markers.
var GoValidMarkers = map[string]struct{}{
	GoValidMarkerRequired: {},
	GoValidMarkerMin:      {},
	GoValidMarkerMax:      {},
}
