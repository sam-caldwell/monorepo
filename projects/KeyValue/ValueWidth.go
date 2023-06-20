package keyvalue

/*
 * KeyValue.ValueWidth
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * 	Given a set of key-value data, return the maximum string length of all values
 */

// ValueWidth - Return the maximum width of all values in the current KeyValue struct
func (kv *KeyValue) ValueWidth() (width int) {
	if kv.data != nil {
		for _, value := range kv.data {
			switch v := value.(type) {
			case string:
				if len(v) > width {
					width = len(v)
				}
			}
		}
	}
	return width
}
