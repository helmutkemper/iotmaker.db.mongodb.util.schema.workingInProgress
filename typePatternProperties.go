package iotmakerdbmongodbutilschema

import (
	"errors"
	"regexp"
)

// PatternProperties (English):
//
// PatternProperties (Português):
type PatternProperties struct {
	regexp     *regexp.Regexp
	properties []MongoDBJsonSchema
}

// AppendProperty (English): Add a new property
//
// AppendProperty (Português): Adiciona uma nova propriedade
func (el *PatternProperties) AppendProperty(property MongoDBJsonSchema) {
	if len(el.properties) == 0 {
		el.properties = make([]MongoDBJsonSchema, 0)
	}

	el.properties = append(el.properties, property)
}

// SetRegexp (English): set a regular expression for property
//
// SetRegexp (Português): define uma expressão regular para a propriedade
func (el *PatternProperties) SetRegexp(value string) (err error) {
	el.regexp, err = regexp.Compile(value)
	return
}

// SetRegexpPOSIX (English): set a posix regular expression for property
//
// SetRegexpPOSIX (Português): define uma expressão regular posix para a propriedade
func (el *PatternProperties) SetRegexpPOSIX(value string) (err error) {
	el.regexp, err = regexp.CompilePOSIX(value)
	return
}

// GetMatch (English): Returns a list of properties for the given key
//
// GetMatch (Português): Retorna a de propriedades para uma determinada chave
func (el *PatternProperties) GetMatch(value string) (propertiesList []MongoDBJsonSchema, err error) {
	propertiesList = make([]MongoDBJsonSchema, 0)

	if el.regexp.MatchString(value) == true {
		if len(el.properties) == 0 {
			el.properties = make([]MongoDBJsonSchema, 0)
		}
		propertiesList = el.properties
		return
	}

	err = errors.New("the key name does not match the regular expression")
	return
}
