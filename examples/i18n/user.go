// Package i18n demonstrates generating validators with localized error messages
// via `govalid -i18n=i18n.yaml`.
package i18n

//go:generate go run github.com/sivchari/govalid/cmd/govalid -i18n=i18n.yaml ./user.go

// User is a sample struct validated with localized messages.
type User struct {
	//govalid:required
	Name string

	//govalid:gt=18
	Age int

	//govalid:maxlength=20
	Nickname string
}
