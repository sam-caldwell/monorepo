package JiraIssue

import "encoding/json"

// Marshall - marshall the jira issue and return its json []byte form
func (jira Issue) Marshall() []byte {

    result, err := json.Marshal(jira)

    if err != nil {
        panic("internal error")
    }

    return result

}
