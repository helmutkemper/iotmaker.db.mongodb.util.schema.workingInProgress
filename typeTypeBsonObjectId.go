package iotmakerdbmongodbutilschema

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The numeric schema type configures the content of numeric fields, such as integers and
// decimals.
// For more information, see the official JSON Schema numeric guide.
// https://json-schema.org/understanding-json-schema/reference/numeric.html
//
//   Example:
//   {
//     "bsonType": "objectId",
//   }
type TypeBsonObjectId struct {
	TypeBsonCommonToAllTypes
}

func (el *TypeBsonObjectId) Verify(value interface{}) (err error) {
	err = el.verifyParent(value)
	if err != nil {
		return
	}

	err = el.VerifyType(value)
	if err != nil {
		return
	}

	return
}

func (el *TypeBsonObjectId) VerifyType(value interface{}) (err error) {
	if value == nil {
		return
	}

	switch converted := value.(type) {
	case primitive.ObjectID:
		if converted.IsZero() == true {
			err = errors.New("objectId not be null")
		}
	default:
		err = errors.New("type must be a objectID")
	}

	return
}

func (el *TypeBsonObjectId) getTypeString() string {
	return "objectId"
}

func (el *TypeBsonObjectId) Populate(schema map[string]interface{}) (err error) {
	err = el.populateGeneric(schema)
	if err != nil {
		return
	}

	return
}
