package iotmakerdbmongodbutilschema

import (
	"errors"
	"fmt"
	"reflect"
)

// The object schema type configures the content of documents.
// For more information, see the official JSON Schema object guide.
// https://json-schema.org/understanding-json-schema/reference/object.html
//
//   Example:
//   {
//     "bsonType": "object",
//     "title": "<Type Name>",
//     "required": ["<Required Field Name>", ...],
//     "properties": {
//       "<Field Name>": <Schema Document>
//     },
//     "minProperties": <integer>,
//     "maxProperties": <integer>,
//     "patternProperties": {
//       "<Field Name Regex>": <Schema Document>
//     },
//     "additionalProperties": <boolean> | <Schema Document>,
//     "dependencies": {
//       "<Field Name>": <Schema Document> | ["<Field Name>", ...]
//     }
//   }
type TypeBsonObject struct {
	TypeBsonCommonToAllTypes

	// An object where each field maps to a field in the parent object by name. The value
	// of each field is a schema document that configures the value of the field.
	Properties map[string]map[string]BsonType

	// The minimum number of fields allowed in the object.
	MinPropertiesHasSet bool
	MinProperties       int64

	// The maximum number of fields allowed in the object.
	MaxPropertiesHasSet bool
	MaxProperties       int64

	// An object where each field is a regular expression string that maps to all fields in
	// the parent object that match. The value of each field is a schema document that
	// configures the value of matched fields.
	PatternProperties []PatternProperties

	// Default: true.
	// If true, a document may contain additional fields that are not defined in the
	// schema.
	// If false, only fields that are explicitly defined in the schema may appear in a
	// document.
	// If the value is a schema object, any additional fields must validate against the
	// schema
	AdditionalPropertiesBoolIsSet bool
	AdditionalPropertiesBoolValue bool
	AdditionalPropertiesMap       map[string]map[string]BsonType

	// Specify property and schema dependencies.
	// https://www.mongodb.com/blog/post/json-schema-validation--dependencies-you-can-depend-on
	Dependencies map[string]map[string]BsonType

	Required map[string]bool

	ErrorList []error
}

func (el *TypeBsonObject) VerifyErros() (errorList []error) {
	return el.ErrorList
}

func (el *TypeBsonObject) getTypeString() string {
	return "object"
}

func (el *TypeBsonObject) Populate(schema map[string]interface{}) (err error) {
	err = el.populateGeneric(schema)
	if err != nil {
		return
	}

	el.MinPropertiesHasSet, el.MinProperties, err = el.getPropertyMinProperties(schema)
	if err != nil {
		return
	}

	el.MaxPropertiesHasSet, el.MaxProperties, err = el.getPropertyMaxProperties(schema)
	if err != nil {
		return
	}

	err = el.populateRequired(schema)
	if err != nil {
		return
	}

	el.Properties, err = el.populateBsonType(schema)
	return
}

// processRequiredFields (English): Process the required fields
//    json schema example:
//    {
//      "bsonType": "object",
//      "title": "<Type Name>",
//      "required": ["<Required Field Name>", ...],
//      "properties": {
//        "<Field Name>": <Schema Document>
//      }
//    }
//
// processRequiredFields (Português): Processa os campos requeridos
//    exemplo de esquema json:
//    {
//      "bsonType": "object",
//      "title": "<Type Name>",
//      "required": ["<Required Field Name>", ...],
//      "properties": {
//        "<Field Name>": <Schema Document>
//      }
//    }
func (el *TypeBsonObject) populateRequired(schema map[string]interface{}) (err error) {
	return el.populateRequiredSupport(&el.Required, "", schema)
}

func (el *TypeBsonObject) populateRequiredSupport(requiredPointer *map[string]bool, key string, schema map[string]interface{}) (err error) {
	var found bool
	//var newSchema map[string]interface{}

	var requiredList []interface{}
	requiredList, found = schema["required"].([]interface{})
	if found == false {
		return
	}

	if *requiredPointer == nil {
		*requiredPointer = make(map[string]bool)
	}

	for _, requiredKeyName := range requiredList {
		if key != "" {
			//todo: verificar string
			requiredKeyName = key + "." + requiredKeyName.(string)
		}
		(*requiredPointer)[requiredKeyName.(string)] = true
	}

	// fixme: início: isto está correto, mas, o objeto só deve verificar aos campos dele ou entrar na árvore e verificar?
	//newSchema, _ = schema["properties"].(map[string]interface{})
	//for schemaCellKey, schemaCell := range newSchema {
	//  if key != "" {
	//    schemaCellKey = key + "." + schemaCellKey
	//  }
	//  err = el.populateRequiredSupport(requiredPointer, schemaCellKey, schemaCell.(map[string]interface{}))
	//}
	// fixme: fim

	return
}

func (el *TypeBsonObject) VerifyRequired(mainKey string, required map[string]bool, value interface{}) (err error) {

	var found bool

	for k, v := range required {
		if v == true {
			_, found = value.(map[string]BsonType)[k]
			if found == false {
				if mainKey != "" {
					k = mainKey + "." + k
				}

				if el.ErrorList == nil {
					el.ErrorList = make([]error, 0)
				}

				el.ErrorList = append(el.ErrorList, errors.New(k+" not found"))
				return
			}
		}
	}

	return
}

func (el *TypeBsonObject) convertGolangTypeToMongoType(goType string) (mongoType string) {
	switch goType {
	case "map":
		return "object"
	}

	return goType
}

func (el *TypeBsonObject) VerifyRules(value interface{}) {
	var err error
	var found bool
	var rules BsonType
	var key = ""

	for key, properties := range el.Properties {
		var pass = false
		var errList = make([]error, 0)
		for dataType, rule := range properties {

			if dataType == "object" {
				err = rule.Verify(value)
			} else {
				switch converted := value.(type) {
				case map[string]interface{}:
					err = rule.Verify(converted[key])
				}
			}

			if err != nil {
				if el.ErrorList == nil {
					el.ErrorList = make([]error, 0)
				}
				el.ErrorList = append(el.ErrorList, err)

			} else {
				pass = true

				if dataType == "object" {
					var t interface{}
					t = rule.ElementType
					if t == nil {
						break
					}
					r := t.(*TypeBsonObject)
					r.VerifyRules(value.(map[string]interface{})[key])
				}

				break
			}
		}

		if pass == false {
			for _, err = range errList {
				fmt.Printf("error: %v\n", err.Error())
			}
		}
	}

	switch reflect.ValueOf(value).Kind() {
	case reflect.Invalid:
		fmt.Println("case reflect.Invalid:")
		rules, found = el.Properties[key]["null"]

	case reflect.Bool:
		fmt.Println("case reflect.Bool:")
		rules, found = el.Properties[key]["bool"]

	case reflect.Int:
		fmt.Println("case reflect.Int:")
		rules, found = el.Properties[key]["int"]

	case reflect.Int8:
		fmt.Println("case reflect.Int8:")
		rules, found = el.Properties[key]["int"]

	case reflect.Int16:
		fmt.Println("case reflect.Int16:")
		rules, found = el.Properties[key]["int"]

	case reflect.Int32:
		fmt.Println("case reflect.Int32:")
		rules, found = el.Properties[key]["int"]

	case reflect.Int64:
		fmt.Println("case reflect.Int64:")
		rules, found = el.Properties[key]["long"]

	case reflect.Uint:
		fmt.Println("case reflect.Uint:")
		rules, found = el.Properties[key]["int"]

	case reflect.Uint8:
		fmt.Println("case reflect.Uint8:")
		rules, found = el.Properties[key]["int"]

	case reflect.Uint16:
		fmt.Println("case reflect.Uint16:")
		rules, found = el.Properties[key]["int"]

	case reflect.Uint32:
		fmt.Println("case reflect.Uint32:")
		rules, found = el.Properties[key]["int"]

	case reflect.Uint64:
		fmt.Println("case reflect.Uint64:")
		rules, found = el.Properties[key]["long"]

	case reflect.Uintptr:
		fmt.Println("case reflect.Uintptr:")
		rules, found = el.Properties[key]["error"]

	case reflect.Float32:
		fmt.Println("case reflect.Float32:")
		rules, found = el.Properties[key]["double"]

	case reflect.Float64:
		fmt.Println("case reflect.Float64:")
		rules, found = el.Properties[key]["double"]

	case reflect.Complex64:
		fmt.Println("case reflect.Complex64:")
		rules, found = el.Properties[key]["decimal"]

	case reflect.Complex128:
		fmt.Println("case reflect.Complex128:")
		rules, found = el.Properties[key]["decimal"]

	case reflect.Array:
		fmt.Println("case reflect.Array:")
		rules, found = el.Properties[key]["array"]

	case reflect.Chan:
		fmt.Println("case reflect.Chan:")
		rules, found = el.Properties[key]["error"]

	case reflect.Func:
		fmt.Println("case reflect.Func:")
		rules, found = el.Properties[key]["error"]

	case reflect.Interface:
		fmt.Println("case reflect.Interface:")
		rules, found = el.Properties[key]["error"]

	case reflect.Map:
		fmt.Println("case reflect.Map:")
		rules, found = el.Properties[key]["object"]

	case reflect.Ptr:
		fmt.Println("case reflect.Ptr:")
		rules, found = el.Properties[key]["error"]

	case reflect.Slice:
		fmt.Println("case reflect.Slice:")
		rules, found = el.Properties[key]["array"]

	case reflect.String:
		fmt.Println("case reflect.String:")
		rules, found = el.Properties[key]["string"]

	case reflect.Struct:
		fmt.Println("case reflect.Struct:")
		rules, found = el.Properties[key]["object"]

	case reflect.UnsafePointer:
		fmt.Println("case reflect.UnsafePointer:")
		rules, found = el.Properties[key]["error"]
	}

	if found == true {
		err = rules.Verify(value)
	}

	return
}

func (el *TypeBsonObject) _VerifyRules(value interface{}) (err error) {

	var valueAsMap map[string]interface{}

	switch converted := value.(type) {
	case map[string]interface{}:
		valueAsMap = converted

	default:
		err = errors.New("data must be a map[string]interface{}")
		return
	}

	var found bool
	var rules BsonType
	var kindAsString string

	for dataKey, dataValue := range valueAsMap {
		kindAsString = reflect.ValueOf(dataValue).Kind().String()
		kindAsString = el.convertGolangTypeToMongoType(kindAsString)

		_, found = el.Properties[dataKey]
		if found == false {
			continue
		}

		rules, found = el.Properties[dataKey][kindAsString]
		if found == false {
			err = errors.New("'" + dataKey + "' wrong data type")

			if el.ErrorList == nil {
				el.ErrorList = make([]error, 0)
			}
			el.ErrorList = append(el.ErrorList, err)
			continue
		}

		err = rules.Verify(dataValue)
		if err != nil {

			if el.ErrorList == nil {
				el.ErrorList = make([]error, 0)
			}
			el.ErrorList = append(el.ErrorList, err)
			continue
		}

		// fixme: melhorar isto - início
		if el.Required == nil {
			continue
		}

		el.Required[dataKey] = false
		// fixme: melhorar isto - fim

	}

	switch value.(type) {
	case map[string]interface{}:
		for k, v := range el.Required {
			if v == false {
				continue
			}

			_, found = value.(map[string]interface{})[k]
			if found == false {
				el.ErrorList = append(el.ErrorList, errors.New(k+" not found"))
			}
		}
	}

	for dataKey, dataValue := range el.Properties {
		for keyType, property := range el.Properties[dataKey] {
			if keyType != "object" {
				continue
			}

			var t interface{}
			t = property.ElementType
			if t == nil {
				continue
			}
			r := t.(*TypeBsonObject)
			rl := r.Required
			err = el.VerifyRequired(dataKey, rl, dataValue)
			if err != nil {
				el.ErrorList = append(el.ErrorList, err)
				continue
			}
		}
	}

	//if len(el.ErrorList) != 0 {
	//  err = errors.New("verification erros found")
	//}
	return
}

func (el *TypeBsonObject) Verify(value interface{}) (err error) {

	err = el.verifyParent(value)
	if err != nil {
		return
	}

	err = el.verifyMaxProperties()
	if err != nil {
		return
	}

	err = el.verifyMinProperties()
	if err != nil {
		return
	}

	//err = el.VerifyRules(value)
	//if err != nil {
	//	return
	//}

	return
}

func (el *TypeBsonObject) getPropertyMinProperties(schema map[string]interface{}) (set bool, minimum int64, err error) {
	var found bool

	_, found = schema["minProperties"]
	if found == false {
		return
	}

	set = true
	minimum, err = el.getPropertyAsInt64(schema, "minProperties")
	return
}

func (el *TypeBsonObject) getPropertyMaxProperties(schema map[string]interface{}) (set bool, minimum int64, err error) {
	var found bool

	_, found = schema["maxProperties"]
	if found == false {
		return
	}

	set = true
	minimum, err = el.getPropertyAsInt64(schema, "maxProperties")
	return
}

func (el *TypeBsonObject) populateBsonType(schema map[string]interface{}) (properties map[string]map[string]BsonType, err error) {

	properties = make(map[string]map[string]BsonType)
	//var typeList []string
	//typeList, err = el.getPropertyBsonTypeAsSlice(schema)
	//for _, v := range typeList {
	//  err = el.typeStringToTypeObjectPopulated(&properties, "", v, schema)
	//  if err != nil {
	//    return
	//  }
	//}

	var newSchema map[string]interface{}
	newSchema, _ = schema["properties"].(map[string]interface{})
	for schemaCellKey, schemaCell := range newSchema {

		var typesInCell []string
		typesInCell, err = el.getPropertyBsonTypeAsSlice(schemaCell.(map[string]interface{}))

		for _, currentType := range typesInCell {
			//if key != "" {
			//  schemaCellKey = key + "." + schemaCellKey
			//}
			err = el.typeStringToTypeObjectPopulated(&properties, schemaCellKey, currentType, schemaCell.(map[string]interface{}))
			if err != nil {
				return
			}
		}
	}

	return
}

func (el *TypeBsonObject) getPropertyBsonTypeAsSlice(schema map[string]interface{}) (value []string, err error) {

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

func (el *TypeBsonObject) typeStringToTypeObjectPopulated(propertiesPointer *map[string]map[string]BsonType, key string, typeString string, schema map[string]interface{}) (err error) {
	//var newSchema map[string]interface{}
	var objType InterfaceBson

	if *propertiesPointer == nil {
		*propertiesPointer = make(map[string]map[string]BsonType)
	}

	switch typeString {

	// English:
	// in case of enum, 'bsonType' can be omitted.
	// when this happens, the 'TypeBsonCommonToAllTypes' object meets the needs
	// note: 'generic' was created by me and is not provided in the documentation
	//
	// Português:
	// em caso de enum, 'bsonType' pode ser omitido.
	// quando isto acontece, o objeto 'TypeBsonCommonToAllTypes' atende as necessidades
	// nota: 'generic' foi criado por mim e não é previsto na documentação
	case "generic":
		objType = &TypeBsonGeneric{}
		err = objType.Populate(schema)
		if err != nil {
			return
		}

	case "object":
		objType = &TypeBsonObject{}
		err = objType.Populate(schema)
		if err != nil {
			return
		}

		//newSchema, _ = schema["properties"].(map[string]interface{})
		//for schemaCellKey, schemaCell := range newSchema {
		//
		//  var typesInCell []string
		//  typesInCell, err = el.getPropertyBsonTypeAsSlice(schemaCell.(map[string]interface{}))
		//
		//  for _, currentType := range typesInCell {
		//    if key != "" {
		//      schemaCellKey = key + "." + schemaCellKey
		//    }
		//    err = el.typeStringToTypeObjectPopulated(propertiesPointer, schemaCellKey, currentType, schemaCell.(map[string]interface{}))
		//    if err != nil {
		//      return
		//    }
		//  }
		//}

	case "double":
		objType = &TypeBsonDouble{}
		err = objType.Populate(schema)
		if err != nil {
			return
		}

	case "string":
		objType = &TypeBsonString{}
		err = objType.Populate(schema)
		if err != nil {
			return
		}

	case "array":
		objType = &TypeBsonArray{}
		err = objType.Populate(schema)
		if err != nil {
			return
		}

	//case "binData":
	//case "objectId":
	case "bool":
		objType = &TypeBsonBool{}
		err = objType.Populate(schema)
		if err != nil {
			return
		}

	//case "date":
	//case "null":
	//case "regex":
	//case "dbPointer":
	//case "javascript":
	//case "symbol":
	//case "javascriptWithScope":
	case "int":
		objType = &TypeBsonInt{}
		err = objType.Populate(schema)
		if err != nil {
			return
		}

	case "timestamp":
	case "long":
		objType = &TypeBsonLong{}
		err = objType.Populate(schema)
		if err != nil {
			return
		}

	case "decimal":
		objType = &TypeBsonDecimal{}
		err = objType.Populate(schema)
		if err != nil {
			return
		}

	default:
		err = errors.New("type not implemented yet")
	}

	if (*propertiesPointer)[key] == nil {
		(*propertiesPointer)[key] = make(map[string]BsonType)
	}

	(*propertiesPointer)[key][typeString] = BsonType{ElementType: objType}

	return
}
