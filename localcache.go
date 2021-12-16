package localcache

import (
	"errors"
)

var (
	// ErrValueNotFound is the error when localcache can not find the vlaue by key
	ErrValueNotFound = errors.New("value not found")
	// ErrDataExpired is the error when the value is found in localcache but it is expired
	ErrDataExpired = errors.New("data expired")
)

// Cache is the interface that provides Get() and Set() two methods
type Cache interface {
	// Get the value/content of cache with a string
	Get(string) (interface{}, error)
	// Set the value/content of cache by a string
	Set(string, interface{}) error
}
