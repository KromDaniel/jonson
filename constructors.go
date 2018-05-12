package jonson

import "encoding/json"

/*
Creates a new Jonson Object with the value
of the interface. the reference to interface is lost
and it is deeply cloned

Possible types:
Primitive
Map[string]interface{}
Slice
struct
 */
func New(value interface{}) *JSON {
	return jonsonize(value)
}

/*
	Creates a new empty Jonson object with null value
 */
func NewEmptyJSON() *JSON {
	return &JSON{
		value:       nil,
		isPrimitive: true,
	}
}

/*
	Creates a new empty Jonson object with empty map
    {}
 */
func NewEmptyJSONMap() *JSON {
	return New(make(map[string]interface{}))
}


/*
	Created a new empty Jonson object with empty array
    []
 */
func NewEmptyJSONArray() *JSON {
	return New(make([]interface{}, 0))
}

/*
	Parses JSON returns err, nil if error
 */
func Parse(data []byte) (err error, jsn *JSON) {
	var m interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return
	}
	jsn = New(m)
	return
}

/*
	Parses JSON returns null json if error
 */
func ParseUnsafe(data []byte) (jsn *JSON) {
	_, jsn = Parse(data)
	if jsn == nil {
		jsn = NewEmptyJSON()
	}
	return jsn
}
