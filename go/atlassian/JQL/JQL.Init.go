package JQL

import (
	"fmt"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
)

// Init - initialize the jql query object.
func (jql *JQL) Init(client *Atlassian.Client) (err error) {
	if client == nil {
		err = fmt.Errorf("client was not initialized (nil)")
	} else {
		jql.client = client
	}
	return err
}
