/*
Implement a package called `localcache` including two public methods `Get` and `Set` within an interface `Cache`.

`Get` gets the value related to a key which is a string type.
`Set` sets the key/value pair in the local cache implemented by a map.


Please satisfy the following requirement:
* `Set` only keeps a key within 30 seconds.
* `Set` overwrites a value with the same key.
* The value in the cache would be any kind of data in Go.
* Separate the interface and implementation in different files
* Need follow up the guideline mentioned in `Code Review Comments`.
* Need unit test for this package. 3rd-party packages including `stretchr/testify`, `onsi/ginkgo` or `testing` is welcome.

HINT:

```
package localcache

type Cache interface {
	Get( ... )
	Set( ... )
}
```
*/
package localcache

//var Data = make(map[string]interface{})
var Data = make(map[string]interface{}) //{"appId": 2, "fcmServerKey": "keyTestTest", "name": "com.app", "version": []int{1, 2, 3}, "xyz": 3}

type Cache interface {
	Get(k string) (interface{}, error)
	Set(k string) error
}