package iotmakerdbmongodbutilschema

import "errors"

func (el *TypeBsonObject) verifyMaxProperties() (err error) {
	if el.MaxProperties != 0 && len(el.Properties) > int(el.MaxProperties) {
		err = errors.New("maximum amount of properties exceeded")
	}
	return
}
