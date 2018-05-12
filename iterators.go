package jonson

// iterates on slice
func (jsn *JSON) SliceForEach(cb func(jsn *JSON, index int)) *JSON {
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

// iterates on slice with map callback, transforming the slice to new slice
func (jsn *JSON) SliceMap(cb func(jsn *JSON, index int) *JSON) *JSON {
	isSlice, slice := jsn.GetSlice()

	if !isSlice {
		return jsn
	}
	jsn.rwMutex.RLock()
	mappedArr := make([]*JSON, len(slice))
	for i, v := range slice {
		mappedArr[i] = cb(v, i)
	}
	jsn.rwMutex.RUnlock()
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	jsn.value = mappedArr
	return jsn
}

// iterates on slice with filter callback, removes values that callback returned false
func (jsn *JSON) SliceFilter(cb func(jsn *JSON, index int) (shouldKeep bool)) *JSON {
	isSlice, slice := jsn.GetSlice()

	if !isSlice {
		return jsn
	}
	jsn.rwMutex.RLock()
	filteredArr := make([]*JSON, 0)
	for i, v := range slice {
		if cb(v, i) {
			filteredArr = append(filteredArr, slice[i])
		}
	}
	jsn.rwMutex.RUnlock()
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	jsn.value = filteredArr
	return jsn
}

// iterates on object
func (jsn *JSON) ObjectForEach(cb func(jsn *JSON, key string)) *JSON {
	isMap, hMap := jsn.GetMap()

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

// iterates on object, replacing each value with new returned value
func (jsn *JSON) ObjectMap(cb func(jsn *JSON, key string) *JSON) *JSON {
	isMap, hMap := jsn.GetMap()

	if !isMap {
		return jsn
	}

	jsn.rwMutex.RLock()
	res := make(map[string]*JSON)
	for k, v := range hMap {
		res[k] = cb(v, k)
	}
	jsn.rwMutex.RUnlock()
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	jsn.value = res
	return jsn
}

// iterates on object, removing each value that cb returns false
func (jsn *JSON) ObjectFilter(cb func(jsn *JSON, key string) (shouldKeep bool)) *JSON {
	isMap, hMap := jsn.GetMap()

	if !isMap {
		return jsn
	}

	jsn.rwMutex.RLock()
	res := make(map[string]*JSON)
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
