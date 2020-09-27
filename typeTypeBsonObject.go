package iotmakerdbmongodbutilschema

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
	Properties map[string]Element

	// The minimum number of fields allowed in the object.
	MinProperties int64

	// The maximum number of fields allowed in the object.
	MaxProperties int64

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
	AdditionalProperties AdditionalProperties

	// Specify property and schema dependencies.
	// https://www.mongodb.com/blog/post/json-schema-validation--dependencies-you-can-depend-on
	Dependencies map[string]Element
}

func (el *TypeBsonObject) getTypeString() string {
	return "object"
}

func (el *TypeBsonObject) Populate(schema map[string]interface{}) (err error) {
	err = el.populateGeneric(schema)
	return
}

func (el *TypeBsonObject) Verify(value interface{}) (err error) {

	//el.mountPatternPropertiesAdditional()
	//el.mountPatternPropertiesPattern(keyList)
	err = el.verifyMaxProperties()
	if err != nil {
		return
	}

	err = el.verifyMinProperties()
	return
}
