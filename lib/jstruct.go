package lib

import (
	"reflect"
	"sync"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"time"
)

type GoSonMap map[string]*GoSon
type GoSon struct {
	threadSafe  bool
	rwMutex     sync.RWMutex
	isPrimitive bool
	kind        reflect.Kind
	value       interface{}
}

/*
	Set thread safety operation is NOT thread safe, do it first thing when creating JSON
 */
func (gs *GoSon) SetThreadSafety(isSafe bool) {
	if gs.threadSafe == isSafe {
		return
	}
	gs.threadSafe = isSafe

}
func (gs *GoSon) SetValue(v interface{}) {
	if gs.threadSafe {
		gs.rwMutex.Lock()
		defer gs.rwMutex.Unlock()
	}

	temp := gosonize(v)
	gs.kind = temp.kind
	gs.value = temp.value
}

func (gs *GoSon) ToJSON() ([]byte, error) {
	return json.Marshal(gs.ToInterface())
}

func (gs *GoSon) ToUnsafeJson() (data []byte) {
	data, err := gs.ToJSON()
	if err != nil {
		return []byte{}
	}
	return
}

func (gs *GoSon) ToJSONString() (string, error) {
	data, err := gs.ToJSON()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (gs *GoSon) ToInterface() interface{} {
	if gs.IsPrimitive() {
		return gs.value
	}

	if gs.IsArray() {
		arr := gs.GetUnsafeArray()
		resArr := make([]interface{}, len(arr))
		for k, v := range arr {
			resArr[k] = v.ToInterface()
		}
		return resArr
	}

	if gs.IsHashMap() {
		hMap := gs.GetUnsafeHashMap()
		resMap := make(map[string]interface{})
		for k, v := range hMap {
			resMap[k] = v.ToInterface()
		}
		return resMap
	}

	return EmptyJSON()
}
func (gs *GoSon) Clone() *GoSon {
	if gs.IsPrimitive() {
		return &GoSon{
			value:      gs.value,
			kind:       gs.kind,
			threadSafe: gs.threadSafe,
		}
	}

	if gs.IsArray() {
		arr := gs.GetUnsafeArray()
		resArr := make([]*GoSon, len(arr))
		for k, v := range arr {
			resArr[k] = v.Clone()
		}
		return &GoSon{
			value:      resArr,
			kind:       reflect.Slice,
			threadSafe: gs.threadSafe,
		}
	}

	if gs.IsHashMap() {
		hMap := gs.GetUnsafeHashMap()
		resMap := make(map[string]*GoSon)
		for k, v := range hMap {
			resMap[k] = v.Clone()
		}
		return &GoSon{
			value:      resMap,
			kind:       reflect.Map,
			threadSafe: gs.threadSafe,
		}
	}

	return EmptyJSON()
}

func (gs *GoSon) IsNil() bool {
	return gs.value == nil
}
func (gs *GoSon) IsHashMap() bool {
	if gs.threadSafe {
		gs.rwMutex.RLock()
		defer gs.rwMutex.RUnlock()
	}
	if gs.IsNil() {
		return false
	}
	return gs.kind == reflect.Map
}

func (gs *GoSon) atLocked(key interface{}, keys ...interface{}) *GoSon {
	var res *GoSon = nil
	switch reflect.TypeOf(key).Kind() {
	case reflect.Int, reflect.Uint:
		isArray, arr := gs.GetArray()
		if isArray {
			index := key.(int)
			if index < len(arr) {
				res = arr[index]
			}
		}
		break
	case reflect.String:
		isObject, obj := gs.GetHashMap()
		if isObject {
			hashKey := key.(string)
			if val, ok := obj[hashKey]; ok {
				res = val
			}
		}
		break
	}
	if len(keys) > 0 && res != nil {
		return res.At(keys[0], keys[1:]...)
	}
	if res == nil {
		res = EmptyJSON()
	}
	return res
}
func (gs *GoSon) At(key interface{}, keys ...interface{}) *GoSon {
	if gs.threadSafe {
		gs.rwMutex.RLock()
		defer gs.rwMutex.RUnlock()
	}
	res := gs.atLocked(key, keys...)
	return res
}

func (gs *GoSon) IsArray() bool {
	if gs.threadSafe {
		gs.rwMutex.RLock()
		defer gs.rwMutex.RUnlock()
	}
	if gs.IsNil() {
		return false
	}
	return gs.kind == reflect.Slice
}

func (gs *GoSon) IsPrimitive() bool {
	return !(gs.IsArray() || gs.IsHashMap() || gs.IsStruct())
}

func (gs *GoSon) IsStruct() bool {
	if gs.threadSafe {
		gs.rwMutex.RLock()
		defer gs.rwMutex.RUnlock()
	}
	if gs.IsNil() {
		return false
	}

	return gs.kind == reflect.Struct
}

func (gs *GoSon) GetValue() interface{} {
	if gs.threadSafe {
		gs.rwMutex.RLock()
		defer gs.rwMutex.RUnlock()
	}

	return gs.value
}

// ==== Getter helpers ====//
func (gs *GoSon) GetInt() (isInt bool, value int) {
	if gs.threadSafe {
		gs.rwMutex.RLock()
		defer gs.rwMutex.RUnlock()
	}
	isInt = gs.kind == reflect.Int
	if isInt {
		value = gs.value.(int)
	}
	return
}

func (gs *GoSon) GetUnsafeInt() (value int) {
	_, value = gs.GetInt()
	return
}

func (gs *GoSon) GetFloat32() (isFloat32 bool, value float32) {
	if gs.threadSafe {
		gs.rwMutex.RLock()
		defer gs.rwMutex.RUnlock()
	}
	isFloat32 = gs.kind == reflect.Float32
	if isFloat32 {
		value = gs.value.(float32)
	}
	return
}

func (gs *GoSon) GetUnsafeFloat32() (value float32) {
	_, value = gs.GetFloat32()
	return
}

func (gs *GoSon) GetFloat64() (isFloat64 bool, value float64) {
	if gs.threadSafe {
		gs.rwMutex.RLock()
		defer gs.rwMutex.RUnlock()
	}
	isFloat64 = gs.kind == reflect.Float64
	if isFloat64 {
		value = gs.value.(float64)
	}
	return
}

func (gs *GoSon) GetUnsafeFloat64() (value float64) {
	_, value = gs.GetFloat64()
	return
}

func (gs *GoSon) GetBool() (isBool bool, value bool) {
	if gs.threadSafe {
		gs.rwMutex.RLock()
		defer gs.rwMutex.RUnlock()
	}
	isBool = gs.kind == reflect.Bool
	if isBool {
		value = gs.value.(bool)
	}
	return
}

func (gs *GoSon) GetUnsafeBool() (value bool) {
	_, value = gs.GetBool()
	return
}

func (gs *GoSon) GetString() (isString bool, value string) {
	if gs.threadSafe {
		gs.rwMutex.RLock()
		defer gs.rwMutex.RUnlock()
	}
	isString = gs.kind == reflect.String
	if isString {
		value = gs.value.(string)
	}
	return
}

func (gs *GoSon) GetUnsafeString() (value string) {
	_, value = gs.GetString()
	return
}

func (gs *GoSon) GetHashMap() (isHashMap bool, value GoSonMap) {
	if gs.threadSafe {
		gs.rwMutex.RLock()
		defer gs.rwMutex.RUnlock()
	}
	isHashMap = gs.kind == reflect.Map
	if isHashMap {
		value = gs.value.(map[string]*GoSon)
	}
	return
}

func (gs *GoSon) GetUnsafeHashMap() (value map[string]*GoSon) {
	isHashMap, m := gs.GetHashMap()
	if isHashMap {
		value = m
		return
	}
	value = make(map[string]*GoSon)
	return
}

func (gs *GoSon) GetArray() (isArray bool, value []*GoSon) {
	if gs.threadSafe {
		gs.rwMutex.RLock()
		defer gs.rwMutex.RUnlock()
	}
	isArray = gs.kind == reflect.Slice
	if isArray {
		value = gs.value.([]*GoSon)
	}

	return
}

func (gs *GoSon) GetUnsafeArray() (value []*GoSon) {
	isArray, m := gs.GetArray();
	if isArray {
		value = m
		return
	}
	value = make([]*GoSon, 0)
	return
}

func gosonize(value interface{}) *GoSon {
	if value == nil {
		return EmptyJSON()
	}
	vo := reflect.ValueOf(value)
	if vo.Kind() == reflect.Ptr {
		vo = vo.Elem()
	}
	switch vo.Kind() {
	case reflect.Ptr:
		return gosonize(vo.Elem())
	case reflect.Map:
		vMap := value.(map[string]interface{})
		return gosonizeMap(&vMap)
	case reflect.Slice:
		return gosonizeSlice(&vo)
	case reflect.String, reflect.Uint, reflect.Bool, reflect.Float64,
		reflect.Float32, reflect.Int, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &GoSon{
			value:       value,
			isPrimitive: true,
			kind:        vo.Kind(),
			threadSafe:  true,
		}
	case reflect.Struct:
		if v, ok := value.(GoSon); ok {
			return &v
		}
		tempMap := make(map[string]interface{})
		typ := reflect.TypeOf(value)
		for i := 0; i < typ.NumField(); i++ {
			if vo.Field(i).CanInterface() {
				fieldValue := typ.Field(i)
				if v, has := fieldValue.Tag.Lookup("json"); has {
					if v != "-" {
						tempMap[v] = vo.Field(i).Interface()
					}
					continue
				}
				tempMap[fieldValue.Name] = vo.Field(i).Interface()
			}
		}
		return gosonizeMap(&tempMap)
	}

	return EmptyJSON()
}

func gosonizeMap(value *map[string]interface{}) *GoSon {
	fmt.Println("GOSONIZE MAP")
	mapValue := make(map[string]*GoSon)
	for k, v := range *value {
		mapValue[k] = gosonize(v)
	}

	return &GoSon{
		value:       mapValue,
		isPrimitive: false,
		kind:        reflect.Map,
		threadSafe:  true,
	}
}

func gosonizeSlice(value *reflect.Value) *GoSon {
	arrValue := make([]*GoSon, value.Len())
	for i:=0; i < value.Len(); i++ {
		arrValue[i] = gosonize(value.Index(i).Interface())
	}

	return &GoSon{
		value:       arrValue,
		isPrimitive: false,
		kind:        reflect.Slice,
		threadSafe:  true,
	}
}

func New(value interface{}) *GoSon {
	return gosonize(value)
}

func EmptyJSON() *GoSon {
	return &GoSon{
		value:       nil,
		isPrimitive: true,
		threadSafe:  true,
	}
}

func EmptyHashJSON() *GoSon {
	return New(make(map[string]*GoSon))
}

func EmptyArray() *GoSon {
	return New(make(map[string]*GoSon))
}

func Parse(data []byte) (err error, goson *GoSon) {
	var m interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return
	}
	goson = New(m)
	return
}

func ParseUnsafe(data []byte) *GoSon {
	_, goson := Parse(data)
	if goson == nil {
		return EmptyJSON()
	}
	return goson
}

func Test() {
	readBegin := time.Now()
	b, err := ioutil.ReadFile("/Users/danielkrom/Dev/sf-city-lots-json/smaller.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var m interface{}
	json.Unmarshal([]byte(string(b)), &m)
	readEnd := time.Now().Sub(readBegin).Seconds()
	fmt.Println("Read and parse", readEnd)
	for i :=0; i < 5; i++{
		createBegin := time.Now()
		t := New(b)
		fmt.Println("Create", time.Now().Sub(createBegin).Seconds())
		fmt.Println(t.threadSafe)
		fmt.Println(t.At("features"))
	}

}
