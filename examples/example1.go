/*
	Example 1

 */

package main

import (
	"github.com/KromDaniel/jonson"
	"fmt"
)

func main() {

	err, json := jonson.Parse([]byte(`{"foo": "bar", "arr": [1,2,"str", {"nestedFooA" : "nestedBar"}]}`))
	if err != nil {
		// error handler
	}

	json.At("arr").SliceMap(func(jsn *jonson.JSON, index int) *jonson.JSON {
		// JSON numbers are always float when parsed
		if jsn.IsFloat64() {
			return jonson.New(jsn.GetUnsafeFloat64() * float64(4))
		}
		if jsn.IsString() {
			return jonson.New("_" + jsn.GetUnsafeString())
		}

		if jsn.IsMap() {
			jsObject := jsn.GetUnsafeMap()
			jsObject["me"] = jonson.New([]int{1, 2, 3})
		}
		return jsn
	})
	// {"arr":[4,8,"_str",{"me":[1,2,3],"nestedFooA":"nestedBar"}],"foo":"bar"}
	fmt.Println(json.ToUnsafeJSONString())

}
