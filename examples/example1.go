/*
	Example 1

 */

package main

import (
	"github.com/KromDaniel/jonson"
	"fmt"
)

func main(){
	json := Jonson.NewEmptyJSONObject()
	json.MapSet("keyA", 1).MapSet("keyB", []int{1,2,3,4,5}) // {"keyA":1,"keyB":[1,2,3,4,5]}
	// map the array and filter it
	json.At("keyB").SliceMap(func(jonson *Jonson.JSON, i int) interface{} {
		return jonson.GetUnsafeInt() * 3 // {"keyA":1,"keyB":[3,6,9,12,15]}
	}).SliceFilter(func(jonson *Jonson.JSON, i int) bool {
		return jonson.GetUnsafeInt() % 5 == 0 // {"keyA":1,"keyB":[15]}
	})

	someMap := make(map[string]interface{})
	someMap["nested"] = "I'm nested value"
	json.At("keyB").SliceAppendBegin("someString", 90,someMap) //{"keyA":1,"keyB":[{"nested":"I'm nested value"},90,"someString",15]}

	// let's remove keyA
	json.ObjectFilter(func(json *Jonson.JSON, s string) bool {
		return !json.IsInt()  //{"keyB":[{"nested":"I'm nested value"},90,"someString",15]}
	})

	fmt.Println(json.ToUnsafeJSONString())

}