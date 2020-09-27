package iotmakerdbmongodbutilschema

func (el Element) AppendType(value interface{}) (err error) {
	if len(el.Type) == 0 {
		el.Type = make([]BsonType, 0)
	}

	var t BsonType
	err = t.VerifyBsonType(value)
	if err != nil {
		return
	}

	el.Type = append(el.Type)
	return
}
