package iotmakerdbmongodbutilschema

import (
	"errors"
)

func (el *Element) typeStringToTypeObjectPopulated(propertiesPointer *map[string]map[string]BsonType, key string, typeString string, schema map[string]interface{}) (err error) {
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
		//	var typesInCell []string
		//	typesInCell, err = el.getPropertyBsonType(schemaCell.(map[string]interface{}))
		//
		//	for _, currentType := range typesInCell {
		//		if key != "" {
		//			schemaCellKey = key + "." + schemaCellKey
		//		}
		//		err = el.typeStringToTypeObjectPopulated(propertiesPointer, schemaCellKey, currentType, schemaCell.(map[string]interface{}))
		//		if err != nil {
		//			return
		//		}
		//	}
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
func (el *Element) populateRequired(key string, typeString string, schema map[string]interface{}) (err error) {
	var newSchema map[string]interface{}

	if el.Properties == nil {
		el.Properties = make(map[string]map[string]BsonType)
	}

	switch typeString {

	case "object":
		var objType = TypeBsonObject{}
		err = objType.Populate(schema)
		if err != nil {
			return
		}

		var requiredList []interface{}
		requiredList, _ = schema["required"].([]interface{})
		_ = requiredList

		if el.Required == nil {
			el.Required = make(map[string]bool)
		}

		if requiredList != nil {
			for _, requiredKeyName := range requiredList {
				if key != "" {
					//todo: verificar string
					requiredKeyName = key + "." + requiredKeyName.(string)
				}
				el.Required[requiredKeyName.(string)] = true
			}
		}

		newSchema, _ = schema["properties"].(map[string]interface{})
		for schemaCellKey, schemaCell := range newSchema {

			var typesInCell []string
			typesInCell, err = el.getPropertyBsonType(schemaCell.(map[string]interface{}))

			for _, currentType := range typesInCell {
				if key != "" {
					schemaCellKey = key + "." + schemaCellKey
				}
				err = el.populateRequired(schemaCellKey, currentType, schemaCell.(map[string]interface{}))
			}
		}
		return

	}
	return
}

func (el *Element) populateDependencies(schema map[string]interface{}) (err error) {

	var found bool
	_, found = schema["dependencies"]
	if found == false {
		return
	}

	return
}
