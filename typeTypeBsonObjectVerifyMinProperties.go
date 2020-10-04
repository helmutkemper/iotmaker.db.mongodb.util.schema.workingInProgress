package iotmakerdbmongodbutilschema

import "errors"

func (el *TypeBsonObject) verifyMinProperties() (err error) {
	if el.MinPropertiesHasSet == true && len(el.Properties) < int(el.MinProperties) {
		err = errors.New("minimum amount of properties not achieved")
	}
	return
}

func (el *TypeBsonObject) verifyType(value ...interface{}) (err error) {
	switch value[0].(type) {
	case map[string]interface{}:
	default:
		err = errors.New("wrong type")
	}
	return
}
