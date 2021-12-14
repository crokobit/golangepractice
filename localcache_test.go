package localcache

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	c := &localcache{}
	c.Set("key1", 100)
	bs, _ := json.Marshal(Data)
	fmt.Println(string(bs))
	got := Data["key1"].content

	expect := 100
	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("expected: %v, got: %v", expect, got)
	}
}

func TestGet(t *testing.T) {
	c := &localcache{}
	Data["key2"] = CacheData{content: 200, expiredTime: time.Now().Add(30 * time.Duration(time.Second))}
	expect := 200
	got, _ := c.Get("key2")

	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("expected: %v, got: %v", expect, got)
	}
}

func TestGetWithExpiredData(t *testing.T) {
	c := &localcache{}
	Data["key3"] = CacheData{content: 200, expiredTime: time.Now().AddDate(0, -1, 0)}
	expectError := ErrDataExpired
	_, err := c.Get("key3")

	if !(expectError == err) {
		t.Fatalf("expected: %v, got: %v", ErrDataExpired, err)
	}
}
