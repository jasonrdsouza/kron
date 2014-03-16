package main

import (
  "sync"
)


type Environment struct {
  sync.RWMutex
  vars map[string]string
}

func NewEnvironment() *Environment {
  return &Environment{ *new(sync.RWMutex), make(map[string]string) }
}

// Gets the specified environment variable. If the key is not found, the empty
// string (zero value) is returned. Just as with regular go maps, the second
// return value will be true or false depending the key's existance.
func (e *Environment) Get(key string) (string, bool) {
  e.RLock()
  defer e.RUnlock()
  val, ok := e.vars[key]
  return val, ok
}

// Sets the specified environment variable, overwriting if it already exists.
func (e *Environment) Set(key, value string) {
  e.Lock()
  defer e.Unlock()
  e.vars[key] = value
}

// Replaces the current environment with a new one
func (e *Environment) Replace(vars map[string]string) {
  e.Lock()
  defer e.Unlock()
  e.vars = vars
}
