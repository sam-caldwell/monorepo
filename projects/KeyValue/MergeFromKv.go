package keyvalue

func (kv *KeyValue) MergeFromKv(source KeyValue) {
	for key, value := range source.data {
		kv.data[key] = value
	}
}
