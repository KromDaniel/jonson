package lib

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type JonsonMap map[string]*Jonson

type Jonson struct {
	rwMutex     sync.RWMutex
	isPrimitive bool
	kind        reflect.Kind
	value       interface{}
}

/*
	Set thread safety operation is NOT thread safe, do it first thing when creating JSON
*/

func (jsn *Jonson) SetValue(v interface{}) {
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	temp := gosonize(v)
	jsn.kind = temp.kind
	jsn.value = temp.value
}

func (jsn *Jonson) ToJSON() ([]byte, error) {
	return json.Marshal(jsn.ToInterface())
}

func (jsn *Jonson) ToUnsafeJson() (data []byte) {
	data, err := jsn.ToJSON()
	if err != nil {
		return []byte{}
	}
	return
}

func (jsn *Jonson) ToJSONString() (string, error) {
	data, err := jsn.ToJSON()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (jsn *Jonson) ToInterface() interface{} {
	if jsn.IsPrimitive() {
		return &jsn.value
	}

	if jsn.IsSlice() {
		begin := time.Now()
		arr := jsn.GetUnsafeSlice()
		unsafeSlice := time.Now().Sub(begin).Seconds()
		begin = time.Now()
		resArr := make([]interface{}, len(arr))
		for k, v := range arr {
			resArr[k] = v.ToInterface()
		}
		loopEnd := time.Now().Sub(begin).Seconds()
		if len(arr) > 5000 {
			fmt.Println("Unsafe Slice ", unsafeSlice, " loopEnd ", loopEnd, "len", len(arr))
		}
		return resArr
	}

	if jsn.IsHashMap() {
		hMap := jsn.GetUnsafeHashMap()
		resMap := make(map[string]interface{})
		for k, v := range hMap {
			resMap[k] = v.ToInterface()
		}
		return &resMap
	}

	return nil
}
func (jsn *Jonson) Clone() *Jonson {
	if jsn.IsPrimitive() {
		return &Jonson{
			value: jsn.value,
			kind:  jsn.kind,
		}
	}

	if jsn.IsSlice() {
		arr := jsn.GetUnsafeSlice()
		resArr := make([]*Jonson, len(arr))
		for k, v := range arr {
			resArr[k] = v.Clone()
		}
		return &Jonson{
			value: resArr,
			kind:  reflect.Slice,
		}
	}

	if jsn.IsHashMap() {
		hMap := jsn.GetUnsafeHashMap()
		resMap := make(map[string]*Jonson)
		for k, v := range hMap {
			resMap[k] = v.Clone()
		}
		return &Jonson{
			value: resMap,
			kind:  reflect.Map,
		}
	}

	return EmptyJSON()
}

func (jsn *Jonson) IsNil() bool {
	return jsn.value == nil
}
func (jsn *Jonson) IsHashMap() bool {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	if jsn.IsNil() {
		return false
	}
	return jsn.kind == reflect.Map
}

func (jsn *Jonson) atLocked(key interface{}, keys ...interface{}) *Jonson {
	var res *Jonson = nil
	switch reflect.TypeOf(key).Kind() {
	case reflect.Int, reflect.Uint:
		isSlice, arr := jsn.GetSlice()
		if isSlice {
			index := key.(int)
			if index < len(arr) {
				res = arr[index]
			}
		}
		break
	case reflect.String:
		isObject, obj := jsn.GetHashMap()
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
func (jsn *Jonson) At(key interface{}, keys ...interface{}) *Jonson {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	res := jsn.atLocked(key, keys...)
	return res
}

func (jsn *Jonson) IsSlice() bool {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	if jsn.IsNil() {
		return false
	}
	return jsn.kind == reflect.Slice
}

func (jsn *Jonson) IsPrimitive() bool {
	return !(jsn.IsSlice() || jsn.IsHashMap())
}

func (jsn *Jonson) GetValue() interface{} {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	return jsn.value
}

func gosonize(value interface{}) *Jonson {
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
	case reflect.String,
		reflect.Bool,
		reflect.Float64,
		reflect.Float32,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return &Jonson{
			value:       value,
			isPrimitive: true,
			kind:        vo.Kind(),
		}
	case reflect.Struct:
		if v, ok := value.(Jonson); ok {
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

func gosonizeMap(value *map[string]interface{}) *Jonson {
	mapValue := make(map[string]*Jonson)
	for k, v := range *value {
		mapValue[k] = gosonize(v)
	}

	return &Jonson{
		value:       mapValue,
		isPrimitive: false,
		kind:        reflect.Map,
	}
}

func gosonizeSlice(value *reflect.Value) *Jonson {
	arrValue := make([]*Jonson, value.Len())
	for i := 0; i < value.Len(); i++ {
		arrValue[i] = gosonize(value.Index(i).Interface())
	}

	return &Jonson{
		value:       arrValue,
		isPrimitive: false,
		kind:        reflect.Slice,
	}
}

func New(value interface{}) *Jonson {
	return gosonize(value)
}

func EmptyJSON() *Jonson {
	return &Jonson{
		value:       nil,
		isPrimitive: true,
	}
}

func EmptyHashJSON() *Jonson {
	return New(make(map[string]*Jonson))
}

func EmptySlice() *Jonson {
	return New(make(map[string]*Jonson))
}

func Parse(data []byte) (err error, goson *Jonson) {
	var m interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return
	}
	goson = New(m)
	return
}

func ParseUnsafe(data []byte) *Jonson {
	_, goson := Parse(data)
	if goson == nil {
		return EmptyJSON()
	}
	return goson
}