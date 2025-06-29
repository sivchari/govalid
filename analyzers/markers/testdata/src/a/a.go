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
package a

type Markers struct {
	// +govalid:required
	Required string // want Required:`Identifier: "govalid:required", Expressions: {no expressions}`

	// +govalid:lt=10
	Min int // want Min:`Identifier: "govalid:lt", Expressions: {govalid:lt: 10}`

	// +govalid:required
	// +govalid:lt=10
	RequiredAndMin string // want RequiredAndMin:`Identifier: "govalid:required", Expressions: {no expressions}` // want RequiredAndMin:`Identifier: "govalid:lt", Expressions: {govalid:lt: 10}`
}
