package keyval

import (
	"testing"
)

func TestKeyVal(t *testing.T) {
	var kvs KeyVals

	kvs.Add("key1", "value1")
	kvs.Add("key2", "value2")
	kvs.Add("key3", "value3")

	for _, kv := range kvs {
		key := kv.Key.(string)
		val := kv.Value.(string)

		t.Logf("key: %s, val: %s", key, val)
	}
}
