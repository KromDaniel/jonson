package lib

func (jsn *Jonson) Set(v interface{}) *Jonson {
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	temp := gosonize(v)
	jsn.kind = temp.kind
	jsn.value = temp.value

	return jsn
}

func (jsn *Jonson) HashSet(key string, value interface{}) *Jonson {
	isHashMap, hashMap := jsn.GetHashMap()
	if !isHashMap {
		return jsn
	}
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()
	hashMap[key] = gosonize(value)
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
		jsn.value = append(arr, gosonize(v))
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
		arr = append([]*Jonson{gosonize(v)}, arr...)
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

	arr[index] = gosonize(value)
	jsn.value = arr
	return jsn
}
