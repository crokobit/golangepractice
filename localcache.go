package localcache

import (
	"errors"
	"time"
)

var Data = make(map[string]CacheData)

var (
	// ErrValueNotFound indicates the key is missing
	ErrValueNotFound = errors.New("value not found")
	// ErrDataExpired indicates the key is found but data is expired
	ErrDataExpired = errors.New("data expired")
)

type Cache interface {
	Get(k string) (interface{}, error)
	Set(k string) error
}

// the struct be saved in the Data var
type CacheData struct {
	content     interface{}
	expiredTime time.Time
}