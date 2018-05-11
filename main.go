package main

import (
	"github.com/KromDaniel/GoSon/lib"
	"fmt"
)

func main() {
	json := lib.EmptyHashJSON()
	json.HashSet("Yoho", "65")

	str, _ := json.ToJSONString()
	fmt.Println(str)
}
