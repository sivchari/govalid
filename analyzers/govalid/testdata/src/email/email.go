//go:generate govalid ./email.go

package email

type Email struct {
	// +govalid:email
	Email string `validate:"email" json:"email"`
}