package iotmakerdbmongodbutilschema

import "errors"

// verifyAdditionalProperties (English): Check the rules for additional properties
//
// verifyAdditionalProperties (Português): Verifica as regras das propriedades adicionais
func (el *TypeBsonObject) verifyAdditionalProperties(keyList []string) (err error) {
	if el.AdditionalProperties.documentMayContainAdditional != -1 {
		return
	}

	for _, property := range el.Properties {
		var pass = false
		for _, keyValue := range keyList {
			if property.Key == keyValue {
				pass = true
				break
			}
		}
		if pass == false {
			err = errors.New("the '" + property.Key + "' key cannot be contained in the document. only keys contained in the mongodb scheme are accepted")
			return
		}
	}

	return
}
