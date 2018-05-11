package Jonson

func (jsn *JSON) SliceForEach(cb func(*JSON, int)) *JSON {
	isSlice, slice := jsn.GetSlice()

	if !isSlice {
		return jsn
	}

	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	for i, v := range slice {
		cb(v, i)
	}
	return jsn
}

func (jsn *JSON) SliceMap(cb func(*JSON, int) interface{})*JSON {
	isSlice, slice := jsn.GetSlice()

	if !isSlice {
		return jsn
	}
	jsn.rwMutex.RLock()
	mappedArr := make([]*JSON, len(slice))
	for i, v := range slice {
		mappedArr[i] = jonsonize(cb(v, i))
	}
	jsn.rwMutex.RUnlock()
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	jsn.value = mappedArr
	return jsn
}

func (jsn *JSON) SliceFilter(cb func(*JSON, int) bool)*JSON {
	isSlice, slice := jsn.GetSlice()

	if !isSlice {
		return jsn
	}
	jsn.rwMutex.RLock()
	filteredArr := make([]*JSON, 0)
	for i, v := range slice {
		if cb(v, i){
			filteredArr = append(filteredArr, slice[i])
		}
	}
	jsn.rwMutex.RUnlock()
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	jsn.value = filteredArr
	return jsn
}

func (jsn *JSON) HashMapForEach(cb func(*JSON, string)) *JSON {
	isMap, hMap := jsn.GetHashMap()

	if !isMap {
		return jsn
	}

	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	for k, v := range hMap {
		cb(v, k)
	}
	return jsn
}

func (jsn *JSON) HashMapMap(cb func(*JSON, string) interface{})  *JSON {
	isMap, hMap := jsn.GetHashMap()

	if !isMap {
		return jsn
	}

	jsn.rwMutex.RLock()
	res := make(JonsonMap)
	for k, v := range hMap {
		res[k] = jonsonize(cb(v, k))
	}
	jsn.rwMutex.RUnlock()
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	jsn.value = res
	return jsn
}

func (jsn *JSON) HashMapFilter(cb func(*JSON, string) bool)  *JSON {
	isMap, hMap := jsn.GetHashMap()

	if !isMap {
		return jsn
	}

	jsn.rwMutex.RLock()
	res := make(JonsonMap)
	for k, v := range hMap {
		if cb(v, k) {
			res[k] = v
		}
	}
	jsn.rwMutex.RUnlock()
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	jsn.value = res
	return jsn
}