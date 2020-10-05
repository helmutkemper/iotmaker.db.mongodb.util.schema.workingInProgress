package iotmakerdbmongodbutilschema

import (
	"errors"
	"regexp"
)

// The string schema type configures the value of string fields.
// For more information, see the official JSON Schema string guide.
// https://json-schema.org/understanding-json-schema/reference/string.html
//
//   Example:
//   {
//     "bsonType": "string",
//     "maxLength": <integer>,
//     "minLength": <integer>,
//     "pattern": "<Regular Expression>"
//   }
type TypeBsonString struct {
	TypeBsonCommonToAllTypes

	// The maximum number of characters in the string.
	MaxLength int64

	// The minimum number of characters in the string.
	MinLength int64

	// A regular expression string that must match the string value.
	Pattern *regexp.Regexp
}

func (el *TypeBsonString) Verify(value interface{}) (err error) {
	err = el.verifyParent(value)
	if err != nil {
		return
	}

	err = el.VerifyType(value)
	if err != nil {
		return
	}

	err = el.VerifyMaxLength(value)
	if err != nil {
		return
	}

	err = el.VerifyMinLength(value)
	if err != nil {
		return
	}

	err = el.VerifyPattern(value)
	if err != nil {
		return
	}

	return
}

func (el *TypeBsonString) VerifyMaxLength(value interface{}) (err error) {
	if value == nil || el.MaxLength == 0 {
		return
	}

	if len(value.(string)) > int(el.MaxLength) {
		err = errors.New("maximum string size exceeded")
	}

	return
}

func (el *TypeBsonString) VerifyMinLength(value interface{}) (err error) {
	if value == nil || el.MinLength == 0 {
		return
	}

	if len(value.(string)) < int(el.MinLength) {
		err = errors.New("minimum string length expected")
	}

	return
}

func (el *TypeBsonString) VerifyPattern(value interface{}) (err error) {
	if value == nil || el.Pattern == nil {
		return
	}

	if el.Pattern.MatchString(value.(string)) == false {
		err = errors.New("the string does not match the regular expression")
	}

	return
}

func (el *TypeBsonString) VerifyType(value interface{}) (err error) {
	switch value.(type) {
	case string:
	case nil:
		if el.Enum.values != nil {
			err = el.Enum.Verify(value)
		}
	default:
		err = errors.New("wrong type")
	}

	return
}

func (el *TypeBsonString) Populate(schema map[string]interface{}) (err error) {
	err = el.populateGeneric(schema)
	if err != nil {
		return
	}

	el.MaxLength, err = el.getPropertyMaxLength(schema)
	if err != nil {
		return
	}

	el.MinLength, err = el.getPropertyMinLength(schema)
	if err != nil {
		return
	}

	el.Pattern, err = el.getPropertyPattern(schema)
	return
}

func (el *TypeBsonString) getPropertyMaxLength(schema map[string]interface{}) (maxLength int64, err error) {
	var found bool

	_, found = schema["maxLength"]
	if found == false {
		return
	}

	maxLength, err = el.getPropertyAsInt64(schema, "maxLength")
	return
}

func (el *TypeBsonString) getPropertyMinLength(schema map[string]interface{}) (minLength int64, err error) {
	var found bool

	_, found = schema["minLength"]
	if found == false {
		return
	}

	minLength, err = el.getPropertyAsInt64(schema, "minLength")
	return
}

func (el *TypeBsonString) getPropertyPattern(schema map[string]interface{}) (pattern *regexp.Regexp, err error) {
	var found bool

	_, found = schema["pattern"]
	if found == false {
		return
	}

	pattern, err = el.getPropertyAsRegExp(schema, "pattern")
	return
}
