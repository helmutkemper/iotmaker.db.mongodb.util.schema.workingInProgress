package iotmakerdbmongodbutilschema

import (
	"errors"
	"reflect"
)

func (el *Element) getPropertyTitle(schema map[string]interface{}) (value string, err error) {
	var title interface{}
	var found bool

	title, found = schema["title"]
	if found == false {
		return
	}

	if reflect.ValueOf(title).Kind() != reflect.String {
		err = errors.New("'title' key must be a string")
		return
	}

	value = title.(string)
	return
}
