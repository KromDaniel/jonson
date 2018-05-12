/*
	Example 1

*/

package main

import (
	"fmt"

	"github.com/KromDaniel/jonson"
)

func main() {
	err, json := jonson.Parse([]byte(`{"foo": "bar", "arr": [1,2,"str", {"nestedFooA" : "nestedBar"}]}`))
	if err != nil {
		// error handler
	}
	// jsn is a clone of the original JSON object, since it's immutable
	json.At("arr").SliceMap(func(jsn *jonson.JSON, index int) *jonson.JSON {
		// JSON numbers are always float when parsed
		if jsn.IsFloat64() {
			return jonson.New(jsn.GetUnsafeFloat64() * float64(4))
		}
		if jsn.IsString() {
			return jonson.New("_" + jsn.GetUnsafeString())
		}

		if jsn.IsMap() {
			jsn.MapSet("me", []int{1, 2, 3})
		}
		return jsn
	})
	// {"arr":[4,8,"_str",{"me":[1,2,3],"nestedFooA":"nestedBar"}],"foo":"bar"}
	fmt.Println(json.ToUnsafeJSONString())

}
