//go:generate ./alphanum.go
package alphanum

// Alphanum is a struct for testing alphanum validation
type Alphanum struct {
	//govalid:alphanum
	ProductCode string `json:"product_code"`

	//govalid:alphanum
	SerialNumber string `json:"serial_number"`

	//govalid:alphanum
	Username string `json:"username"`

	// Non-string field: the alphanum rule does not apply and is skipped.
	//govalid:alphanum
	Revision int `json:"revision"`
}
