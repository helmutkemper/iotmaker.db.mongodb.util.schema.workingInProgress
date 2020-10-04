package iotmakerdbmongodbutilschema

import "errors"

func (el BsonType) Verify(value ...interface{}) (err error) {

	err = el.ElementType.Verify(value[0])
	if err != nil {
		if len(value) > 1 {
			err = errors.New("key " + value[1].(string) + " - " + err.Error())
		}
	}

	return
}
