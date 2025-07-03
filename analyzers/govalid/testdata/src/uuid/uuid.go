//go:generate govalid ./uuid.go

package uuid

type UUID struct {
	// +govalid:uuid
	UUID string `validate:"uuid" json:"uuid"`
}