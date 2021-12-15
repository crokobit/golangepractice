package localcache

import (
	"errors"
	"time"
)

var (
	// the error when can not find vlaue by key
	ErrValueNotFound = errors.New("value not found")
	// the error when value found but it is expired
	ErrDataExpired = errors.New("data expired")
)

// the format using for cache, including content and expiredTime
type CacheData struct {
	// the content of cache
	content interface{}
	// the expiredTime of cache
	expiredTime time.Time
}

// the interface of cache, provide Get() and Set() two methods
type Cache interface {
	// Get the value/content of cache with a string
	Get(string) (interface{}, error)
	// Set the value/content of cache by a string
	Set(string, interface{}) error
}