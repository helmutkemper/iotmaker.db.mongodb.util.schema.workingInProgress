package iotmakerdbmongodbutilschema

import (
	"errors"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type InterfaceBson interface {
	Populate(schema map[string]interface{}) (err error)
	Verify(value interface{}) (err error)
	VerifyErros() (errorList []error)
}

type AdditionalItems interface{}

type Items interface{}

type TypeBsonGeneric struct {
	TypeBsonCommonToAllTypes
}

func (el *TypeBsonGeneric) Verify(value interface{}) (err error) {
	err = el.verifyParent(value)
	return
}

func (el *TypeBsonGeneric) Populate(schema map[string]interface{}) (err error) {
	err = el.populateGeneric(schema)
	return
}

// General
//
//   Example:
//   {
//     "type": "<JSON Type>" | ["<JSON Type>", ...],
//     "bsonType": "<BSON Type>" | ["<BSON Type>", ...],
//     "enum": [<Value 1>, <Value 2>, ...],
//     "description": "<Descriptive Text>,
//     "title": "<Short Description>"
//   }
type TypeBsonCommonToAllTypes struct {
	// The JSON type of the property the schema describes.
	// If the property’s value can be of multiple types, specify an array of JSON types.
	// Cannot be used with the bsonType field.
	//
	// The following standard JSON types are available:
	// object
	// array
	// number
	// boolean
	// string
	// null
	//
	// Note: MongoDB’s JSON Schema implementation does not support the integer JSON type.
	// Instead, use the bsonType field with int or long as the value.
	//Type TypeJson

	// The BSON type of the property the schema describes. If the property’s value can be
	// of multiple types, specify an array of BSON types.
	// Cannot be used with the type field.
	//
	// BSON types include all JSON types as well as additional types that you can reference
	// by their [string alias](https://docs.mongodb.com/manual/reference/operator/query/type/#document-type-available-types),
	// such as:
	// objectId
	// int
	// long
	// double
	// decimal
	// date
	// timestamp
	// regex
	//TypeBson TypeBson

	// Accept double converted to integer
	AcceptDoubleConvertedToInteger bool

	// An array that includes all valid values for the data that the schema describes
	Enum Enum

	// A short title or name for the data that the schema models. This field is used for
	// metadata purposes only and has no impact on schema validation.
	Title string

	// A detailed description of the data that the schema models. This field is used for
	// metadata purposes only and has no impact on schema validation.
	Description string
}

func (el *TypeBsonCommonToAllTypes) VerifyErros() (errorList []error) {
	return nil
}

func (el *TypeBsonCommonToAllTypes) round(value float64, places float64) float64 {
	var roundOn float64 = 0.5

	var round float64
	pow := math.Pow(10, places)
	digit := pow * value
	_, div := math.Modf(digit)

	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}

	return round / pow
}

func (el *TypeBsonCommonToAllTypes) verify(value interface{}) (err error) {
	err = el.verifyParent(value)
	return
}

func (el *TypeBsonCommonToAllTypes) verifyParent(value interface{}) (err error) {
	err = el.verifyEnum(value)
	return
}

func (el *TypeBsonCommonToAllTypes) verifyEnum(value interface{}) (err error) {
	if el.Enum.values == nil {
		return
	}

	err = el.Enum.Verify(value)
	return
}

func (el *TypeBsonCommonToAllTypes) Populate(schema map[string]interface{}) (err error) {
	err = el.populateGeneric(schema)
	return
}

func (el *TypeBsonCommonToAllTypes) populateGeneric(schema map[string]interface{}) (err error) {
	el.Enum, err = el.getPropertyEnum(schema)
	if err != nil {
		return
	}

	el.Title, err = el.getPropertyTitle(schema)
	if err != nil {
		return
	}

	el.Description, err = el.getPropertyDescription(schema)
	if err != nil {
		return
	}

	return
}

func (el *TypeBsonCommonToAllTypes) getPropertyAsInt64(schema map[string]interface{}, key string) (number int64, err error) {

	var value interface{}
	var found bool

	value, found = schema[key]
	if found == false {
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int64 {
		number = value.(int64)
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int32 {
		number = int64(value.(int32))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int16 {
		number = int64(value.(int16))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int8 {
		number = int64(value.(int8))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int {
		number = int64(value.(int))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint64 {
		err = errors.New("impossible to convert 'uint64' to 'int64'")
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint32 {
		number = int64(value.(uint32))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint16 {
		number = int64(value.(uint16))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint8 {
		number = int64(value.(uint8))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint {
		number = int64(value.(uint))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Float32 {
		number = int64(value.(float32))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Float64 {
		number = int64(value.(float64))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Complex64 {
		err = errors.New("impossible to convert 'complex64' to 'int64'")
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Complex128 {
		err = errors.New("impossible to convert 'complex128' to 'int64'")
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.String {
		number, err = strconv.ParseInt(value.(string), 10, 64)
		return
	}

	err = errors.New("value is not numeric")
	return
}

func (el *TypeBsonCommonToAllTypes) getPropertyAsFloat64(schema map[string]interface{}, key string) (number float64, err error) {

	var value interface{}
	var found bool

	value, found = schema[key]
	if found == false {
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int64 {
		number = float64(value.(int64))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int32 {
		number = float64(value.(int32))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int16 {
		number = float64(value.(int16))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int8 {
		number = float64(value.(int8))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int {
		number = float64(value.(int))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint64 {
		number = float64(value.(uint64))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint32 {
		number = float64(value.(uint32))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint16 {
		number = float64(value.(uint16))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint8 {
		number = float64(value.(uint8))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint {
		number = float64(value.(uint))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Float32 {
		number = float64(value.(float32))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Float64 {
		number = value.(float64)
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Complex64 {
		err = errors.New("impossible to convert 'complex64' to 'float64'")
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Complex128 {
		err = errors.New("impossible to convert 'complex128' to 'float64'")
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.String {
		number, err = strconv.ParseFloat(value.(string), 64)
		return
	}

	err = errors.New("value is not numeric")
	return
}

func (el *TypeBsonCommonToAllTypes) getPropertyAsFloat32(schema map[string]interface{}, key string) (number float32, err error) {

	var value interface{}
	var found bool

	value, found = schema[key]
	if found == false {
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int64 {
		number = float32(value.(int64))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int32 {
		number = float32(value.(int32))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int16 {
		number = float32(value.(int16))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int8 {
		number = float32(value.(int8))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Int {
		number = float32(value.(int))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint64 {
		number = float32(value.(uint64))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint32 {
		number = float32(value.(uint32))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint16 {
		number = float32(value.(uint16))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint8 {
		number = float32(value.(uint8))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Uint {
		number = float32(value.(uint))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Float32 {
		number = value.(float32)
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Float64 {
		number = float32(value.(float64))
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Float64 {
		err = errors.New("impossible to convert 'float64' to 'float32'")
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Complex64 {
		err = errors.New("impossible to convert 'complex64' to 'float32'")
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Complex128 {
		err = errors.New("impossible to convert 'complex128' to 'float32'")
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.String {
		var tmp float64
		tmp, err = strconv.ParseFloat(value.(string), 32)
		number = float32(tmp)
		return
	}

	err = errors.New("value is not numeric")
	return
}

func (el *TypeBsonCommonToAllTypes) getPropertyAsBool(schema map[string]interface{}, key string) (boolean bool, err error) {

	var value interface{}
	var found bool

	value, found = schema[key]
	if found == false {
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.Bool {
		boolean = value.(bool)
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.String {
		boolean, err = strconv.ParseBool(strings.ToLower(value.(string)))
		return
	}

	err = errors.New("value is not boolean")
	return
}

func (el *TypeBsonCommonToAllTypes) getPropertyAsMapStringInterface(schema map[string]interface{}, key string) (mapStringInterface map[string]interface{}, err error) {

	var value interface{}
	var found bool

	value, found = schema[key]
	if found == false {
		return
	}

	if reflect.ValueOf(value).Kind() != reflect.Map {
		err = errors.New("value is not map[string]interface{}")
		return
	}

	switch converted := value.(type) {
	case map[string]interface{}:
		return converted, nil
	default:
		err = errors.New("value is not map[string]interface{}")
		return
	}
}

func (el *TypeBsonCommonToAllTypes) getPropertyAsRegExp(schema map[string]interface{}, key string) (pattern *regexp.Regexp, err error) {
	var value interface{}
	var found bool

	value, found = schema[key]
	if found == false {
		return
	}

	pattern, err = regexp.Compile(value.(string))
	if err != nil {
		pattern, err = regexp.CompilePOSIX(value.(string))
	}
	return
}

func (el *TypeBsonCommonToAllTypes) getPropertyEnum(schema map[string]interface{}) (enum Enum, err error) {

	//enum.values = make([]interface{}, 0)
	var value interface{}
	var found bool

	value, found = schema["enum"]
	if found == false {
		return
	}

	if reflect.ValueOf(value).Kind() != reflect.Slice {
		err = errors.New("'enum' key must be a array")
		return
	}

	enum.values = value.([]interface{})
	return
}

func (el *TypeBsonCommonToAllTypes) getPropertyTitle(schema map[string]interface{}) (title string, err error) {
	var value interface{}
	var found bool

	value, found = schema["title"]
	if found == false {
		return
	}

	if reflect.ValueOf(value).Kind() != reflect.String {
		err = errors.New("'title' key must be a string")
		return
	}

	title = value.(string)
	return
}

func (el *TypeBsonCommonToAllTypes) getPropertyString(schema map[string]interface{}, key string) (text string, err error) {
	var value interface{}
	var found bool

	value, found = schema[key]
	if found == false {
		return
	}

	if reflect.ValueOf(value).Kind() != reflect.String {
		err = errors.New("'title' key must be a string")
		return
	}

	text = value.(string)
	return
}

func (el *TypeBsonCommonToAllTypes) getPropertyDescription(schema map[string]interface{}) (description string, err error) {
	var value interface{}
	var found bool

	value, found = schema["description"]
	if found == false {
		return
	}

	if reflect.ValueOf(value).Kind() != reflect.String {
		err = errors.New("'description' key must be a string")
		return
	}

	description = value.(string)
	return
}

func (el *TypeBsonCommonToAllTypes) parentConvertInterfaceToInt(value interface{}) (converted int, err error) {
	switch value.(type) {
	case int:
		converted = value.(int)
	case int8:
		converted = int(value.(int8))
	case int16:
		converted = int(value.(int16))
	case int32:
		converted = int(value.(int32))
	case int64:
		converted = int(value.(int64))
		if int64(converted) != value.(int64) {
			err = errors.New("wrong type")
		}
	case uint:
		converted = int(value.(uint))
	case uint8:
		converted = int(value.(uint8))
	case uint16:
		converted = int(value.(uint16))
	case uint32:
		converted = int(value.(uint32))
	case uint64:
		converted = int(value.(uint64))
		if uint64(converted) != value.(uint64) {
			err = errors.New("wrong type")
		}
	case float32:
		if el.AcceptDoubleConvertedToInteger == true {
			var tmp = int32(value.(float32))
			if value.(float32) == float32(tmp) {
				converted = int(value.(float32))
				return
			}
			err = errors.New("wrong type")

		} else {
			err = errors.New("wrong type")
		}
	case float64:
		if el.AcceptDoubleConvertedToInteger == true {
			var tmp = int32(value.(float64))
			if value.(float64) == float64(tmp) {
				converted = int(value.(float64))
				return
			}
			err = errors.New("wrong type")

		} else {
			err = errors.New("wrong type")
		}
	default:
		err = errors.New("wrong type")
	}

	return
}

func (el *TypeBsonCommonToAllTypes) parentConvertInterfaceToInt64(value interface{}) (converted int64, err error) {
	switch value.(type) {
	case int:
		converted = int64(value.(int))
	case int8:
		converted = int64(value.(int8))
	case int16:
		converted = int64(value.(int16))
	case int32:
		converted = int64(value.(int32))
	case int64:
		converted = value.(int64)
	case uint:
		converted = int64(value.(uint))
	case uint8:
		converted = int64(value.(uint8))
	case uint16:
		converted = int64(value.(uint16))
	case uint32:
		converted = int64(value.(uint32))
	case uint64:
		converted = int64(value.(uint64))
	case float32:
		if el.AcceptDoubleConvertedToInteger == true {
			var tmp = int32(value.(float32))
			if value.(float32) == float32(tmp) {
				converted = int64(value.(float32))
				return
			}
			err = errors.New("wrong type")

		} else {
			err = errors.New("wrong type")
		}
	case float64:
		if el.AcceptDoubleConvertedToInteger == true {
			var tmp = int64(value.(float64))
			if value.(float64) == float64(tmp) {
				converted = int64(value.(float64))
				return
			}
			err = errors.New("wrong type")

		} else {
			err = errors.New("wrong type")
		}
	default:
		err = errors.New("wrong type")
	}

	return
}

// fixme: verificar perda de valor
func (el *TypeBsonCommonToAllTypes) parentConvertInterfaceToFloat32(value interface{}) (converted float32, err error) {
	switch value.(type) {
	case int:
		converted = float32(value.(int))
	case int8:
		converted = float32(value.(int8))
	case int16:
		converted = float32(value.(int16))
	case int32:
		converted = float32(value.(int32))
	case int64:
		converted = float32(value.(int64))
	case uint:
		converted = float32(value.(uint))
	case uint8:
		converted = float32(value.(uint8))
	case uint16:
		converted = float32(value.(uint16))
	case uint32:
		converted = float32(value.(uint32))
	case uint64:
		converted = float32(value.(uint64))
	case float32:
		converted = value.(float32)
	case float64:
		converted = float32(value.(float64))
	default:
		err = errors.New("wrong type")
	}

	return
}

// fixme: verificar perda de valor
func (el *TypeBsonCommonToAllTypes) parentConvertInterfaceToFloat64(value interface{}) (converted float64, err error) {
	switch value.(type) {
	case int:
		converted = float64(value.(int))
	case int8:
		converted = float64(value.(int8))
	case int16:
		converted = float64(value.(int16))
	case int32:
		converted = float64(value.(int32))
	case int64:
		converted = float64(value.(int64))
	case uint:
		converted = float64(value.(uint))
	case uint8:
		converted = float64(value.(uint8))
	case uint16:
		converted = float64(value.(uint16))
	case uint32:
		converted = float64(value.(uint32))
	case uint64:
		converted = float64(value.(uint64))
	case float32:
		converted = float64(value.(float32))
	case float64:
		converted = value.(float64)
	default:
		err = errors.New("wrong type")
	}

	return
}

func (el *TypeBsonCommonToAllTypes) parentVerifyInterfaceTypeIsArray(value interface{}) (err error) {
	switch value.(type) {
	case []map[string]interface{}:
	default:
		err = errors.New("wrong type")
	}

	return
}

func (el *TypeBsonCommonToAllTypes) parentVerifyInterfaceTypeIsInt(value interface{}) (err error) {
	_, err = el.parentConvertInterfaceToInt(value)
	return
}

func (el *TypeBsonCommonToAllTypes) parentVerifyInterfaceTypeIsInt64(value interface{}) (err error) {
	_, err = el.parentConvertInterfaceToInt64(value)
	return
}

func (el *TypeBsonCommonToAllTypes) parentVerifyInterfaceTypeIsFloat32(value interface{}) (err error) {
	_, err = el.parentConvertInterfaceToFloat32(value)
	return
}

func (el *TypeBsonCommonToAllTypes) parentVerifyInterfaceTypeIsFloat64(value interface{}) (err error) {
	_, err = el.parentConvertInterfaceToFloat64(value)
	return
}
