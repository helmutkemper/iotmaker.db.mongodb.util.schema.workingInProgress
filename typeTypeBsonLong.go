package iotmakerdbmongodbutilschema

import (
	"errors"
	"strconv"
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
type TypeBsonLong struct {
	TypeBsonCommonToAllTypes

	// An integer divisor of the field value. For example, if multipleOf is set to 3, 6 is
	// a valid value but 7 is not.
	MultipleOf int64

	// The maximum value of the number.
	Maximum int64

	// Default: false
	// If true, the field value must be strictly less than the maximum value.
	// If false, the field value may also be equal to the maximum value.
	ExclusiveMaximum bool

	// The minimum value of the number.
	Minimum       int64
	MinimumHasSet bool

	// Default: false
	// If true, the field value must be strictly greater than the minimum value.
	// If false, the field value may also be equal to the minimum value.
	ExclusiveMinimum bool
}

func (el *TypeBsonLong) Verify(value interface{}) (err error) {
	if value != nil {
		value, err = el.TypeBsonCommonToAllTypes.parentConvertInterfaceToInt64(value)
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

func (el *TypeBsonLong) VerifyType(value interface{}) (err error) {
	if value == nil && el.Enum.values == nil {
		return
	}

	if el.Enum.values != nil {
		err = el.Enum.Verify(value)
		return
	}
	err = el.parentVerifyInterfaceTypeIsInt64(value)
	return
}

func (el *TypeBsonLong) VerifyMultipleOf(value interface{}) (err error) {
	if value == nil {
		return
	}

	var module int64
	if el.MultipleOf == 0 {
		return
	}

	var converted int64
	converted, err = el.parentConvertInterfaceToInt64(value)
	if err != nil {
		return
	}

	module = converted % el.MultipleOf
	if module != 0 {
		err = errors.New("number must be multiple of " + strconv.FormatInt(el.MultipleOf, 10))
	}

	return
}

func (el *TypeBsonLong) VerifyMaximum(value interface{}) (err error) {
	if value == nil {
		return
	}

	if el.Maximum == 0 {
		return
	}

	var converted int64
	converted, err = el.parentConvertInterfaceToInt64(value)
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

func (el *TypeBsonLong) VerifyMinimum(value interface{}) (err error) {
	if value == nil {
		return
	}

	if el.MinimumHasSet == false {
		return
	}

	var converted int64
	converted, err = el.parentConvertInterfaceToInt64(value)
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

func (el *TypeBsonLong) getTypeString() string {
	return "long"
}

func (el *TypeBsonLong) Populate(schema map[string]interface{}) (err error) {
	err = el.populateGeneric(schema)
	if err != nil {
		return
	}

	err = el.Enum.turnValuesIntoInt64()
	if err != nil {
		return
	}

	var multipleOf int64
	var maximum int64
	var minimum int64
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

func (el *TypeBsonLong) getPropertyMultipleOf(schema map[string]interface{}) (multipleOf int64, err error) {
	var found bool

	_, found = schema["multipleOf"]
	if found == false {
		return
	}

	multipleOf, err = el.getPropertyAsInt64(schema, "multipleOf")
	return
}

func (el *TypeBsonLong) getPropertyMaximum(schema map[string]interface{}) (maximum int64, err error) {
	var found bool

	_, found = schema["maximum"]
	if found == false {
		return
	}

	maximum, err = el.getPropertyAsInt64(schema, "maximum")
	return
}

func (el *TypeBsonLong) getPropertyExclusiveMaximum(schema map[string]interface{}) (exclusiveMaximum bool, err error) {
	var found bool

	_, found = schema["exclusiveMaximum"]
	if found == false {
		return
	}

	exclusiveMaximum, err = el.getPropertyAsBool(schema, "exclusiveMaximum")
	return
}

func (el *TypeBsonLong) getPropertyMinimum(schema map[string]interface{}) (set bool, minimum int64, err error) {
	var found bool

	_, found = schema["minimum"]
	if found == false {
		return
	}

	set = true
	minimum, err = el.getPropertyAsInt64(schema, "minimum")
	return
}

func (el *TypeBsonLong) getPropertyExclusiveMinimum(schema map[string]interface{}) (exclusiveMinimum bool, err error) {
	var found bool

	_, found = schema["exclusiveMinimum"]
	if found == false {
		return
	}

	exclusiveMinimum, err = el.getPropertyAsBool(schema, "exclusiveMinimum")
	return
}
