package lib

func (jsn *Jonson) SliceForEach(cb func(*Jonson, int)) {
	isSlice, slice := jsn.GetSlice()

	if !isSlice {
		return
	}

	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	for i, v := range slice {
		cb(v, i)
	}
}

func (jsn *Jonson) SliceMap(cb func(*Jonson, int) interface{}) {
	isSlice, slice := jsn.GetSlice()

	if !isSlice {
		return
	}
	jsn.rwMutex.RLock()
	mappedArr := make([]*Jonson, len(slice))
	for i, v := range slice {
		mappedArr[i] = gosonize(cb(v, i))
	}
	jsn.rwMutex.RUnlock()
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	jsn.value = mappedArr
}

func (jsn *Jonson) SliceFilter(cb func(*Jonson, int) bool) {
	isSlice, slice := jsn.GetSlice()

	if !isSlice {
		return
	}
	jsn.rwMutex.RLock()
	filteredArr := make([]*Jonson, 0)
	for i, v := range slice {
		if keep := cb(v, i); keep {
			filteredArr = append(filteredArr, v)
		}
	}
	jsn.rwMutex.RUnlock()
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	jsn.value = filteredArr
}

func (jsn *Jonson) HashMapForEach(cb func(*Jonson, string)) {
	isMap, hMap := jsn.GetHashMap()

	if !isMap {
		return
	}

	jsn.rwMutex.RLock()
	defer jsn.rwMutex.RUnlock()
	for k, v := range hMap {
		cb(v, k)
	}
}

func (jsn *Jonson) HashMapMap(cb func(*Jonson, string) interface{}) {
	isMap, hMap := jsn.GetHashMap()

	if !isMap {
		return
	}

	jsn.rwMutex.RLock()
	res := make(JonsonMap)
	for k, v := range hMap {
		res[k] = gosonize(cb(v, k))
	}
	jsn.rwMutex.RUnlock()
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	jsn.value = res
}

func (jsn *Jonson) HashMapFilter(cb func(*Jonson, string) bool) {
	isMap, hMap := jsn.GetHashMap()

	if !isMap {
		return
	}

	jsn.rwMutex.RLock()
	res := make(JonsonMap)
	for k, v := range hMap {
		if keep := cb(v, k); keep {
			res[k] = v
		}
	}
	jsn.rwMutex.RUnlock()
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	jsn.value = res
}