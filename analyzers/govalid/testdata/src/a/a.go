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
package a

//go:generate govalid ./a.go

type User struct {
	// +govalid:required
	String string

	// +govalid:required
	Int int

	// +govalid:required
	Array [1]string

	// +govalid:required
	Slice []string

	// +govalid:required
	Map map[string]string

	// +govalid:required
	Interface any

	// +govalid:required
	Pointer *string

	// +govalid:required
	Struct struct{}

	// +govalid:required
	Channel chan string

	// +govalid:required
	Func func(string) string
}
