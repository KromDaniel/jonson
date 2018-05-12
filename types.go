package Jonson

import (
	"reflect"
	"sync"
)

type JSONObject map[string]*JSON

type JSON struct {
	rwMutex     sync.RWMutex
	isPrimitive bool
	kind        reflect.Kind
	value       interface{}
}
