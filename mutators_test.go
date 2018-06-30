package jonson

import (
	"testing"
)

func TestMutators(t *testing.T) {
	jsn := New(56.43)
	jsn.MutateToInt()
	if jsn.GetUnsafeInt() != 56 {
		t.Errorf("jsn.GetUnsafeInt() -> %d  does not equal 56", jsn.GetUnsafeInt())
	}

	jsn.Set("76.9")


	if jsn.MutateToInt() {
		t.Errorf("Shouldn't convert \"67.9\" directly to int")
	}

	jsn.Set("78.9")
	jsn.MutateToFloat()

	if jsn.GetUnsafeFloat64() != 78.9 {
		t.Errorf("jsn.GetUnsafeFlot() -> %f  does not equal 78.9", jsn.GetUnsafeFloat64())
	}
}