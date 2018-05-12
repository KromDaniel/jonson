# Jonson

Fast, lightweight, thread-safe, schemaless Golang JSON handler

##### Beta version
----



## Table of Contents

1. [Quick start](#install)
2. [Getters](#getters)
3. [Setters](#setters)
4. [Constructors](#constructors)
5. [Types](#types)
6. [Convertors](#convertors)
7. [Iterators](#iterators)
8. [Threads](#threads)
9. [License](#license)
10. [Contact](#contact)
 
## Install

```shell
go get github.com/KromDaniel/jonson
```

## Quick start


##### Parsing and working with JSON

```go
import "github.com/KromDaniel/jonson"
```

```go
err, json := jonson.Parse([]byte(`{"foo": "bar", "arr": [1,2,"str", {"nestedFooA" : "nestedBar"}]}`))
if err != nil {
    // error handler
}
// Array mapper
json.At("arr").SliceMap(func(jsn *jonson.JSON, index int) *jonson.JSON {
    // JSON numbers are always float when parsed
    if jsn.IsFloat64() {
        return jonson.New(jsn.GetUnsafeFloat64() * float64(4))
    }
    if jsn.IsString() {
        return jonson.New("_" + jsn.GetUnsafeString())
    }

    if jsn.IsMap() {
        jsn.MapSet("me", []int{1, 2, 3})
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

## Getters

Getters are the way to retreive the actual value of the JSON,
Since jonson is thread safe, the value is cloned before returned, unless using `At` method

#### IsType
Jonson supports most of the reflect types
each Jonson object can be asked for `IsType(t reflect.Kind)` or directly e.g `IsInt()`.

##### Example
```go
json := jonson.New("str")
json.IsString() // true
json.IsSlice() // false
json.IsInt() // false
```

A legal JSON value can be one of the following types:

* string
* number
* object
* array
* boolean
* null

Jonson supports the getters `IsSlice` `IsMap` and special `IsPrimitive` for string, number, boolean and null. 

Since there are many type of numbers, There's a getter for each type  e.g `IsUint8` `IsFloat32`

**Note** When parsing JSON string, the default value of each number is `Float64`


#### Value type getters

Each of reflect.Kind type has a getter and unsafe getter, unsafe getter returns the zero value for that type if type is wrong

##### Example
```go
json := jonson.New(96)
isInt, val := json.GetInt()
if isInt {
    // safe, it's int
}

json.GetUnsafeFloat64() //0 value
```

#### Methods
* `JSON.At(keys ...interface{})` returns a pointer to the current key, can be chained to null values
* `JSON.GetSlice()` returns  `[]*jonson.JSON`
* `JSON.GetMap()` returns `map[string]*jonson.JSON`
* `JSON.GetObjectKeys()` returns `[]string` if JSON is map else nil
* `JSON.GetSliceLen()` returns `int`, the length of the slice if JSON is slice, else 0

#### At Method
`JSON.At` method accepts as argument `int` or `string`</br>
if passed a string and the JSON is not map, returns the zero JSON </br>
if passed int and the JSON is not slice, returns the zero JSON
It can be chained even to none existing value

#### At Example

```go
js := jonson.NewEmptyJSONMap()
js.At("KeyOfObjectWithArrayAsValue").At(12).At(54).At("key") // 12, 54 is index of slice, strin gis key of map
// Same as
js.At("KeyOfObjectWithArrayAsValue", 12, 54, "key")
// same as
js.At("KeyOfObjectWithArrayAsValue", 12).At(54, "key") 
// same as goes on...
```


## Setters

Setters are used to set the value of the current JSON pointer

#### How set works
Since jonson is thread safe, it must be aware when trying to read or write a value, in order
to gurantee that, value is deeply cloned, if value passed as pointer, the jonson will use the actual element it points to
(For better performance it is usually better to pass value as pointer, so the deep clone will happen only once at the jonson cloner)

For performance, use the setters provided by the library that can safely mutate each value and able to directly access it's pointer

#### Methods

* `JSON.SetValue(v interface{})` sets the passed value to current JSON pointer, overrides the type and the existing value
* `JSON.MapSet(key string, v interface{})` sets value to current JSON as the current key (works only if current JSON is map type)
* `JSON.SliceAppend(v ...interface{})` append all given values to slice (works only if current JSON is slice type)
* `JSON.SliceAppendBegin(v ...interface{})` same as `SliceAppend` but at the start of the slice instead at the end
* `JSON.SliceSet(index int, v interace{})` overrides value at specific index on slice (works only if current JSON is slice type)

##### Example

```go
json := jonson.NewEmptyJSON() // nil value
exampleMap := make(map[string]int)
exampleMap["1"] = 1
exampleMap["2"] = 2

json.Set(&exampleMap)
exampleMap["1"] = 4

// key 1 is different value, because setters do deep clone
fmt.Println(exampleMap) // map[1:4 2:2]
fmt.Println(json.ToUnsafeJSONString()) // {"1":1,"2":2}
```

### Constructors

Constructors are the way to initialize a new JSON object

##### Methods

* `jonson.New(value interface{}) *JSON` creates a new JSON containing the passed value
* `jonson.NewEmptyJSON() *JSON` creates a new empty JSON with the value of nil
* `jonson.NewEmptyJSONMap() *JSON` creates a new empty JSON with the value `map[string]*JSON`
* `jonson.NewEmptyJSONArray() *JSON` creates a new empty JSON with the value 0 length slice
* `jonson.Parse([]byte) (error, *JSON)` parses the byte (assumed to be UTF-8 JSON string)
* `jonson.ParseUnsafe([]byte) *JSON` same as `jonson.Parse` but returns the `jonson.NewEmptyJSON()` if error

## Types

Jonson supports all valid types for JSON, here's how it works:

##### Map

JSON Object (key, value) is valid only for strings key, it means that only `map[string]interface{}` will work, a map with none string keys, the key will be ignored

###### Example
```go
keyMixedMap := make(map[interface{}]interface{})
keyMixedMap[1] = "key is integer"
keyMixedMap["key"] = "key is string"

fmt.Println(jonson.New(&keyMixedMap).ToUnsafeJSONString()) //{"key":"key is string"}
```

##### Struct

Struct behaves the same as with `encoding/json`

Only public fields are exported, the name of the field is the key on the struct, unless there's a field descriptors with json tag `json:"customKey"`.
If key is public and tagged with `json:"-"` it is ignored.

**Note** When passing a struct to Jonson, it is immediately being "Jonsonized" means the keys are converted instantly

###### Example
```go
type MyStruct struct {
    Public  string
    private string
    Custom  string `json:"customKey"`
    Ignored string `json:"-"`
}

structExample := jonson.New(&MyStruct{
    Public:  "public value",
    private: "private value",
    Custom:  "custom value",
    Ignored: "Ignored value",
})

fmt.Println(structExample.At("private").IsNil()) // true
fmt.Println(structExample.ToUnsafeJSONString())  // {"Public":"public value","customKey":"custom value"}
```

##### Slice

Slice is the array type of JSON, jonson supports all kind of slices, as long as each element is JSON legal

## Convertors

Convertors is a group of methods that converts the JSON object without changing it

##### Methods

* `JSON.ToJSON() ([]byte, error)` stringify the JSON to `[]byte`
* `JSON.ToUnsafeJSON() []byte` stringify the JSON, if error returns empty `[]byte`
* `JSON.ToJSONString() (string, error)`
* `JSON.ToUnsafeJSONString() string` empty string if error
* `JSON.ToInterface() interface{}` returns the entire JSON tree as interface
* `JSON.Clone() *JSON` Deep clone the current JSON tree


## Iterators

Iterators is a group of methods that allows iteration on slice or map, it accepts a function as argument for callback

The methods will do nothing is JSON is not slice or map (according the relevant method)

**Note** Using map or filter (Similar to other languages Array.map and Array.filter), It won't return a new copy of the slice or map, it will **mutate the existing one**

##### Methods
*  `JSON.SliceForEach(cb func(jsn *JSON, index int))` Iterate on JSON slice
*  `JSON.SliceMap(cb func(jsn *JSON, index int) *JSON)` Iterate on JSON slice, replacing each element with returned JSON
*  `JSON.SliceFilter(cb func(jsn *JSON, index int) bool)` Iterate on JSON slice, removing element if cb returned false 
*  `JSON.ObjectForEach(cb func(jsn *JSON, key string))` Iterate on JSON map 
*  `JSON.ObjectMap(cb func(jsn *JSON, key string) *JSON)` Iterate on JSON map, replacing each value with returned JSON
*  `JSON.ObjectFilter(cb func(jsn *JSON, key string) bool)` Iterate on JSON map, removing value if cb returned false 

##### Example

```go
jsn := jonson.NewEmptyJSONMap()

jsn.MapSet("keyA", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
jsn.MapSet("KeyB", 1)
jsn.MapSet("KeyC", 2)
fmt.Println(jsn.ToUnsafeJSONString()) // {"KeyB":1,"KeyC":2,"keyA":[1,2,3,4,5,6,7,8,9,10]}

// Object Map, multiply integer values * 3
jsn.ObjectMap(func(jsn *jonson.JSON, key string) *jonson.JSON {
    if jsn.IsInt() {
        return jonson.New(jsn.GetUnsafeInt() * 3)
    }
    return jsn
    // iterate on the array, keep only evens
}).At("keyA").SliceFilter(func(jsn *jonson.JSON, index int) (shouldKeep bool) {
    shouldKeep = jsn.GetUnsafeInt()%2 == 0
    return
})
fmt.Println(jsn.ToUnsafeJSONString()) // {"KeyB":3,"KeyC":6,"keyA":[2,4,6,8,10]}
```
## Threads

Jonson managed thread safety by it self, it using read-writer mutex `sync.RWMutex`
allowing multple readers the same time

#### Example
```go
func writer(jsn *jonson.JSON, wg *sync.WaitGroup) {
	for i :=0 ; i < 100000; i++ {
		jsn.SliceAppend(i)
	}
	wg.Done()
}

func reader(jsn *jonson.JSON, wg *sync.WaitGroup) {
	time.Sleep(time.Nanosecond * 1000)
	fmt.Println("Reader", jsn.GetSliceLen())
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	arr := jonson.NewEmptyJSONArray()
	wg.Add(5)
	go writer(arr, &wg)
	for i:=0; i < 4; i++{
		go reader(arr, &wg)
	}
	wg.Wait()
	fmt.Println("Final len", arr.GetSliceLen())
}

/* Output
Reader 5650
Reader 5651
Reader 5652
Reader 5651
Final len 100000
*/
```

## License
Apache 2.0

## Contact
For any question or contribution, feel free to contact me at
kromdan@gmail.com