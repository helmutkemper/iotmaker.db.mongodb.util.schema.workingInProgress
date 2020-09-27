package iotmakerdbmongodbutilschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func (el *Element) UnmarshalJSON(data []byte) (err error) {
	var schema = make(map[string]interface{})

	err = json.Unmarshal(data, &schema)
	if err != nil {
		return
	}

	schema = el.filterSchemaElements(schema)

	err = el.slicerAndAssemblerForBsonType(schema)
	if err != nil {
		return
	}

	//err = el.populateRequired("", schema)
	//if err != nil {
	//	return
	//}

	return
}

func (el *Element) VerifyDocument(document map[string]interface{}) (err error) {
	return el.VerifyDocumentByProperties(&el.Properties, document)
}

func (el *Element) VerifyDocumentByProperties(propertiesPointer *map[string]map[string]BsonType, document map[string]interface{}) (err error) {

	var found bool
	var properties map[string]BsonType
	var kind reflect.Kind
	var rule interface{}

	for fieldKey, fieldValue := range document {
		found, properties = el.getRules(propertiesPointer, fieldKey)

		if found == false {
			continue
		}

		var pass = false
		var errList = make([]error, 0)
		for _, v := range properties {
			err = v.Verify(fieldValue)
			if err == nil {
				pass = true
				break
			}
			errList = append(errList, err)
		}
		if pass == true {
			continue
		}

		if len(errList) != 0 {
			err = errors.New(fmt.Sprintf("document key '%v': '%v' (%v - type: %v)", fieldKey, fieldValue, errList[0].Error(), reflect.ValueOf(fieldValue).Kind().String()))
			return err
		}

		err = errors.New("0 wrong type")
		kind = reflect.ValueOf(fieldValue).Kind()
		switch kind {
		case reflect.String:
			rule, found = properties["string"]
			if found == true {
				err = rule.(BsonType).Verify(fieldValue)
				break
			}
		case reflect.Float32:
			rule, found = properties["decimal"]
			if found == true {
				err = rule.(BsonType).Verify(fieldValue)
				break
			}
			rule, found = properties["long"]
			if found == true {
				err = rule.(BsonType).Verify(fieldValue)
				break
			}
			rule, found = properties["int"]
			if found == true {
				err = rule.(BsonType).Verify(fieldValue)
				break
			}
		case reflect.Int64:
			rule, found = properties["long"]
			if found == true {
				err = rule.(BsonType).Verify(fieldValue)
				break
			}
		case reflect.Int:
			rule, found = properties["decimal"]
			if found == true {
				err = rule.(BsonType).Verify(fieldValue)
				break
			}
			rule, found = properties["long"]
			if found == true {
				err = rule.(BsonType).Verify(fieldValue)
				break
			}
			rule, found = properties["int"]
			if found == true {
				err = rule.(BsonType).Verify(fieldValue)
				break
			}
		}

		if err != nil {
			err = errors.New(fmt.Sprintf("document key '%v': '%v' (%v - type: %v)", fieldKey, fieldValue, err.Error(), reflect.ValueOf(fieldValue).Kind().String()))
			return
		}
	}

	_ = found
	_ = properties

	return
}

// getRules (English): Returns the specific rule for the key contained in the MongoDB
// data
//
// getRules (Português): Retorna a regra específica da chave contida no dado do MongoDB
func (el *Element) getRules(propertiesPointer *map[string]map[string]BsonType, key string) (found bool, properties map[string]BsonType) {
	properties, found = (*propertiesPointer)[key]
	return
}
