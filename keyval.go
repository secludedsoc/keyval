package keyval

/* Very simple Key-Value structure, as map does not guarantee sorting order */

type KeyVal struct {
	Key   interface{}
	Value interface{}
}

type KeyVals []KeyVal

/* Add an item to the KeyVal slice */
func (kvs *KeyVals) Add(key interface{}, val interface{}) {
	kv := KeyVal{key, val}

	*kvs = append(*kvs, kv)
}
