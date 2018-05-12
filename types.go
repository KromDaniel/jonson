package Jonson

import (
	"reflect"
	"sync"
)

type JSON struct {
	rwMutex     sync.RWMutex
	isPrimitive bool
	kind        reflect.Kind
	value       interface{}
}
