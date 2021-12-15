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
	content     interface{}
	expiredTime time.Time
}

// the interface of cache, provide Get() and Set() two methods
type Cache interface {
	Get(string) (interface{}, error)
	Set(string, interface{}) error
}