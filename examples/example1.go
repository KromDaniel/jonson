/*
	Example 1

 */

package main

import (
	"github.com/KromDaniel/Jonson"
	"fmt"
)

func main(){
	json := Jonson.NewEmptyHashJSON()
	json.HashSet("keyA", 1).HashSet("keyB", []int{1,2,3,4,5}) // {"keyA":1,"keyB":[1,2,3,4,5]}
	// map the array and filter it
	json.At("keyB").SliceMap(func(jonson *Jonson.JSON, i int) interface{} {
		return jonson.GetUnsafeInt() * 3 // {"keyA":1,"keyB":[3,6,9,12,15]}
	}).SliceFilter(func(jonson *Jonson.JSON, i int) bool {
		return jonson.GetUnsafeInt() % 5 == 0 // {"keyA":1,"keyB":[15]}
	})

	someMap := make(map[string]interface{})
	someMap["nested"] = "I'm nested value"
	json.At("keyB").SliceAppendBegin("someString", 90,someMap)

	fmt.Println(json.ToUnsafeJSONString())
}