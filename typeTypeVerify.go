package iotmakerdbmongodbutilschema

func (el BsonType) Verify(value interface{}) (err error) {

	err = el.ElementType.Verify(value)
	return
}
