package main

import (
	"github.com/KromDaniel/jonson"
	"fmt"
)

func main(){
	js := jonson.New([]interface{}{55.6, 70.8, 10.4, 1, "48", "-90"})

	js.SliceMap(func(jsn *jonson.JSON, index int) *jonson.JSON {
		jsn.MutateToInt()
		return jsn
	}).SliceMap(func(jsn *jonson.JSON, index int) *jonson.JSON {
		if jsn.GetUnsafeInt() > 50{
			jsn.MutateToString()
		}
		return jsn
	})
	fmt.Println(js.ToUnsafeJSONString()) // ["55","70",10,1,48,-90]

}
