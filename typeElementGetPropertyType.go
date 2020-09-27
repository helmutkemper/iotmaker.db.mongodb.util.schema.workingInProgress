package iotmakerdbmongodbutilschema

import (
	"errors"
	"reflect"
)

func (el *Element) getPropertyBsonType(schema map[string]interface{}) (value []string, err error) {

	value = make([]string, 0)

	var bsonType interface{}
	var found bool

	bsonType, found = schema["bsonType"]
	if found == false {
		return
	}

	if reflect.ValueOf(bsonType).Kind() == reflect.Slice {
		for _, v := range bsonType.([]interface{}) {
			if reflect.ValueOf(v).Kind() != reflect.String {
				err = errors.New("the 'bsonType' values must be a string")
				return
			}

			value = append(value, v.(string))
		}
		return
	}

	if reflect.ValueOf(bsonType).Kind() == reflect.String {
		value = append(value, bsonType.(string))
		return
	}

	err = errors.New("the 'bsonType' a string or a array of string")
	return
}
