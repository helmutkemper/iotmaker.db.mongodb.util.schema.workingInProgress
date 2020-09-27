package iotmakerdbmongodbutilschema

// Default: true.
// If true, a document may contain additional fields that are not defined in the
// schema.
// If false, only fields that are explicitly defined in the schema may appear in a
// document.
// If the value is a schema object, any additional fields must validate against the
// schema
type AdditionalProperties struct {
	documentMayContainAdditional int8
	properties                   []Element
}

func (el *AdditionalProperties) SetDocumentMayContainAdditional() {
	el.documentMayContainAdditional = 1
}

func (el *AdditionalProperties) SetOnlyFieldsThatAreExplicitlyDefinedInTheSchema() {
	el.documentMayContainAdditional = -1
}

func (el *AdditionalProperties) SetProperties(properties []Element) {
	if len(properties) == 0 {
		properties = make([]Element, 0)
	}
	el.properties = properties
}

func (el *AdditionalProperties) AppendProperty(property Element) {
	if len(el.properties) == 0 {
		el.properties = make([]Element, 0)
	}
	el.properties = append(el.properties, property)
}
