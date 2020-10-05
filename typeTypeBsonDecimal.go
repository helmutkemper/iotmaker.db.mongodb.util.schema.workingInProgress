package iotmakerdbmongodbutilschema

import (
	"errors"
	"fmt"
)

// The numeric schema type configures the content of numeric fields, such as integers and
// decimals.
// For more information, see the official JSON Schema numeric guide.
// https://json-schema.org/understanding-json-schema/reference/numeric.html
//
//   Example:
//   {
//     "bsonType": "int" | "long" | "double" | "decimal",
//     "multipleOf": <number>,
//     "maximum": <number>,
//     "exclusiveMaximum": <boolean>,
//     "minimum": <number>,
//     "exclusiveMinimum": <boolean>
//   }
type TypeBsonDecimal struct {
	TypeBsonCommonToAllTypes

	// An integer divisor of the field value. For example, if multipleOf is set to 3, 6 is
	// a valid value but 7 is not.
	MultipleOf float32

	// The maximum value of the number.
	Maximum float32

	// Default: false
	// If true, the field value must be strictly less than the maximum value.
	// If false, the field value may also be equal to the maximum value.
	ExclusiveMaximum bool

	// The minimum value of the number.
	Minimum       float32
	MinimumHasSet bool

	// Default: false
	// If true, the field value must be strictly greater than the minimum value.
	// If false, the field value may also be equal to the minimum value.
	ExclusiveMinimum bool
}

func (el *TypeBsonDecimal) Verify(value interface{}) (err error) {
	if value != nil {
		value, err = el.TypeBsonCommonToAllTypes.parentConvertInterfaceToFloat32(value)
		if err != nil {
			return
		}
	}

	err = el.verifyParent(value)
	if err != nil {
		return
	}

	err = el.VerifyType(value)
	if err != nil {
		return
	}

	err = el.VerifyMultipleOf(value)
	if err != nil {
		return
	}

	err = el.VerifyMaximum(value)
	if err != nil {
		return
	}

	err = el.VerifyMinimum(value)
	return
}

func (el *TypeBsonDecimal) VerifyType(value interface{}) (err error) {
	if value == nil && el.Enum.values == nil {
		return
	}

	if el.Enum.values != nil {
		err = el.Enum.Verify(value)
		return
	}
	err = el.parentVerifyInterfaceTypeIsFloat32(value)
	return
}

func (el *TypeBsonDecimal) VerifyMultipleOf(value interface{}) (err error) {
	var module float64
	if el.MultipleOf == 0 {
		return
	}

	var converted float32
	converted, err = el.parentConvertInterfaceToFloat32(value)
	if err != nil {
		return
	}

	module = el.round(float64(converted)/float64(el.MultipleOf), 1.0)
	if module != float64(int64(module)) {
		err = errors.New(fmt.Sprintf("number must be multiple of %1.2f", el.MultipleOf))
	}

	return
}

func (el *TypeBsonDecimal) VerifyMaximum(value interface{}) (err error) {
	if el.Maximum == 0 {
		return
	}

	var converted float32
	converted, err = el.parentConvertInterfaceToFloat32(value)
	if err != nil {
		return
	}

	if el.ExclusiveMaximum == true && converted >= el.Maximum {
		err = errors.New("maximum value exceeded")
		return
	}

	if el.ExclusiveMaximum == false && converted > el.Maximum {
		err = errors.New("maximum value exceeded")
		return
	}

	return
}

func (el *TypeBsonDecimal) VerifyMinimum(value interface{}) (err error) {
	if el.MinimumHasSet == false {
		return
	}

	var converted float32
	converted, err = el.parentConvertInterfaceToFloat32(value)
	if err != nil {
		return
	}

	if el.ExclusiveMinimum == true && el.Minimum >= converted {
		err = errors.New("expected minimum value")
		return
	}

	if el.ExclusiveMinimum == false && el.Minimum > converted {
		err = errors.New("expected minimum value")
	}

	return
}

func (el *TypeBsonDecimal) getTypeString() string {
	return "decimal"
}

func (el *TypeBsonDecimal) Populate(schema map[string]interface{}) (err error) {
	err = el.populateGeneric(schema)
	if err != nil {
		return
	}

	err = el.Enum.turnValuesIntoFloat32()
	if err != nil {
		return
	}

	var multipleOf float32
	var maximum float32
	var minimum float32
	var minimumHasSet bool

	multipleOf, err = el.getPropertyMultipleOf(schema)
	if err != nil {
		return
	}

	maximum, err = el.getPropertyMaximum(schema)
	if err != nil {
		return
	}

	el.ExclusiveMaximum, err = el.getPropertyExclusiveMaximum(schema)
	if err != nil {
		return
	}

	minimumHasSet, minimum, err = el.getPropertyMinimum(schema)
	if err != nil {
		return
	}

	el.ExclusiveMinimum, err = el.getPropertyExclusiveMinimum(schema)
	if err != nil {
		return
	}

	el.MultipleOf = multipleOf
	el.Maximum = maximum
	el.Minimum = minimum
	el.MinimumHasSet = minimumHasSet

	return
}

func (el *TypeBsonDecimal) getPropertyMultipleOf(schema map[string]interface{}) (multipleOf float32, err error) {
	var found bool

	_, found = schema["multipleOf"]
	if found == false {
		return
	}

	multipleOf, err = el.getPropertyAsFloat32(schema, "multipleOf")
	return
}

func (el *TypeBsonDecimal) getPropertyMaximum(schema map[string]interface{}) (maximum float32, err error) {
	var found bool

	_, found = schema["maximum"]
	if found == false {
		return
	}

	maximum, err = el.getPropertyAsFloat32(schema, "maximum")
	return
}

func (el *TypeBsonDecimal) getPropertyExclusiveMaximum(schema map[string]interface{}) (exclusiveMaximum bool, err error) {
	var found bool

	_, found = schema["exclusiveMaximum"]
	if found == false {
		return
	}

	exclusiveMaximum, err = el.getPropertyAsBool(schema, "exclusiveMaximum")
	return
}

func (el *TypeBsonDecimal) getPropertyMinimum(schema map[string]interface{}) (set bool, minimum float32, err error) {
	var found bool

	_, found = schema["minimum"]
	if found == false {
		return
	}

	set = true
	minimum, err = el.getPropertyAsFloat32(schema, "minimum")
	return
}

func (el *TypeBsonDecimal) getPropertyExclusiveMinimum(schema map[string]interface{}) (exclusiveMinimum bool, err error) {
	var found bool

	_, found = schema["exclusiveMinimum"]
	if found == false {
		return
	}

	exclusiveMinimum, err = el.getPropertyAsBool(schema, "exclusiveMinimum")
	return
}
