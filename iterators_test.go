package jonson

import "testing"
const testJsonString = `
[
  {
    "key": 0.8215845637650305,
    "date": "2018-06-30T13:39:19.867Z"
  },
  {
    "key": 0.8773275487707828,
    "date": "2018-06-30T13:39:19.867Z"
  },
  {
    "key": 0.8470551881353823,
    "date": "2018-06-30T13:39:19.867Z"
  },
  {
    "key": 0.5198612869871533,
    "date": "2018-06-30T13:39:19.867Z"
  },
  {
    "key": 0.14774937566969926,
    "date": "2018-06-30T13:39:19.867Z"
  }
]`
func performEqualTest(left *JSON, right *JSON, t *testing.T, shouldEqual bool) {
	if shouldEqual != EqualsDeep(left, right) {
		var testError string
		if shouldEqual {
			testError = "Should equal to"
		}else {
			testError = "Shouldn't equal to"
		}
		t.Errorf("%s %s %s", left.ToUnsafeJSONString(),testError, right.ToUnsafeJSONString())
	}
}
func TestEqualsDeep(t *testing.T) {
	left := New(5)
	right := New("5")

	performEqualTest(left, right,t,false)

	left.Set([]int{1,2,3,4,5})
	right.Set([]int{1,2,3,4,5})

	performEqualTest(left, right,t,true)

	left.At(0).Set(4)

	performEqualTest(left, right,t,false)

	left = NewEmptyJSONMap()
	right = NewEmptyJSONArray()

	performEqualTest(left, right,t,false)

	left = ParseUnsafe([]byte(testJsonString))
	right = ParseUnsafe([]byte(testJsonString))

	performEqualTest(left, right, t, true)

	right.At(3).MapSet("newKey", "newValue")

	performEqualTest(left, right, t, false)

	left = ParseUnsafe([]byte(testJsonString))
	right = ParseUnsafe([]byte(testJsonString))

	left.SliceAppend(56)

	performEqualTest(left, right, t, false)
}
func TestSliceIterators(t *testing.T) {
	originalArr := []int{1,2,3,4,5}
	counter := len(originalArr)
	jsnSlice := New(originalArr)

	jsnSlice.SliceForEach(func(jsn *JSON, index int) {
		if jsn.GetUnsafeInt() != originalArr[index] {
			t.Errorf("expected %d at index %d but instead got %d", originalArr[index], index, jsn.GetUnsafeInt())
			return
		}
		counter--
	})

	if counter != 0 {
		t.Errorf("Length didn't match original array, expected counter to be 0 but instead got %d", counter)
	}

	jsnSlice.SliceMap(func(jsn *JSON, index int) *JSON {
		return New(jsn.GetUnsafeInt() * 3)
	})

	jsnSlice.SliceForEach(func(jsn *JSON, index int) {
		if jsn.GetUnsafeInt() != (originalArr[index] * 3) {
			t.Errorf("expected %d at index %d but instead got %d", originalArr[index], index, jsn.GetUnsafeInt())
			return
		}
	})

	jsnSlice.Set(originalArr)

	jsnSlice.SliceFilter(func(jsn *JSON, index int) (shouldKeep bool) {
		return jsn.GetUnsafeInt() % 2 == 0
	})

	if !EqualsDeep(jsnSlice, New([]int{2,4})) {
		t.Errorf("Expected jsn to equal deep [2,4] but instead got %s", jsnSlice.ToUnsafeJSONString())
	}
}

func TestMapIterators(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["keyA"] = "string"
	testMap["keyB"] = 56
	testMap["keyC"] = []int{1,2,3,4,5}

	jsnMap := New(testMap)
	counter := 3

	jsnMap.ObjectForEach(func(jsn *JSON, key string) {
		if _, exist := testMap[key]; !exist {
			t.Errorf("key %s does not exist on testMap", key)
		}
		if !EqualsDeep(jsn, New(testMap[key])){
			t.Errorf("Expected %v for key %s but got %s", testMap[key], key, jsn.ToInterface())
			return
		}
		counter--
	})

	if counter != 0 {
		t.Errorf("Length didn't match original object keys size, expected counter to be 0 but instead got %d", counter)
	}

	jsnMap.ObjectFilter(func(jsn *JSON, key string) (shouldKeep bool) {
		return key != "keyA"
	})

	if jsnMap.ObjectKeyExists("keyA") {
		t.Errorf("ObjectFilter should had remove keyA but key still exists")
	}

	jsnMap.ObjectMap(func(jsn *JSON, key string) *JSON {
		if key == "keyB" {
			return New(10)
		}

		if key == "keyC" {
			return jsn.SliceAppend(6)
		}
		return jsn
	})

	shouldEqualsJson := NewEmptyJSONMap()
	shouldEqualsJson.MapSet("keyB", 10)
	shouldEqualsJson.MapSet("keyC", []int{1,2,3,4,5,6})

	if !EqualsDeep(jsnMap, shouldEqualsJson) {
		t.Errorf("Expected maps to be equal, %v, %v", jsnMap.ToInterface(), shouldEqualsJson.ToInterface())
	}
}