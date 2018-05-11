package Jonson

func (jsn *Jonson) SliceForEach(cb func(*Jonson, int)) *Jonson {
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

func (jsn *Jonson) SliceMap(cb func(*Jonson, int) interface{})*Jonson {
	isSlice, slice := jsn.GetSlice()

	if !isSlice {
		return jsn
	}
	jsn.rwMutex.RLock()
	mappedArr := make([]*Jonson, len(slice))
	for i, v := range slice {
		mappedArr[i] = jonsonize(cb(v, i))
	}
	jsn.rwMutex.RUnlock()
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	jsn.value = mappedArr
	return jsn
}

func (jsn *Jonson) SliceFilter(cb func(*Jonson, int) bool)*Jonson {
	isSlice, slice := jsn.GetSlice()

	if !isSlice {
		return jsn
	}
	jsn.rwMutex.RLock()
	filteredArr := make([]*Jonson, 0)
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

func (jsn *Jonson) HashMapForEach(cb func(*Jonson, string)) *Jonson {
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

func (jsn *Jonson) HashMapMap(cb func(*Jonson, string) interface{})  *Jonson {
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

func (jsn *Jonson) HashMapFilter(cb func(*Jonson, string) bool)  *Jonson {
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