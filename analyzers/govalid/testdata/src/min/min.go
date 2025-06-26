/*
Copyright {{license-year}} sivchari.

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
package min

//go:generate govalid ./min.go

type Min struct {
	// +govalid:min=1
	Int int

	// +govalid:min=1
	Int8 int8

	// +govalid:min=1
	Int16 int16

	// +govalid:min=1
	Int32 int32

	// +govalid:min=1
	Int64 int64

	// +govalid:min=1
	Float32 float32

	// +govalid:min=1
	Float64 float64

	// +govalid:min=1
	Uint uint

	// +govalid:min=1
	Uint8 uint8

	// +govalid:min=1
	Uint16 uint16

	// +govalid:min=1
	Uint32 uint32

	// +govalid:min=1
	Uint64 uint64

	// +govalid:min=1
	Uintptr uintptr

	// +govalid:min=1
	Complex64 complex64

	// +govalid:min=1
	Complex128 complex128

	// +govalid:min=1
	String string

	Struct struct {
		// +govalid:min=1
		Int int
	}
}
