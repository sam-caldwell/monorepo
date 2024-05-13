package JQL

import (
	"encoding/json"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"net/http"
	"strings"
)

// Parse - Parse an input string (passed by reference) and store as internal state.
func (jql *JQL) Parse(query *string) (err error) {

	const path = "/rest/api/3/jql/parse?validation=strict"

	var response []byte

	if jql.client.GetDebug() {
		ansi.Blue().Println("Parsing JQL").LF().Reset()
	}

	if jql.client == nil {
		return fmt.Errorf("client is nil")
	}

	payload := []byte(fmt.Sprintf("{\"queries\":[\"%s\"]}", *query))

	response, err = jql.client.Send(
		http.MethodPost,
		path,
		payload)

	if err != nil {
		return err
	}

	if jql.client.GetDebug() {
		ansi.Blue().
			Line("-", 40).
			Printf("response (raw): '%s'", response).LF().
			Line("-", 40).
			Reset()
	}

	if len(response) == 0 {
		return fmt.Errorf("empty response from server")
	}
	if err = json.Unmarshal(response, &jql); err != nil {
		return fmt.Errorf("unmarshal failed: %v", err)
	}
	var e []string
	e = append(e, "{\"errors\":[")
	for i, _ := range jql.Queries {
		if jql.Queries[i].Errors != nil && len(jql.Queries[i].Errors) > 0 {
			e = append(e, fmt.Sprintf(" %d: '%s'", i, strings.Join(jql.Queries[i].Errors, words.Comma)))
		}
	}
	if len(e) > 1 {
		e = append(e, "]}")
		err = fmt.Errorf("%s", strings.Join(e, "\n"))
	}
	return err
}
