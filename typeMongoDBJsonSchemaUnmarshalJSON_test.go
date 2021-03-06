package iotmakerdbmongodbutilschema

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"runtime/debug"
	"time"
)

func _ExampleMongoDBJsonSchema_UnmarshalJSON() {

	var jsonSchema = `
  {
    "validator": {
      "$jsonSchema": {
        "title": "main schema",
        "bsonType": "object",
        "required": [
          "name"
        ],
        "properties": {
          "_id": {
            "title": "ObjectID",
            "description": "MongoDB ObjectID"
          },
          "stringCompleteKey_1": {
            "bsonType": "string",
            "title": "string complete key 1",
            "description": "complete string key 1 for test",
            "maxLength": 20,
            "minLength": 3,
            "pattern": "^[a-z][a-z0-9]+$"
          },
          "stringCompleteKey_2": {
            "bsonType": "string",
            "title": "string complete key 2",
            "description": "complete string key 2 for test",
            "enum": ["um", "dois", "três", null]
          },
          "intComplexKey_1": {
            "bsonType": "int",
            "title": "int Complex Key 1",
            "description": "int Complex Key test 1",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": false,
            "minimum": 10,
            "exclusiveMinimum": false
          },
          "intComplexKey_2": {
            "bsonType": "int",
            "title": "int Complex Key 2",
            "description": "int Complex Key test 2",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": true,
            "minimum": 10,
            "exclusiveMinimum": true
          },
          "intComplexKey_3": {
            "bsonType": "int",
            "title": "int Complex Key 3",
            "description": "int Complex Key test 3",
            "enum": [1, 10, 100]
          },
          "longComplexKey_1": {
            "bsonType": "long",
            "title": "long Complex Key 1",
            "description": "long Complex Key test 1",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": false,
            "minimum": 10,
            "exclusiveMinimum": false
          },
          "longComplexKey_2": {
            "bsonType": "long",
            "title": "long Complex Key 2",
            "description": "long Complex Key test 2",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": true,
            "minimum": 10,
            "exclusiveMinimum": true
          },
          "longComplexKey_3": {
            "bsonType": "long",
            "title": "long Complex Key 3",
            "description": "long Complex Key test 3",
            "enum": [2, 20, 200]
          },
          "decimalComplexKey_1": {
            "bsonType": "decimal",
            "title": "decimal Complex Key 1",
            "description": "decimal Complex Key test 1",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": false,
            "minimum": 8.8,
            "exclusiveMinimum": false
          },
          "decimalComplexKey_2": {
            "bsonType": "decimal",
            "title": "decimal Complex Key 2",
            "description": "decimal Complex Key test 2",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": true,
            "minimum": 8.8,
            "exclusiveMinimum": true
          },
          "decimalComplexKey_3": {
            "bsonType": "decimal",
            "title": "decimal Complex Key 3",
            "description": "decimal Complex Key test 3",
            "enum": [1.1, 1.2, 1.3, 1.4]
          },
          "doubleComplexKey_1": {
            "bsonType": "double",
            "title": "double Complex Key 1",
            "description": "double Complex Key test 1",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": false,
            "minimum": 8.8,
            "exclusiveMinimum": false
          },
          "doubleComplexKey_2": {
            "bsonType": "double",
            "title": "double Complex Key 2",
            "description": "double Complex Key test 2",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": true,
            "minimum": 8.8,
            "exclusiveMinimum": true
          },
          "doubleComplexKey_3": {
            "bsonType": "double",
            "title": "double Complex Key 2",
            "description": "double Complex Key test 2",
            "enum": [4.13566743E-15, 6.674184E-11, 3.14159265358979]
          },
          "arrayComplex_1": {
            "bsonType": "array",
            "items": {
              "$jsonSchema": {
                "title": "sub schema",
                "bsonType": "object",
                "required": [
                  "name"
                ],
                "properties": {
                  "title": {
                    "bsonType": "string",
                    "title": "title",
                    "description": "'title' is a required string",
                    "enum": ["Dr.", "Dra.", null]
                  },
                  "name": {
                    "bsonType": "string",
                    "title": "name text",
                    "description": "'name' is an optional boolean value",
                    "pattern": "^[A-Z][a-z]+\\s+[A-Z][a-z]+$"
                  }
                }
              }
            }
          },
          "simpleBool": {
            "bsonType": [
              "bool"
            ],
            "title": "title text",
            "description": "'simpleBool' is an optional boolean value"
          },
          "graduated": {
            "bsonType": [
              "bool", "int"
            ],
            "title": "title text",
            "description": "'graduated' is an optional boolean value"
          },
          "street": {
            "bsonType": "object",
            "required": [
              "name", "number"
            ],
            "description": "street data",
            "properties": {
              "name": {
                "description": "street name",
                "bsonType": "string"
              },
              "number": {
                "description": "house number",
                "bsonType": "int"
              }
            }
          }
        },
        "dependencies": {
          "graduated": ["mailing_address"]
        }
      }
    }
  }
  `

	jsonSchema = `
  {
    "validator": {
      "$jsonSchema": {
        "title": "main schema",
        "bsonType": "object",
        "required": [
          "name"
        ],
        "properties": {
          "_id": {
            "title": "ObjectID",
            "description": "MongoDB ObjectID"
          },
          "stringCompleteKey_1": {
            "bsonType": "string",
            "title": "string complete key 1",
            "description": "complete string key 1 for test",
            "maxLength": 20,
            "minLength": 3,
            "pattern": "^[a-z][a-z0-9]+$"
          },
          "stringCompleteKey_2": {
            "bsonType": "string",
            "title": "string complete key 2",
            "description": "complete string key 2 for test",
            "enum": ["um", "dois", "três", null]
          },
          "intComplexKey_1": {
            "bsonType": "int",
            "title": "int Complex Key 1",
            "description": "int Complex Key test 1",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": false,
            "minimum": 10,
            "exclusiveMinimum": false
          },
          "intComplexKey_2": {
            "bsonType": "int",
            "title": "int Complex Key 2",
            "description": "int Complex Key test 2",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": true,
            "minimum": 10,
            "exclusiveMinimum": true
          },
          "intComplexKey_3": {
            "bsonType": "int",
            "title": "int Complex Key 3",
            "description": "int Complex Key test 3",
            "enum": [1, 10, 100]
          },
          "longComplexKey_1": {
            "bsonType": "long",
            "title": "long Complex Key 1",
            "description": "long Complex Key test 1",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": false,
            "minimum": 10,
            "exclusiveMinimum": false
          },
          "longComplexKey_2": {
            "bsonType": "long",
            "title": "long Complex Key 2",
            "description": "long Complex Key test 2",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": true,
            "minimum": 10,
            "exclusiveMinimum": true
          },
          "longComplexKey_3": {
            "bsonType": "long",
            "title": "long Complex Key 3",
            "description": "long Complex Key test 3",
            "enum": [2, 20, 200]
          },
          "decimalComplexKey_1": {
            "bsonType": "decimal",
            "title": "decimal Complex Key 1",
            "description": "decimal Complex Key test 1",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": false,
            "minimum": 8.8,
            "exclusiveMinimum": false
          },
          "decimalComplexKey_2": {
            "bsonType": "decimal",
            "title": "decimal Complex Key 2",
            "description": "decimal Complex Key test 2",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": true,
            "minimum": 8.8,
            "exclusiveMinimum": true
          },
          "decimalComplexKey_3": {
            "bsonType": "decimal",
            "title": "decimal Complex Key 3",
            "description": "decimal Complex Key test 3",
            "enum": [1.1, 1.2, 1.3, 1.4]
          },
          "doubleComplexKey_1": {
            "bsonType": "double",
            "title": "double Complex Key 1",
            "description": "double Complex Key test 1",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": false,
            "minimum": 8.8,
            "exclusiveMinimum": false
          },
          "doubleComplexKey_2": {
            "bsonType": "double",
            "title": "double Complex Key 2",
            "description": "double Complex Key test 2",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": true,
            "minimum": 8.8,
            "exclusiveMinimum": true
          },
          "doubleComplexKey_3": {
            "bsonType": "double",
            "title": "double Complex Key 2",
            "description": "double Complex Key test 2",
            "enum": [4.13566743E-15, 6.674184E-11, 3.14159265358979]
          },
          "arrayComplex_1": {
            "bsonType": "array",
            "items": {
              "title": "sub schema",
              "bsonType": "object",
              "required": [
                "name"
              ],
              "properties": {
                "title": {
                  "bsonType": "string",
                  "title": "title",
                  "description": "'title' is a required string",
                  "enum": ["Dr.", "Dra.", null]
                },
                "name": {
                  "bsonType": "string",
                  "title": "name text",
                  "description": "'name' is an optional boolean value",
                  "pattern": "^[A-Z][a-z]+\\s+[A-Z][a-z]+$"
                }
              }
            }
          },
          "graduated": {
            "bsonType": [
              "bool", "int"
            ],
            "title": "title text",
            "description": "'graduated' is an optional boolean value"
          },
          "street": {
            "bsonType": "object",
            "required": [
              "name", "number"
            ],
            "description": "street data",
            "properties": {
              "name": {
                "description": "street name",
                "bsonType": "string"
              },
              "number": {
                "description": "house number",
                "bsonType": "int"
              }
            }
          }
        },
        "dependencies": {
          "graduated": ["mailing_address"]
        }
      }
    }
  }
  `

	jsonSchema = `
  {
    "validator": {
      "$jsonSchema": {
        "title": "main schema",
        "bsonType": "object",
        "required": [
          "object", "doubleComplexKey_3", "stringCompleteKey_2"
        ],
        "properties": {
          "_id": {
            "title": "ObjectID",
            "description": "MongoDB ObjectID"
          },
          "object": {
            "title": "object test",
            "bsonType": "object",
            "required": ["name"],
            "properties": {
              "name": {
                "bsonType": "string"
              }
            }
          },
          "stringCompleteKey_1": {
            "bsonType": "string",
            "title": "string complete key 1",
            "description": "complete string key 1 for test",
            "maxLength": 20,
            "minLength": 3,
            "pattern": "^[a-z][a-z0-9]+$"
          },
          "stringCompleteKey_2": {
            "bsonType": "string",
            "title": "string complete key 2",
            "description": "complete string key 2 for test",
            "enum": ["um", "dois", "três", null]
          },
          "intComplexKey_1": {
            "bsonType": "int",
            "title": "int Complex Key 1",
            "description": "int Complex Key test 1",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": false,
            "minimum": 10,
            "exclusiveMinimum": false
          },
          "intComplexKey_2": {
            "bsonType": "int",
            "title": "int Complex Key 2",
            "description": "int Complex Key test 2",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": true,
            "minimum": 10,
            "exclusiveMinimum": true
          },
          "intComplexKey_3": {
            "bsonType": "int",
            "title": "int Complex Key 3",
            "description": "int Complex Key test 3",
            "enum": [1, 10, 100]
          },
          "longComplexKey_1": {
            "bsonType": "long",
            "title": "long Complex Key 1",
            "description": "long Complex Key test 1",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": false,
            "minimum": 10,
            "exclusiveMinimum": false
          },
          "longComplexKey_2": {
            "bsonType": "long",
            "title": "long Complex Key 2",
            "description": "long Complex Key test 2",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": true,
            "minimum": 10,
            "exclusiveMinimum": true
          },
          "longComplexKey_3": {
            "bsonType": "long",
            "title": "long Complex Key 3",
            "description": "long Complex Key test 3",
            "enum": [2, 20, 200]
          },
          "decimalComplexKey_1": {
            "bsonType": "decimal",
            "title": "decimal Complex Key 1",
            "description": "decimal Complex Key test 1",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": false,
            "minimum": 8.8,
            "exclusiveMinimum": false
          },
          "decimalComplexKey_2": {
            "bsonType": "decimal",
            "title": "decimal Complex Key 2",
            "description": "decimal Complex Key test 2",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": true,
            "minimum": 8.8,
            "exclusiveMinimum": true
          },
          "decimalComplexKey_3": {
            "bsonType": "decimal",
            "title": "decimal Complex Key 3",
            "description": "decimal Complex Key test 3",
            "enum": [1.1, 1.2, 1.3, 1.4]
          },
          "doubleComplexKey_1": {
            "bsonType": "double",
            "title": "double Complex Key 1",
            "description": "double Complex Key test 1",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": false,
            "minimum": 8.8,
            "exclusiveMinimum": false
          },
          "doubleComplexKey_2": {
            "bsonType": "double",
            "title": "double Complex Key 2",
            "description": "double Complex Key test 2",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": true,
            "minimum": 8.8,
            "exclusiveMinimum": true
          },
          "doubleComplexKey_3": {
            "bsonType": "double",
            "title": "double Complex Key 2",
            "description": "double Complex Key test 2",
            "enum": [4.13566743E-15, 6.674184E-11, 3.14159265358979]
          },
          "arrayComplex_1": {
            "bsonType": "array",
            "items": {
              "title": "sub schema",
              "bsonType": "object",
              "required": [
                "name"
              ],
              "properties": {
                "title": {
                  "bsonType": "string",
                  "title": "title",
                  "description": "'title' is a required string",
                  "enum": ["Dr.", "Dra.", null]
                },
                "name": {
                  "bsonType": "string",
                  "title": "name text",
                  "description": "'name' is an optional boolean value",
                  "pattern": "^[A-Z][a-z]+\\s+[A-Z][a-z]+$"
                }
              }
            },
            "additionalProperties": {
              "bsonType": "array",
              "items": {
                "title": "sub schema",
                "bsonType": "object",
                "required": [
                  "name"
                ],
                "properties": {
                  "title": {
                    "bsonType": "string",
                    "title": "title",
                    "description": "'title' is a required string",
                    "enum": ["Dr.", "Dra.", null]
                  },
                  "name": {
                    "bsonType": "string",
                    "title": "name text",
                    "description": "'name' is an optional boolean value",
                    "pattern": "^[A-Z][a-z]+\\s+[A-Z][a-z]+$"
                  }
                }
              }
            }
          },
          "graduated": {
            "bsonType": [
              "bool", "int"
            ],
            "title": "title text",
            "description": "'graduated' is an optional boolean value"
          },
          "street": {
            "bsonType": "object",
            "required": [
              "name", "number"
            ],
            "description": "street data",
            "properties": {
              "name": {
                "description": "street name",
                "bsonType": "string"
              },
              "number": {
                "description": "house number",
                "bsonType": "int"
              }
            }
          }
        },
        "dependencies": {
          "graduated": ["mailing_address"]
        }
      }
    }
  }
  `

	//jsonSchema = `
	//{
	//  "validator": {
	//    "$jsonSchema": {
	//      "title": "main schema",
	//      "bsonType": "object",
	//      "required": [
	//        "name", "object"
	//      ],
	//      "properties": {
	//        "_id": {
	//          "title": "ObjectID",
	//          "description": "MongoDB ObjectID"
	//        },
	//        "object": {
	//          "title": "object test",
	//          "bsonType": "object",
	//          "required": ["name"],
	//          "properties": {
	//            "name": {
	//              "bsonType": "string"
	//            }
	//          }
	//        },
	//        "arrayComplex_1": {
	//          "bsonType": "array",
	//          "items": {
	//            "title": "sub schema",
	//            "bsonType": "object",
	//            "required": [
	//              "name"
	//            ],
	//            "properties": {
	//              "title": {
	//                "bsonType": "string",
	//                "title": "title",
	//                "description": "'title' is a required string",
	//                "enum": ["Dr.", "Dra.", null]
	//              },
	//              "name": {
	//                "bsonType": "string",
	//                "title": "name text",
	//                "description": "'name' is an optional boolean value",
	//                "pattern": "^[A-Z][a-z]+\\s+[A-Z][a-z]+$"
	//              }
	//            }
	//          }
	//        },
	//        "street": {
	//          "bsonType": "object",
	//          "required": [
	//            "name", "number"
	//          ],
	//          "description": "street data",
	//          "properties": {
	//            "name": {
	//              "description": "street name",
	//              "bsonType": "string"
	//            },
	//            "number": {
	//              "description": "house number",
	//              "bsonType": "int"
	//            }
	//          }
	//        }
	//      }
	//    }
	//  }
	//}
	//`

	var err error
	var schema = MongoDBJsonSchema{}
	err = schema.UnmarshalJSON([]byte(jsonSchema))
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		debug.PrintStack()
	}

	mongoData := map[string]interface{}{
		"_id":                 primitive.ObjectID([12]byte{0x5f, 0x49, 0xa1, 0x33, 0xa8, 0xf1, 0x30, 0x21, 0x42, 0xba, 0x60, 0x69}),
		"stringCompleteKey_1": "dinossauro",
		"stringCompleteKey_2": nil,
		"intComplexKey_1":     10,
		"intComplexKey_2":     40,
		"intComplexKey_3":     100,
		"longComplexKey_1":    50,
		"longComplexKey_2":    45,
		"longComplexKey_3":    2,
		"decimalComplexKey_1": 22,
		"decimalComplexKey_2": 19.8,
		"decimalComplexKey_3": 1.3,
		"doubleComplexKey_1":  22,
		"doubleComplexKey_2":  15.4,
		"doubleComplexKey_3":  6.674184e-11,
		"arrayComplex_1": []map[string]interface{}{
			{
				"title": "Dr.",
				"name":  "Dino Sauro",
			},
		},
		"object": map[string]interface{}{
			"name": "1",
		},
		"street": map[string]interface{}{
			"name":   "Dino Sauro",
			"number": 123,
		},
	}
	schema.VerifyRules(mongoData)
	schema.VerifyErros()

	// Output:
	//
}

func UnmarshalJSON0() {

	var jsonSchema string

	jsonSchema = `
  {
    "validator": {
      "$jsonSchema": {
        "title": "main schema",
        "bsonType": "object",
        "required": [
          "object", 
          "doubleComplexKey_3", 
          "stringCompleteKey_2",
          "simpleBool"
        ],
        "properties": {
          "_id": {
            "title": "ObjectID",
            "description": "MongoDB ObjectID"
          },
          "object": {
            "title": "object test",
            "bsonType": "object",
            "required": ["name"],
            "properties": {
              "name": {
                "bsonType": "string"
              }
            }
          },
          "stringCompleteKey_1": {
            "bsonType": "string",
            "title": "string complete key 1",
            "description": "complete string key 1 for test",
            "maxLength": 20,
            "minLength": 3,
            "pattern": "^[a-z][a-z0-9]+$"
          },
          "stringCompleteKey_2": {
            "bsonType": "string",
            "title": "string complete key 2",
            "description": "complete string key 2 for test",
            "enum": ["um", "dois", "três", null]
          },
          "intComplexKey_1": {
            "bsonType": "int",
            "title": "int Complex Key 1",
            "description": "int Complex Key test 1",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": false,
            "minimum": 10,
            "exclusiveMinimum": false
          },
          "intComplexKey_2": {
            "bsonType": "int",
            "title": "int Complex Key 2",
            "description": "int Complex Key test 2",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": true,
            "minimum": 10,
            "exclusiveMinimum": true
          },
          "intComplexKey_3": {
            "bsonType": "int",
            "title": "int Complex Key 3",
            "description": "int Complex Key test 3",
            "enum": [1, 10, 100]
          },
          "longComplexKey_1": {
            "bsonType": "long",
            "title": "long Complex Key 1",
            "description": "long Complex Key test 1",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": false,
            "minimum": 10,
            "exclusiveMinimum": false
          },
          "longComplexKey_2": {
            "bsonType": "long",
            "title": "long Complex Key 2",
            "description": "long Complex Key test 2",
            "multipleOf": 5,
            "maximum": 50,
            "exclusiveMaximum": true,
            "minimum": 10,
            "exclusiveMinimum": true
          },
          "longComplexKey_3": {
            "bsonType": "long",
            "title": "long Complex Key 3",
            "description": "long Complex Key test 3",
            "enum": [2, 20, 200]
          },
          "decimalComplexKey_1": {
            "bsonType": "decimal",
            "title": "decimal Complex Key 1",
            "description": "decimal Complex Key test 1",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": false,
            "minimum": 8.8,
            "exclusiveMinimum": false
          },
          "decimalComplexKey_2": {
            "bsonType": "decimal",
            "title": "decimal Complex Key 2",
            "description": "decimal Complex Key test 2",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": true,
            "minimum": 8.8,
            "exclusiveMinimum": true
          },
          "decimalComplexKey_3": {
            "bsonType": "decimal",
            "title": "decimal Complex Key 3",
            "description": "decimal Complex Key test 3",
            "enum": [1.1, 1.2, 1.3, 1.4]
          },
          "doubleComplexKey_1": {
            "bsonType": "double",
            "title": "double Complex Key 1",
            "description": "double Complex Key test 1",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": false,
            "minimum": 8.8,
            "exclusiveMinimum": false
          },
          "doubleComplexKey_2": {
            "bsonType": "double",
            "title": "double Complex Key 2",
            "description": "double Complex Key test 2",
            "multipleOf": 2.2,
            "maximum": 22,
            "exclusiveMaximum": true,
            "minimum": 8.8,
            "exclusiveMinimum": true
          },
          "doubleComplexKey_3": {
            "bsonType": "double",
            "title": "double Complex Key 2",
            "description": "double Complex Key test 2",
            "enum": [4.13566743E-15, 6.674184E-11, 3.14159265358979]
          },
          "arrayComplex_1": {
            "bsonType": "array",
            "items": {
              "title": "sub schema",
              "bsonType": "object",
              "required": [
                "name"
              ],
              "properties": {
                "title": {
                  "bsonType": "string",
                  "title": "title",
                  "description": "'title' is a required string",
                  "enum": ["Dr.", "Dra.", null]
                },
                "name": {
                  "bsonType": "string",
                  "title": "name text",
                  "description": "'name' is an optional boolean value",
                  "pattern": "^[A-Z][a-z]+\\s+[A-Z][a-z]+$"
                }
              }
            },
            "additionalProperties": {
              "bsonType": "array",
              "items": {
                "title": "sub schema",
                "bsonType": "object",
                "required": [
                  "name"
                ],
                "properties": {
                  "title": {
                    "bsonType": "string",
                    "title": "title",
                    "description": "'title' is a required string",
                    "enum": ["Dr.", "Dra.", null]
                  },
                  "name": {
                    "bsonType": "string",
                    "title": "name text",
                    "description": "'name' is an optional boolean value",
                    "pattern": "^[A-Z][a-z]+\\s+[A-Z][a-z]+$"
                  }
                }
              }
            }
          },
          "simpleBool": {
            "bsonType": [
              "bool"
            ],
            "title": "title text",
            "description": "'simpleBool' is an optional boolean value"
          },
          "graduated": {
            "bsonType": [
              "bool", "int"
            ],
            "title": "title text",
            "description": "'graduated' is an optional boolean value"
          },
          "street": {
            "bsonType": "object",
            "required": [
              "name", "number"
            ],
            "description": "street data",
            "properties": {
              "name": {
                "description": "street name",
                "bsonType": "string"
              },
              "number": {
                "description": "house number",
                "bsonType": "int"
              }
            }
          }
        },
        "dependencies": {
          "graduated": ["mailing_address"]
        }
      }
    }
  }
  `

	var err error
	var schema = MongoDBJsonSchema{}
	err = schema.UnmarshalJSON([]byte(jsonSchema))
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		debug.PrintStack()
	}

	mongoData := map[string]interface{}{
		"_id":                 primitive.ObjectID([12]byte{0x5f, 0x49, 0xa1, 0x33, 0xa8, 0xf1, 0x30, 0x21, 0x42, 0xba, 0x60, 0x69}),
		"stringCompleteKey_1": "dinossauro",
		"stringCompleteKey_2": nil,
		"intComplexKey_1":     10,
		"intComplexKey_2":     40,
		"intComplexKey_3":     100,
		"longComplexKey_1":    50,
		"longComplexKey_2":    45,
		"longComplexKey_3":    2,
		"decimalComplexKey_1": 22,
		"decimalComplexKey_2": 19.8,
		"decimalComplexKey_3": 1.3,
		"doubleComplexKey_1":  22,
		"doubleComplexKey_2":  15.4,
		"doubleComplexKey_3":  6.674184e-11,
		"simpleBool":          false,
		"arrayComplex_1": []map[string]interface{}{
			{
				"title": "Dr.",
				"name":  "Dino Sauro",
			},
		},
		"object": map[string]interface{}{
			"name": "1",
		},
		"street": map[string]interface{}{
			"name":   "Dino Sauro",
			"number": 123,
		},
	}
	schema.VerifyRules(mongoData)
	schema.VerifyErros()
}

func UnmarshalJSON1() {

	var jsonSchema string

	jsonSchema = `
  {
    "validator": {
      "$jsonSchema": {
        "title": "main schema",
        "bsonType": "object",
        "required": [ "street" ],
        "properties": {
          "street": {
            "bsonType": "object",
            "required": [
              "name", "number"
            ],
            "description": "street data",
            "properties": {
              "name": {
                "description": "street name",
                "bsonType": "string",
                "title": "name text",
                "description": "'name' must be a string",
                "pattern": "^[A-Z][a-z]+\\s+[A-Z][a-z]+$"
              },
              "number": {
                "description": "house number",
                "bsonType": "int"
              }
            }
          }
        }
      }
    }
  }
  `
	var err error
	var schema = MongoDBJsonSchema{}
	err = schema.UnmarshalJSON([]byte(jsonSchema))
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		debug.PrintStack()
	}

	mongoData := map[string]interface{}{
		"_id": primitive.ObjectID([12]byte{0x5f, 0x49, 0xa1, 0x33, 0xa8, 0xf1, 0x30, 0x21, 0x42, 0xba, 0x60, 0x69}),
		"street": map[string]interface{}{
			"name":   "Dino Sauro",
			"number": 123,
		},
	}
	schema.VerifyRules(mongoData)
	schema.VerifyErros()
}

func UnmarshalJSON2() {

	var jsonSchema string

	jsonSchema = `
  {
    "validator": {
      "$jsonSchema": {
        "title": "main schema",
        "bsonType": "object",
        "required": [ "address" ],
        "properties": {
          "address": {
            "bsonType": "object",
            "required": [ "name", "number" ],
            "properties": {
              "street": {
                "bsonType": "object",
                "properties": {
                  "name": {
                    "description": "street name",
                    "bsonType": "string",
                    "title": "name text",
                    "description": "'name' must be a string",
                    "pattern": "^[A-Z][a-z]+\\s+[A-Z][a-z]+$"
                  },
                  "number": {
                    "description": "house number",
                    "bsonType": "int"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
  `
	var err error
	var schema = MongoDBJsonSchema{}
	err = schema.UnmarshalJSON([]byte(jsonSchema))
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		debug.PrintStack()
	}

	mongoData := map[string]interface{}{
		"_id":     primitive.ObjectID([12]byte{0x5f, 0x49, 0xa1, 0x33, 0xa8, 0xf1, 0x30, 0x21, 0x42, 0xba, 0x60, 0x69}),
		"address": "rua",
	}
	schema.VerifyRules(mongoData)
	schema.VerifyErros()

	mongoData = map[string]interface{}{
		"_id": primitive.ObjectID([12]byte{0x5f, 0x49, 0xa1, 0x33, 0xa8, 0xf1, 0x30, 0x21, 0x42, 0xba, 0x60, 0x69}),
		"address": map[string]interface{}{
			"street": map[string]interface{}{
				"name":   "Dino Sauro",
				"number": 123,
			},
		},
	}
	schema.VerifyRules(mongoData)
	schema.VerifyErros()
}

func UnmarshalJSON3() {

	var err error
	var jsonSchemaFile []byte
	jsonSchemaFile, err = ioutil.ReadFile("./schema.json")
	if err != nil {
		panic(err)
	}

	var schema = MongoDBJsonSchema{}
	err = schema.UnmarshalJSON(jsonSchemaFile)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		debug.PrintStack()
	}

	mongoData := map[string]interface{}{
		"_id":     primitive.ObjectID([12]byte{0x5f, 0x49, 0xa1, 0x33, 0xa8, 0xf1, 0x30, 0x21, 0x42, 0xba, 0x60, 0x69}),
		"ID":      123456,
		"address": "rua",
	}
	schema.VerifyRules(mongoData)
	schema.VerifyErros()

	mongoData = map[string]interface{}{
		"_id": primitive.ObjectID([12]byte{0x5f, 0x49, 0xa1, 0x33, 0xa8, 0xf1, 0x30, 0x21, 0x42, 0xba, 0x60, 0x69}),
		"ID":  123456,
		"address": map[string]interface{}{
			"street": map[string]interface{}{
				"name":   "Dino Sauro",
				"number": 123,
			},
			"Date": time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	schema.VerifyRules(mongoData)
	schema.VerifyErros()
}

func ExampleMongoDBJsonSchema_UnmarshalJSON() {

	//UnmarshalJSON0()
	//UnmarshalJSON1()
	//UnmarshalJSON2()
	UnmarshalJSON3()

	// Output:
	//
}
