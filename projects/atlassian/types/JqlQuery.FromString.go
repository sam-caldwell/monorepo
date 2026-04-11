package AtlassianTypes

import "encoding/json"

// FromString - Unmarshall string to JqlQuery internal state
func (jql *JqlQuery) FromString(jqlString *string) (err error) {
	return json.Unmarshal([]byte(*jqlString), jql)
}
