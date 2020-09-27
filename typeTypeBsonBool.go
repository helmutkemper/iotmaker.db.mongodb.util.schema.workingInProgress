package iotmakerdbmongodbutilschema

// The boolean schema type configures the content of fields that are either true or
// false.
// For more information, see the official JSON Schema boolean guide.
// https://json-schema.org/understanding-json-schema/reference/boolean.html
type TypeBsonBool struct {
	TypeBsonCommonToAllTypes
}

func (el *TypeBsonBool) getTypeString() string {
	return "bool"
}

func (el *TypeBsonBool) Populate(schema map[string]interface{}) (err error) {
	err = el.populateGeneric(schema)
	return
}

func (el *TypeBsonBool) Verify(value interface{}) (err error) {
	err = el.verifyParent(value)
	return
}
