package localcache

import (
	"reflect"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	c := New()
	c.Set("key1", 100)
	got := data["key1"].content

	expect := 100
	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("expected: %v, got: %v", expect, got)
	}
}

func TestGet(t *testing.T) {
	c := New()
	data["key2"] = &CacheData{content: 200, expiredTime: time.Now().Add(30 * time.Duration(time.Second))}
	expect := 200
	got, _ := c.Get("key2")

	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("expected: %v, got: %v", expect, got)
	}
}

func TestGetWithExpiredData(t *testing.T) {
	c := New()
	data["key3"] = &CacheData{content: 200, expiredTime: time.Now().AddDate(0, -1, 0)}
	expectError := ErrDataExpired
	_, err := c.Get("key3")

	if !(expectError == err) {
		t.Fatalf("expected: %v, got: %v", ErrDataExpired, err)
	}
}
