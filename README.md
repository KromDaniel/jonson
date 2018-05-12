# Jonson

Fast, lightweight, thread-safe, schemaless Golang JSON handler

----

### Install

```shell
go get github.com/KromDaniel/jonson
```

### Quick start


##### Parsing and working with JSON

```go
import "github.com/KromDaniel/jonson"


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

```

##### Creating JSON

```go
json := jonson.NewEmptyJSONMap()
json.MapSet("arr", []interface{}{1, "str", []uint16{50,60,70}})
json.MapSet("numbers", []interface{}{})

for i:=0; i < 100; i++ {
    json.At("numbers").SliceAppend(i)
}

json.At("numbers").SliceFilter(func(jsn *jonson.JSON, index int) (shouldKeep bool) {
    return IsPrime(jsn.GetUnsafeInt())
})

// {"arr":[1,"str",[50,60,70]],"numbers":[2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97]}
fmt.Println(json.ToUnsafeJSONString())
```