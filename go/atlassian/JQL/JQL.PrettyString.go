package JQL

import "encoding/json"

// PrettyString - Marshall the internal state to an indented (pretty) output
func (jql *JQL) PrettyString() (output string, err error) {

	b, err := json.MarshalIndent(jql, "", "  ")

	return string(b), err

}
