package iotmakerdbmongodbutilschema

// mountPatternPropertiesAdditional (English): Adds additional properties
//
// mountPatternPropertiesAdditional (Português): Adiciona as propriedades adicionais
func (el *TypeBsonObject) mountPatternPropertiesAdditional() {
	if len(el.AdditionalProperties.properties) == 0 {
		return
	}

	/*if len(el.Properties) == 0 {
	    el.Properties = make([]Element, 0)
	  }

	  el.Properties = append(el.Properties, el.AdditionalProperties.properties...)*/
}
