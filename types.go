package Jonson

import (
	"reflect"
	"sync"
)

type JonsonMap map[string]*JSON

type JSON struct {
	rwMutex     sync.RWMutex
	isPrimitive bool
	kind        reflect.Kind
	value       interface{}
}