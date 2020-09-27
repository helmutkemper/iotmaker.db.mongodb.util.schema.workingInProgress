package iotmakerdbmongodbutilschema

import "testing"

func TestRound(t *testing.T) {
	var value float64

	c := TypeBsonCommonToAllTypes{}
	value = c.round(0.99999, 1.0)
	if value != 1.0 {
		t.Fail()
	}

	value = c.round(0.09999, 1.0)
	if value != 0.1 {
		t.Fail()
	}
}
