package localcache

import (
	"time"
)

type localcache struct {
}

func (lc *localcache) Get(k string) (interface{}, error) {
	cd := Data[k]
	if &cd == nil {
		return nil, ErrValueNotFound
	}
	if time.Now().After(cd.expiredTime) {
		delete(Data, k)
		return nil, ErrDataExpired
	}

	return cd.content, nil
}

func (lc *localcache) Set(k string, d interface{}) error {
	Data[k] = CacheData{content: d, expiredTime: expiredTime()}
	return nil
}

func New() Cache {
	c := localcache{}
	return &c
}

func expiredTime() time.Time {
	t := time.Now().Add(30 * time.Duration(time.Second))
	return t
}
