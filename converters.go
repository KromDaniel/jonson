package main

import (
	"encoding/json"
	"reflect"
)

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

func (jsn *Jonson) ToUnsafeJSONString() string {
	data, err := jsn.ToJSON()
	if err != nil {
		return ""
	}
	return string(data)
}

func (jsn *Jonson) ToInterface() interface{} {
	if jsn.IsPrimitive() {
		return &jsn.value
	}

	if jsn.IsSlice() {
		arr := jsn.GetUnsafeSlice()
		resArr := make([]interface{}, len(arr))
		for k, v := range arr {
			resArr[k] = v.ToInterface()
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

	return NewEmptyJSON()
}
