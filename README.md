# KeyVal

Very simple Key-Value structure, as [golang](https://golang.org)'s map
does not guarantee sorting order but we do want to have a standardized
type that can be passed around for this.

## Usage / Example

Short usage example.

```
package main

import (
	"fmt"
	"trident.li/keyval"
)

func main() {
	var kvs keyval.KeyVals

	kvs.Add("key1", "value1")
	kvs.Add("key2", "value2")
	kvs.Add("key3", "value3")

	for _, kv := range kvs {
		key := kv.Key.(string)
		val := kv.Value.(string)

		fmt.Println("key: %s, val: %s", key, val)
	}
}
```
