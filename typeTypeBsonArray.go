package iotmakerdbmongodbutilschema

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// Arrays
// The array schema type configures the content of array fields.
// For more information, see the official JSON Schema array guide.
// https://json-schema.org/understanding-json-schema/reference/array.html
// https://www.mongodb.com/blog/post/json-schema-validation--locking-down-your-model-the-smart-way
//
//   Example:
//   {
//     "bsonType": "array",
//     "items": <Schema Document> | [<Schema Document>, ...],
//     "additionalItems": <boolean> | <Schema Document>,
//     "maxItems": <integer>,
//     "minItems": <integer>,
//     "uniqueItems": <boolean>
//   }
type TypeBsonArray struct {
	TypeBsonCommonToAllTypes

	// A schema for all array items, or an array of schemas where order matters.
	Items map[string]map[string]BsonType

	// Default: true.
	// If true, the array may contain additional values that are not defined in the schema.
	// If false, only values that are explicitly listed in the items array may appear in
	// the array.
	//
	// If the value is a schema object, any additional fields must validate against the
	// schema.
	//
	// Note: The additionalItems field only affects array schemas that have an array-valued
	// items field. If the items field is a single schema object, additionalItems has no
	// effect.
	AdditionalItemsBoolIsSet bool
	AdditionalItemsBoolValue bool
	AdditionalItemsMap       map[string]map[string]BsonType

	// The maximum length of the array.
	MaxItems int64

	// The minimum length of the array.
	MinItems       int64
	MinItemsHasSet bool

	// Default: false
	// If true, each item in the array must be unique. If false, multiple array items may
	// be identical.
	UniqueItems bool
}

func (el *TypeBsonArray) Verify(value interface{}) (err error) {
	err = el.verifyParent(value)
	if err != nil {
		return
	}

	if el.Enum.values != nil {
		err = errors.New("enum is not compatible with the bson array type")
		return
	}

	err = el.VerifyType(value)
	if err != nil {
		return
	}

	err = el.VerifyMaxItems(value)
	if err != nil {
		return
	}

	err = el.VerifyMinItems(value)
	if err != nil {
		return
	}

	err = el.verifyItems(value)
	return
}

func (el *TypeBsonArray) VerifyType(value interface{}) (err error) {
	if value == nil && el.Enum.values != nil {
		err = el.Enum.Verify(value)
		return
	}
	err = el.parentVerifyInterfaceTypeIsArray(value)
	return
}

func (el *TypeBsonArray) VerifyMaxItems(value interface{}) (err error) {

	if el.MaxItems == 0 {
		return
	}

	switch converted := value.(type) {
	case nil:
	case map[string]interface{}:
		if len(converted) > int(el.MaxItems) {
			err = errors.New("the maximum number of items must be respected")
			return
		}
	default:
		err = errors.New("wrong type. value must be a map[string]interface{}")
	}

	return
}

func (el *TypeBsonArray) VerifyMinItems(value interface{}) (err error) {

	if el.MinItemsHasSet == false {
		return
	}

	switch converted := value.(type) {
	case nil:
	case map[string]interface{}:
		if int(el.MinItems) > len(converted) {
			err = errors.New("the minimum number of items must be respected")
			return
		}
	default:
		err = errors.New("wrong type. value must be a map[string]interface{}")
	}

	return
}

func (el *TypeBsonArray) verifyItems(value interface{}) (err error) {
	var element Element

	switch value.(type) {
	case []map[string]interface{}:
		for _, dataItemValue := range value.([]map[string]interface{}) {
			err = element.VerifyDocumentByProperties(&el.Items, dataItemValue)
			if err != nil {
				return
			}
		}
	default:
		err = errors.New("value must be a map[string]interface{}")
	}

	return
}

func (el *TypeBsonArray) getTypeString() string {
	return "array"
}

func (el *TypeBsonArray) Populate(schema map[string]interface{}) (err error) {
	err = el.populateGeneric(schema)
	if err != nil {
		return
	}

	//err = el.Enum.turnValuesIntoInt()
	//if err != nil {
	//  return
	//}

	el.MaxItems, err = el.getPropertyMaxItems(schema)
	if err != nil {
		return
	}

	el.MinItemsHasSet, el.MinItems, err = el.getPropertyMinItems(schema)
	if err != nil {
		return
	}

	el.UniqueItems, err = el.getPropertyUniqueItems(schema)
	if err != nil {
		return
	}

	el.Items, err = el.PopulateItens(schema)
	if err != nil {
		return
	}

	el.getPropertyAdditionalItens(schema)

	//todo: AdditionalItems
	//todo: UniqueItems

	return
}

func (el *TypeBsonArray) PopulateItens(schema map[string]interface{}) (items map[string]map[string]BsonType, err error) {

	var found bool

	_, found = schema["items"].(map[string]interface{})
	if found == false {
		return
	}

	items = make(map[string]map[string]BsonType)

	var key string
	var newSchemaMap = make(map[string]interface{})
	var newSchemaArray = make([]interface{}, 0)
	var element Element

	switch schema["items"].(type) {
	case []interface{}:
		newSchemaArray, _ = schema["items"].([]interface{})
		_ = newSchemaArray
	case map[string]interface{}:
		newSchemaMap = element.filterSchemaElements(schema["items"].(map[string]interface{}))
		newSchemaMap, found = newSchemaMap["properties"].(map[string]interface{})
		if found == false {
			err = errors.New("bsonType 'array' has key 'items', but not has key 'items.properties'")
			return
		}

		for schemaCellKey, schemaCell := range newSchemaMap {
			var typesInCell []string
			typesInCell, err = element.getPropertyBsonType(schemaCell.(map[string]interface{}))

			for _, currentType := range typesInCell {
				if key != "" {
					schemaCellKey = key + "." + schemaCellKey
				}
				err = element.typeStringToTypeObjectPopulated(&items, schemaCellKey, currentType, schemaCell.(map[string]interface{}))
				if err != nil {
					return
				}
			}
		}
	}

	return
}

func (el *TypeBsonArray) getPropertyMaxItems(schema map[string]interface{}) (maxItems int64, err error) {
	var found bool

	_, found = schema["maxItems"]
	if found == false {
		return
	}

	maxItems, err = el.getPropertyAsInt64(schema, "maxItems")
	return
}

func (el *TypeBsonArray) getPropertyMinItems(schema map[string]interface{}) (set bool, minItems int64, err error) {
	var found bool

	_, found = schema["minItems"]
	if found == false {
		return
	}

	set = true
	minItems, err = el.getPropertyAsInt64(schema, "minItems")
	return
}

func (el *TypeBsonArray) getPropertyUniqueItems(schema map[string]interface{}) (uniqueItems bool, err error) {
	var found bool

	_, found = schema["uniqueItems"]
	if found == false {
		return
	}

	uniqueItems, err = el.getPropertyAsBool(schema, "uniqueItems")
	return
}

// AdditionalItemsBoolIsSet bool
// AdditionalItemsBoolValue bool
// AdditionalItemsMap map[string]map[string]BsonType
func (el *TypeBsonArray) getPropertyAdditionalItens(schema map[string]interface{}) (boolIsSet bool, boolValue bool, itemsMap map[string]map[string]BsonType, err error) {
	var found bool

	_, found = schema["additionalProperties"]
	if found == false {
		return
	}

	var value interface{}

	value, found = schema["additionalProperties"]
	if found == false {
		return
	}

	switch converted := value.(type) {
	case bool:
		boolValue = converted
		boolIsSet = true
		return
	case string:
		boolValue, err = strconv.ParseBool(strings.ToLower(converted))
		if err == nil {
			boolIsSet = true
		}
		return
	case map[string]interface{}:
		boolValue = false
		boolIsSet = false
		itemsMap = value.(map[string]map[string]BsonType)
	}

	if reflect.ValueOf(value).Kind() == reflect.Bool {
		boolValue = value.(bool)
		boolIsSet = true
		return
	}

	if reflect.ValueOf(value).Kind() == reflect.String {
		boolValue, err = strconv.ParseBool(strings.ToLower(value.(string)))
		if err == nil {
			boolIsSet = true
		}
		return
	}

	return
}
