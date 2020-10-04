package iotmakerdbmongodbutilschema

import "fmt"

/*
  bsonType: "object",
  required: ["name"],
  properties: {
    _id: {},
    name: {
      bsonType: ["string"],
      description: "'name' is a required string"
    },
    graduated: {
      bsonType: ["bool"],
      description: "'graduated' is an optional boolean value"
    }
  },
  dependencies: {
    graduated: {
      required: ["mailing_address"],
      properties: {
        mailing_address: {
          bsonType: ["string"]
        }
      }
    }
  }

  todo: tipo array 'MaxItems' necessita de 'has set'
  todo: tipos int, double, decimal, long 'Maximum' necessita de 'has set'
*/

type Element struct {
	TypeBsonObject
}

func (el *Element) walking(key string, properties map[string]map[string]BsonType) {
	for keyProperty, properties := range properties {
		for dataType, rule := range properties {
			key = key + "." + keyProperty

			if dataType == "object" {
				var t interface{}
				t = rule.ElementType
				if t == nil {
					break
				}
				r := t.(*TypeBsonObject)
				for _, err := range r.ErrorList {
					fmt.Printf("error (%v): %v\n", key, err.Error())
				}

				el.walking(key, r.Properties)
			}
		}
	}
}

func (el *Element) VerifyErros() {
	for _, err := range el.ErrorList {
		fmt.Printf("error: %v\n", err.Error())
	}

	for key, properties := range el.Properties {
		for dataType, rule := range properties {

			if dataType == "object" {
				var t interface{}
				t = rule.ElementType
				if t == nil {
					break
				}
				r := t.(*TypeBsonObject)
				for _, err := range r.ErrorList {
					fmt.Printf("error (%v): %v\n", key, err.Error())
				}

				el.walking(key, r.Properties)
			}

		}
	}
}

type _Element struct {
	// schema document key name
	Key string

	// The BSON type of the property the schema describes. If the property’s value can be
	// of multiple types, specify an array of BSON types.
	// Cannot be used with the type field.
	//
	// BSON types include all JSON types as well as additional types that you can reference
	// by their 'string alias' such as:
	// (https://docs.mongodb.com/manual/reference/operator/query/type/#document-type-available-types),
	//
	// objectId
	// int
	// long
	// double
	// decimal
	// date
	// timestamp
	// regex
	BsonType map[string]BsonType

	// An array of required MongoDB document key names
	Required map[string]bool

	// Properties list
	// properties[data.key][type]Element
	Properties map[string]map[string]BsonType

	// A short title or name for the data that the schema models. This field is used for
	// metadata purposes only and has no impact on schema validation.
	Title string

	// A detailed description of the data that the schema models. This field is used for
	// metadata purposes only and has no impact on schema validation.
	Description string

	ErrorList []error
}

// mountPatternPropertiesPattern (English): Receives the name of all keys contained in the
// object and populates the rules by regular expression from PatternProperties
//
// mountPatternPropertiesPattern (Português): Recebe o nome de todas as chaves contidas no
// objeto e popula as regras por expressão regular a partir de PatternProperties
func (el *Element) mountPatternPropertiesPattern(keyList []string) {
	if len(keyList) == 0 {
		return
	}

	/*if len(el.PatternProperties) == 0 {
	    return
	  }

	  var err error
	  var propertyList []Element
	  for _, keyValue := range keyList {
	    for _, pattern := range el.PatternProperties {
	      propertyList, err = pattern.GetMatch(keyValue)
	      if err != nil {
	        continue
	      }

	      el.Properties = append(el.Properties, propertyList...)
	    }
	  }*/
}

func (el *Element) verifyRequired() (requiredList []string, err error) {
	/*
	   requiredList = make([]string, 0)
	   for _, requiredKey := range el.Required {
	     var pass = false
	     for _, property := range el.Properties {
	       if property.Key == requiredKey {
	         pass = true
	         break
	       }
	     }

	     if pass == false {
	       requiredList = append(requiredList, requiredKey)
	     }
	   }

	   if len(requiredList) != 0 {
	     err = errors.New("required keys not found")
	   }
	*/
	return
}
