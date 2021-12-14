package localcache

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	c := &localcache{}
	c.Set("key1", 100)
	bs, _ := json.Marshal(Data)
	fmt.Println(string(bs))
	got := Data

	expect := map[string]interface{}{"key1": 100}
	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("expected: %v, got: %v", expect, got)
	}
}

func TestGet(t *testing.T) {
	c := &localcache{}
	Data["key2"] = 200
	expect := 200
	got, _ := c.Get("key2")

	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("expected: %v, got: %v", expect, got)
	}
}
