package iotmakerdbmongodbutilschema

import "errors"

func (el *TypeBsonObject) verifyMinProperties() (err error) {
	if el.MinPropertiesHasSet == true && len(el.Properties) < int(el.MinProperties) {
		err = errors.New("minimum amount of properties not achieved")
	}
	return
}
