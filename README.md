# Jonson

Fast, lightweight thread-safe Golang JSON handler\

----

### Install

```shell
go get github.com/panthesingh/goson
```

### Quick start

#### Example 1
```go
import "github.com/KromDaniel/Jonson"


json := Jonson.NewEmptyHashJSON()
//creates {"keyA":1,"keyB":[1,2,3,4,5]}
json.HashSet("keyA", 1).HashSet("keyB", []int{1,2,3,4,5}) 
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
json.HashMapFilter(func(json *Jonson.JSON, s string) bool {
    return !json.IsInt()  //{"keyB":[{"nested":"I'm nested value"},90,"someString",15]}
})

return json.ToUnsafeJSONString()

```



