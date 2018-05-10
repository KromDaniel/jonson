package lib

import "reflect"

// ==== Getter helpers ====//

func (jsn *Jonson) IsType(p reflect.Kind) bool{
	return jsn.kind == p
}

func (jsn *Jonson) IsInt() bool{
	return jsn.IsType(reflect.Int)
}

func (jsn *Jonson) IsInt8() bool{
	return jsn.IsType(reflect.Int8)
}

func (jsn *Jonson) IsInt16() bool{
	return jsn.IsType(reflect.Int16)
}

func (jsn *Jonson) IsInt32() bool{
	return jsn.IsType(reflect.Int32)
}

func (jsn *Jonson) IsInt64() bool{
	return jsn.IsType(reflect.Int64)
}

func (jsn *Jonson) GetInt() (isInt bool, value int) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isInt = jsn.kind == reflect.Int
	if isInt {
		value = jsn.value.(int)
	}
	return
}

func (jsn *Jonson) GetUnsafeInt() (value int) {
	_, value = jsn.GetInt()
	return
}

func (jsn *Jonson) GetInt8() (isInt8 bool, value int8) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isInt8 = jsn.kind == reflect.Int8
	if isInt8 {
		value = jsn.value.(int8)
	}
	return
}

func (jsn *Jonson) GetUnsafeInt8() (value int8) {
	_, value = jsn.GetInt8()
	return
}

func (jsn *Jonson) GetInt16() (isInt16 bool, value int16) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isInt16 = jsn.kind == reflect.Int16
	if isInt16 {
		value = jsn.value.(int16)
	}
	return
}

func (jsn *Jonson) GetUnsafeInt16() (value int16) {
	_, value = jsn.GetInt16()
	return
}

func (jsn *Jonson) GetInt32() (isInt32 bool, value int32) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isInt32 = jsn.kind == reflect.Int32
	if isInt32 {
		value = jsn.value.(int32)
	}
	return
}

func (jsn *Jonson) GetUnsafeInt32() (value int32) {
	_, value = jsn.GetInt32()
	return
}

func (jsn *Jonson) GetInt64() (isInt64 bool, value int64) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isInt64 = jsn.kind == reflect.Int64
	if isInt64 {
		value = jsn.value.(int64)
	}
	return
}

func (jsn *Jonson) GetUnsafeInt64() (value int64) {
	_, value = jsn.GetInt64()
	return
}

func (jsn *Jonson) GetFloat32() (isFloat32 bool, value float32) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isFloat32 = jsn.kind == reflect.Float32
	if isFloat32 {
		value = jsn.value.(float32)
	}
	return
}

func (jsn *Jonson) GetUnsafeFloat32() (value float32) {
	_, value = jsn.GetFloat32()
	return
}

func (jsn *Jonson) GetFloat64() (isFloat64 bool, value float64) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isFloat64 = jsn.kind == reflect.Float64
	if isFloat64 {
		value = jsn.value.(float64)
	}
	return
}

func (jsn *Jonson) GetUnsafeFloat64() (value float64) {
	_, value = jsn.GetFloat64()
	return
}

func (jsn *Jonson) GetBool() (isBool bool, value bool) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isBool = jsn.kind == reflect.Bool
	if isBool {
		value = jsn.value.(bool)
	}
	return
}

func (jsn *Jonson) GetUnsafeBool() (value bool) {
	_, value = jsn.GetBool()
	return
}

func (jsn *Jonson) GetString() (isString bool, value string) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isString = jsn.kind == reflect.String
	if isString {
		value = jsn.value.(string)
	}
	return
}

func (jsn *Jonson) GetUnsafeString() (value string) {
	_, value = jsn.GetString()
	return
}

func (jsn *Jonson) GetHashMap() (isHashMap bool, value JonsonMap) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isHashMap = jsn.kind == reflect.Map
	if isHashMap {
		value = jsn.value.(map[string]*Jonson)
	}
	return
}

func (jsn *Jonson) GetUnsafeHashMap() (value map[string]*Jonson) {
	isHashMap, m := jsn.GetHashMap()
	if isHashMap {
		value = m
		return
	}
	value = make(map[string]*Jonson)
	return
}

func (jsn *Jonson) GetSlice() (isSlice bool, value []*Jonson) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isSlice = jsn.kind == reflect.Slice
	if isSlice {
		value = jsn.value.([]*Jonson)
	}

	return
}

func (jsn *Jonson) GetUnsafeSlice() (value []*Jonson) {
	isSlice, m := jsn.GetSlice()
	if isSlice {
		value = m
		return
	}
	value = make([]*Jonson, 0)
	return
}

func (jsn *Jonson) GetUint() (isUint bool, value uint) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isUint = jsn.kind == reflect.Uint
	if isUint {
		value = jsn.value.(uint)
	}
	return
}

func (jsn *Jonson) GetUnsafeUint() (value uint) {
	_, value = jsn.GetUint()
	return
}

func (jsn *Jonson) GetUint8() (isUint8 bool, value uint8) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isUint8 = jsn.kind == reflect.Uint8
	if isUint8 {
		value = jsn.value.(uint8)
	}
	return
}

func (jsn *Jonson) GetUnsafeUint8() (value uint8) {
	_, value = jsn.GetUint8()
	return
}

func (jsn *Jonson) GetUint16() (isUint16 bool, value uint16) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isUint16 = jsn.kind == reflect.Uint16
	if isUint16 {
		value = jsn.value.(uint16)
	}
	return
}

func (jsn *Jonson) GetUnsafeUint16() (value uint16) {
	_, value = jsn.GetUint16()
	return
}

func (jsn *Jonson) GetUint32() (isUint32 bool, value uint32) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isUint32 = jsn.kind == reflect.Uint32
	if isUint32 {
		value = jsn.value.(uint32)
	}
	return
}

func (jsn *Jonson) GetUnsafeUint32() (value uint32) {
	_, value = jsn.GetUint32()
	return
}

func (jsn *Jonson) GetUint64() (isUint64 bool, value uint64) {
	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	isUint64 = jsn.kind == reflect.Int64
	if isUint64 {
		value = jsn.value.(uint64)
	}
	return
}

func (jsn *Jonson) GetUnsafeUint64() (value uint64) {
	_, value = jsn.GetUint64()
	return
}