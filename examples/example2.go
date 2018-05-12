/*
	Example 1

*/

package main

import (
	"fmt"
	"math"

	"github.com/KromDaniel/jonson"
)



func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func main() {

	json := jonson.NewEmptyJSONMap()

	json.MapSet("arr", []interface{}{1, "str", []uint16{50, 60, 70}})
	json.MapSet("numbers", []interface{}{})

	for i := 0; i < 100; i++ {
		json.At("numbers").SliceAppend(i)
	}

	json.At("numbers").SliceFilter(func(jsn *jonson.JSON, index int) (shouldKeep bool) {
		return IsPrime(jsn.GetUnsafeInt())
	})

	// {"arr":[1,"str",[50,60,70]],"numbers":[2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97]}
	fmt.Println(json.ToUnsafeJSONString())

}
