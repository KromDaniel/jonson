package main

import (
	"reflect"
	"sync"
)

type JonsonMap map[string]*Jonson

type Jonson struct {
	rwMutex     sync.RWMutex
	isPrimitive bool
	kind        reflect.Kind
	value       interface{}
}