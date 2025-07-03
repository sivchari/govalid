//go:generate govalid ./url.go

package url

type URL struct {
	// +govalid:url
	URL string `validate:"url" json:"url"`
}