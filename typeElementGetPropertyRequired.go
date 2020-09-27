package iotmakerdbmongodbutilschema

import (
	"errors"
	"reflect"
)

func (el *Element) getPropertyRequired(schema map[string]interface{}) (value []string, err error) {

	value = make([]string, 0)

	var required interface{}
	var found bool

	required, found = schema["required"]
	if found == false {
		return
	}

	if reflect.ValueOf(required).Kind() != reflect.Slice {
		err = errors.New("'title' key must be a string")
		return
	}

	for _, v := range required.([]interface{}) {
		if reflect.ValueOf(v).Kind() != reflect.String {
			err = errors.New("the required values must be string")
			return
		}
		value = append(value, v.(string))
	}

	return
}
