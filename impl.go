package localcache

import (
	"errors"
)

type localcache struct{}

func (lc *localcache) Get(k string) (interface{}, error) {
	t := Data[k]
	if t != nil {
		return t, nil
	}
	return nil, errors.New("value not found")
}

func (lc *localcache) Set(k string, d interface{}) error {
	Data[k] = d
	return nil
}
