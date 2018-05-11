package Jonson

import (
	"encoding/json"
	"reflect"
)

func (jsn *JSON) ToJSON() ([]byte, error) {
	return json.Marshal(jsn.ToInterface())
}

func (jsn *JSON) ToUnsafeJson() (data []byte) {
	data, err := jsn.ToJSON()
	if err != nil {
		return []byte{}
	}
	return
}

func (jsn *JSON) ToJSONString() (string, error) {
	data, err := jsn.ToJSON()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (jsn *JSON) ToUnsafeJSONString() string {
	data, err := jsn.ToJSON()
	if err != nil {
		return ""
	}
	return string(data)
}

func (jsn *JSON) ToInterface() interface{} {
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
func (jsn *JSON) Clone() *JSON {
	if jsn.IsPrimitive() {
		return &JSON{
			value: jsn.value,
			kind:  jsn.kind,
		}
	}

	if jsn.IsSlice() {
		arr := jsn.GetUnsafeSlice()
		resArr := make([]*JSON, len(arr))
		for k, v := range arr {
			resArr[k] = v.Clone()
		}
		return &JSON{
			value: resArr,
			kind:  reflect.Slice,
		}
	}

	if jsn.IsHashMap() {
		hMap := jsn.GetUnsafeHashMap()
		resMap := make(map[string]*JSON)
		for k, v := range hMap {
			resMap[k] = v.Clone()
		}
		return &JSON{
			value: resMap,
			kind:  reflect.Map,
		}
	}

	return NewEmptyJSON()
}
