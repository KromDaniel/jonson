package jonson

import (
	"testing"
)

func testStringEquals(expected, test string, t *testing.T) {
	if expected != test {
		t.Errorf("Expected output to be the string %s, instead got %s", expected, test)
	}
}
func TestPrimitiveJSON(t *testing.T) {
	jsn := NewEmptyJSON()
	jsn.Set(56)
	testStringEquals("56", jsn.ToUnsafeJSONString(), t)
	jsn.MapSet("notMap", "someKey")
	testStringEquals("56", jsn.ToUnsafeJSONString(), t)
	jsn.SliceAppend(5, 6, 7, 8)
	testStringEquals("56", jsn.ToUnsafeJSONString(), t)
	jsn.SliceAppendBegin(7, 8, 9, 10)
	testStringEquals("56", jsn.ToUnsafeJSONString(), t)
}

func TestSlice(t *testing.T) {
	jsn := NewEmptyJSONArray()
	jsn.SliceAppend(1)
	jsn.SliceAppendBegin(2, 3, 4)
	jsn.SliceAppend(4, 5, 6)
	jsn.MapSet("notMap", "someKey")
	for i, v := range []int{4, 3, 2, 1, 4, 5, 6} {
		if !EqualsDeep(New(v), jsn.At(i)) {
			t.Errorf("Expected slice at index %d to be %d but instead got %s", i, v, jsn.At(i).ToUnsafeJSONString())
		}
	}
}

func TestMap(t *testing.T) {
	jsn := NewEmptyJSONMap()
	jsn.MapSet("keyA", "valueA")
	jsn.MapSet("keyB", 67)
	testStringEquals("valueA", jsn.At("keyA").GetUnsafeString(), t)
	if !jsn.At("keyB").IsNumber() {
		t.Errorf("expected keyB to be number, instead type is %d", jsn.At("keyB").kind)
	}
}
