package Jonson

import "encoding/json"

func New(value interface{}) *Jonson {
	return jonsonize(value)
}

func NewEmptyJSON() *Jonson {
	return &Jonson{
		value:       nil,
		isPrimitive: true,
	}
}

func NewEmptyHashJSON() *Jonson {
	return New(make(map[string]interface{}))
}

func NewEmptySlice() *Jonson {
	return New(make([]interface{}, 0))
}

func Parse(data []byte) (err error, jsn *Jonson) {
	var m interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return
	}
	jsn = New(m)
	return
}

func ParseUnsafe(data []byte) *Jonson {
	_, goson := Parse(data)
	if goson == nil {
		return NewEmptyJSON()
	}
	return goson
}
