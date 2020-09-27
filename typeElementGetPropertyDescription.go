package iotmakerdbmongodbutilschema

import (
	"errors"
	"reflect"
)

func (el *Element) getPropertyDescription(schema map[string]interface{}) (value string, err error) {
	var description interface{}
	var found bool

	description, found = schema["description"]
	if found == false {
		return
	}

	if reflect.ValueOf(description).Kind() != reflect.String {
		err = errors.New("'description' key must be a string")
		return
	}

	value = description.(string)
	return
}
