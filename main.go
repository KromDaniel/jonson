package main

import (
	"fmt"
)

func main() {
	json := NewEmptyHashJSON()
	json.HashSet("Yoho", "65").HashSet("me", 432).HashSet("arr", []int{1, 2, 3, 4, 5})
	for i:= 0; i < 1000; i++{
		json.At("arr").SliceAppend(i)
	}
	json.At("arr").SliceMap(func(jonson *Jonson, i int) interface{} {
		return jonson.GetUnsafeInt() * 2
	}).SliceFilter(func(jonson *Jonson, i int) bool {
		return jonson.GetUnsafeInt() % 7 == 0
	}).SliceAppendBegin("dsa", 321321, []string{"A", "B", "C"})
	str, _ := json.ToJSONString()
	fmt.Println(str)
}
