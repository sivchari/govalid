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

// Package errors defines all errors used in the govalid package.
package errors

import "errors"

var (
	// ErrCouldNotGetInspector is returned when the inspector could not be retrieved.
	ErrCouldNotGetInspector = errors.New("could not get inspector")

	// ErrCouldNotCreateMarkers is returned when the markers could not be created.
	ErrCouldNotCreateMarkers = errors.New("could not create markers")
)
