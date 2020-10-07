package iotmakerdbmongodbutilschema

// filterSchemaElements (English): Scans the data for the 'validator' and '$ jsonSchema'
// keys to assemble the data map.
//
// filterSchemaElements (Português): Varre o dado em busca das chaves 'validator' e
// '$jsonSchema' para montar o mapa de dados.
func (el *MongoDBJsonSchema) filterSchemaElements(schema map[string]interface{}) map[string]interface{} {
	var found bool
	_, found = schema["validator"]
	if found == true {
		schema = schema["validator"].(map[string]interface{})
	}

	_, found = schema["$jsonSchema"]
	if found == true {
		schema = schema["$jsonSchema"].(map[string]interface{})
	}

	return schema
}
