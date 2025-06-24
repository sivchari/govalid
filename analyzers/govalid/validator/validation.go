package validator

type Validator interface {
	Validate() string
	FieldName() string
	Err() string
}

var generatorMemory = map[string]bool{}
