package Jonson

/*
Sets a value to the current JSON,
Makes a deep copy of the interface, removing the original reference
 */
func (jsn *JSON) Set(v interface{}) *JSON {
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()

	temp := jonsonize(v)
	jsn.kind = temp.kind
	jsn.value = temp.value
	jsn.isPrimitive = temp.isPrimitive
	return jsn
}

/*
Set a value to MapObject
if key doesn't exists, it creates it

if current json is not map, it does nothing
 */
func (jsn *JSON) MapSet(key string, value interface{}) *JSON {
	isHashMap, hashMap := jsn.GetMap()
	if !isHashMap {
		return jsn
	}
	jsn.rwMutex.Lock()
	defer jsn.rwMutex.Unlock()
	hashMap[key] = jonsonize(value)
	jsn.value = hashMap

	return jsn
}

/*
Append a value at the end of the slice

if current json slice, it does nothing

multiple values will append in the order of the values
SliceAppend(1,2,3,4) -> [oldSlice..., 1,2,3,4]
 */
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

/*
Append a value at the start of the slice

if current json slice, it does nothing
multiple values will append begin in the order of the values
SliceAppend(1,2,3,4) -> [4,3,2,1, oldSlice...]
 */
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

/*
Sets a value at index to current slice

if value isn't slice, it does nothing
User must make sure the length of the slice contains the index
 */
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
