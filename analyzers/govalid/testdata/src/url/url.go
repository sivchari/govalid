//go:generate govalid ./url.go

package url

type URL struct {
	// +govalid:url
	WebsiteURL string `validate:"url" json:"website_url"`

	// +govalid:url
	HomepageURL string `validate:"url" json:"homepage_url"`

	// +govalid:url
	ApiURL string `validate:"url" json:"api_url"`

	// +govalid:url
	ProfileURL string `validate:"url" json:"profile_url"`

	// +govalid:url
	DownloadURL string `validate:"url" json:"download_url"`
}
