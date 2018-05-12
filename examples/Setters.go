/*
	Setters example

 */

package main

import (
	"github.com/KromDaniel/jonson"
	"fmt"
)

func main() {

	 json := jonson.NewEmptyJSON() // nil value
	 exampleMap := make(map[string]int)
	 exampleMap["1"] = 1
	 exampleMap["2"] = 2

	 json.Set(&exampleMap)
	 exampleMap["1"] = 4

	 // key 1 is different, because setters does deep clone
	 fmt.Println(exampleMap) // map[1:4 2:2]
	 fmt.Println(json.ToUnsafeJSONString()) // {"1":1,"2":2}
}
