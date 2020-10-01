package iotmakerdbmongodbutilschema

import (
	"errors"
	"reflect"
)

// slicerAndAssemblerForBsonType (English):
//
// slicerAndAssemblerForBsonType (Português):
func (el *Element) slicerAndAssemblerForBsonType(schema map[string]interface{}) (err error) {

	var value = schema["bsonType"]

	switch reflect.ValueOf(value).Kind() {
	case reflect.String:
		if reflect.ValueOf(value).Kind() != reflect.String {
			err = errors.New("MongoDB schema, key '$jsonSchema' must be a key 'bsonType' and 'bsonType' key must be a string with value 'object'")
			return
		}

		err = el.typeStringToTypeObjectPopulated(&el.Properties, "main", value.(string), schema)
		if err != nil {
			return
		}

		//err = el.populateRequired("main", value.(string), schema)
		//if err != nil {
		//	return
		//}

		err = el.populateDependencies(schema)
		if err != nil {
			return
		}

	default:
		err = errors.New("MongoDB schema, key '$jsonSchema' must be a key 'bsonType' and 'bsonType' key must be a string with value 'object'")
		return
	}

	return
}
