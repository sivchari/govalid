package a

type Markers struct {
	//govalid:required
	Required string // want Required:`Identifier: "govalid:required", Expressions: {no expressions}`

	//govalid:lt=10
	Min int // want Min:`Identifier: "govalid:lt", Expressions: {govalid:lt: 10}`

	//govalid:required
	//govalid:lt=10
	RequiredAndMin string // want RequiredAndMin:`Identifier: "govalid:required", Expressions: {no expressions}` // want RequiredAndMin:`Identifier: "govalid:lt", Expressions: {govalid:lt: 10}`
}

//govalid:required
type TypeLevelMarkers struct { // want TypeLevelMarkers:`Identifier: "govalid:required", Expressions: {no expressions}`
	String string

	//govalid:required
	RequiredString int // want RequiredString:`Identifier: "govalid:required", Expressions: {no expressions}`

	//govalid:email
	Email string // want Email:`Identifier: "govalid:email", Expressions: {no expressions}`
}
