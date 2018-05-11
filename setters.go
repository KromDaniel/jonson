package Jonson

func (jsn *JSON) Set(v interface{}) *JSON {
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	temp := jonsonize(v)
	jsn.kind = temp.kind
	jsn.value = temp.value
	jsn.isPrimitive = temp.isPrimitive
	return jsn
}

func (jsn *JSON) HashSet(key string, value interface{}) *JSON {
	isHashMap, hashMap := jsn.GetHashMap()
	if !isHashMap {
		return jsn
	}
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()
	hashMap[key] = jonsonize(value)
	jsn.value = hashMap

	return jsn
}

func (jsn *JSON) SliceAppend(value ...interface{}) *JSON {
	isSlice, arr := jsn.GetSlice()
	if !isSlice {
		return jsn
	}
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()
	for _, v := range value {
		jsn.value = append(arr, jonsonize(v))
	}

	return jsn
}

func (jsn *JSON) SliceAppendBegin(value ...interface{}) *JSON {
	isSlice, arr := jsn.GetSlice()
	if !isSlice {
		return jsn
	}
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()
	for _, v := range value {
		arr = append([]*JSON{jonsonize(v)}, arr...)
	}

	jsn.value = arr

	return jsn
}

func (jsn *JSON) SliceSet(index int, value interface{}) *JSON {
	isSlice, arr := jsn.GetSlice()
	if !isSlice {
		return jsn
	}
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	arr[index] = jonsonize(value)
	jsn.value = arr
	return jsn
}
