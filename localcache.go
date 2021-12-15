package localcache

import (
	"errors"
	"time"
)

var data = make(map[string]CacheData)

var (
	ErrValueNotFound = errors.New("value not found")
	ErrDataExpired   = errors.New("data expired")
)

type CacheData struct {
	content     interface{}
	expiredTime time.Time
}

type Cache interface {
	Get(k string) (interface{}, error)
	Set(k string, d interface{}) error
}