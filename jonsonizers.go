package Jonson

import "reflect"

func jonsonize(value interface{}) *Jonson {
	if value == nil {
		return NewEmptyJSON()
	}
	vo := reflect.ValueOf(value)
	if vo.Kind() == reflect.Ptr {
		vo = vo.Elem()
	}
	switch vo.Kind() {
	case reflect.Ptr:
		return jonsonize(vo.Elem())
	case reflect.Map:
		return jonsonizeMap(&vo)
	case reflect.Slice:
		return jonsonizeSlice(&vo)
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
			return v.Clone()
		}
		return jonsonizeStruct(&vo)
	}

	return NewEmptyJSON()
}

func jonsonizeMap(value *reflect.Value) *Jonson {
	mapValue := make(JonsonMap)
	for _, k := range value.MapKeys() {
		// map should be only string as keys
		keyValue := reflect.ValueOf(k)
		if keyValue.Kind() != reflect.String{
			continue
		}
		mapValue[keyValue.String()] = jonsonize(keyValue.MapIndex(k))
	}

	return &Jonson{
		value:       mapValue,
		isPrimitive: false,
		kind:        reflect.Map,
	}
}

func jonsonizeSlice(value *reflect.Value) *Jonson {
	arrValue := make([]*Jonson, value.Len())
	for i := 0; i < value.Len(); i++ {
		arrValue[i] = jonsonize(value.Index(i).Interface())
	}

	return &Jonson{
		value:       arrValue,
		isPrimitive: false,
		kind:        reflect.Slice,
	}
}

func jonsonizeStruct(vo *reflect.Value) *Jonson{
	tempMap := make(map[string]interface{})
	typ := vo.Type()
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
	return jonsonize(&tempMap)
}