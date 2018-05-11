package Jonson

import "encoding/json"

func New(value interface{}) *JSON {
	return jonsonize(value)
}

func NewEmptyJSON() *JSON {
	return &JSON{
		value:       nil,
		isPrimitive: true,
	}
}

func NewEmptyHashJSON() *JSON {
	return New(make(map[string]interface{}))
}

func NewEmptySlice() *JSON {
	return New(make([]interface{}, 0))
}

func Parse(data []byte) (err error, jsn *JSON) {
	var m interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return
	}
	jsn = New(m)
	return
}

func ParseUnsafe(data []byte) *JSON {
	_, goson := Parse(data)
	if goson == nil {
		return NewEmptyJSON()
	}
	return goson
}
