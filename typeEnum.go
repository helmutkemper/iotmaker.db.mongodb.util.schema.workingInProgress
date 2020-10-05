package iotmakerdbmongodbutilschema

import (
	"errors"
	"strconv"
)

type Enum struct {
	values []interface{}
}

func (el *Enum) Verify(value interface{}) (err error) {
	if el.values == nil {
		return
	}

	for _, v := range el.values {
		if v == value {
			return
		}
	}

	err = errors.New("the value does not match any value contained in the array")
	return
}

func (el *Enum) turnValuesIntoInt() (err error) {
	if el.values == nil {
		return
	}

	for k, v := range el.values {
		switch v.(type) {
		case nil:
			el.values[k] = nil
		case string:
			el.values[k], err = strconv.ParseInt(v.(string), 10, 32)
			if err != nil {
				return
			}
		case int:
			el.values[k] = int(v.(int))
		case int8:
			el.values[k] = int(v.(int8))
		case int16:
			el.values[k] = int(v.(int16))
		case int32:
			el.values[k] = int(v.(int32))
		case int64:
			el.values[k] = int(v.(int64))
		case uint:
			el.values[k] = int(v.(uint))
		case uint8:
			el.values[k] = int(v.(uint8))
		case uint16:
			el.values[k] = int(v.(uint16))
		case uint32:
			el.values[k] = int(v.(uint32))
		case uint64:
			el.values[k] = int(v.(uint64))
		case float32:
			el.values[k] = int(v.(float32))
		case float64:
			el.values[k] = int(v.(float64))
		default:
			err = errors.New("impossible to convert in int")
			return
		}
	}
	return
}

func (el *Enum) turnValuesIntoInt64() (err error) {
	if el.values == nil {
		return
	}

	for k, v := range el.values {
		switch v.(type) {
		case nil:
			el.values[k] = nil
		case string:
			el.values[k], err = strconv.ParseInt(v.(string), 10, 64)
			if err != nil {
				return
			}
		case int:
			el.values[k] = int64(v.(int))
		case int8:
			el.values[k] = int64(v.(int8))
		case int16:
			el.values[k] = int64(v.(int16))
		case int32:
			el.values[k] = int64(v.(int32))
		case int64:
			el.values[k] = int64(v.(int64))
		case uint:
			el.values[k] = int64(v.(uint))
		case uint8:
			el.values[k] = int64(v.(uint8))
		case uint16:
			el.values[k] = int64(v.(uint16))
		case uint32:
			el.values[k] = int64(v.(uint32))
		case uint64:
			el.values[k] = int64(v.(uint64))
		case float32:
			el.values[k] = int64(v.(float32))
		case float64:
			el.values[k] = int64(v.(float64))
		default:
			err = errors.New("impossible to convert in int")
			return
		}
	}
	return
}

func (el *Enum) turnValuesIntoFloat32() (err error) {
	if el.values == nil {
		return
	}

	for k, v := range el.values {
		switch v.(type) {
		case nil:
			el.values[k] = nil
		case string:
			el.values[k], err = strconv.ParseFloat(v.(string), 32)
			if err != nil {
				return
			}
		case int:
			el.values[k] = float32(v.(int))
		case int8:
			el.values[k] = float32(v.(int8))
		case int16:
			el.values[k] = float32(v.(int16))
		case int32:
			el.values[k] = float32(v.(int32))
		case int64:
			el.values[k] = float32(v.(int64))
		case uint:
			el.values[k] = float32(v.(uint))
		case uint8:
			el.values[k] = float32(v.(uint8))
		case uint16:
			el.values[k] = float32(v.(uint16))
		case uint32:
			el.values[k] = float32(v.(uint32))
		case uint64:
			el.values[k] = float32(v.(uint64))
		case float32:
			el.values[k] = v.(float32)
		case float64:
			el.values[k] = float32(v.(float64))
		default:
			err = errors.New("impossible to convert in int")
			return
		}
	}
	return
}

func (el *Enum) turnValuesIntoFloat64() (err error) {
	if el.values == nil {
		return
	}

	for k, v := range el.values {
		switch v.(type) {
		case nil:
			el.values[k] = nil
		case string:
			el.values[k], err = strconv.ParseFloat(v.(string), 64)
			if err != nil {
				return
			}
		case int:
			el.values[k] = float64(v.(int))
		case int8:
			el.values[k] = float64(v.(int8))
		case int16:
			el.values[k] = float64(v.(int16))
		case int32:
			el.values[k] = float64(v.(int32))
		case int64:
			el.values[k] = float64(v.(int64))
		case uint:
			el.values[k] = float64(v.(uint))
		case uint8:
			el.values[k] = float64(v.(uint8))
		case uint16:
			el.values[k] = float64(v.(uint16))
		case uint32:
			el.values[k] = float64(v.(uint32))
		case uint64:
			el.values[k] = float64(v.(uint64))
		case float32:
			el.values[k] = float64(v.(float32))
		case float64:
			el.values[k] = v.(float64)
		default:
			err = errors.New("impossible to convert in int")
			return
		}
	}
	return
}
