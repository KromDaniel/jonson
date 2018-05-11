package main

func (jsn *Jonson) Set(v interface{}) *Jonson {
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	temp := jonsonize(v)
	jsn.kind = temp.kind
	jsn.value = temp.value
	jsn.isPrimitive = temp.isPrimitive
	return jsn
}

func (jsn *Jonson) HashSet(key string, value interface{}) *Jonson {
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

func (jsn *Jonson) SliceAppend(value ...interface{}) *Jonson {
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

func (jsn *Jonson) SliceAppendBegin(value ...interface{}) *Jonson {
	isSlice, arr := jsn.GetSlice()
	if !isSlice {
		return jsn
	}
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()
	for _, v := range value {
		arr = append([]*Jonson{jonsonize(v)}, arr...)
	}

	jsn.value = arr

	return jsn
}

func (jsn *Jonson) SliceSet(index int, value interface{}) *Jonson {
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
