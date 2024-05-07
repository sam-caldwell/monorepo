package JiraIssue

import (
	"fmt"
	AtlassianTypes "github.com/sam-caldwell/monorepo/go/atlassian/JiraTransition"
	"net/http"
)

// Transition - Transition workflow step for a given ticket
func (jira *Issue) Transition(step *string) (output []byte, err error) {

	const path = "/rest/api/3/issue/%s/transitions"

	var transition AtlassianTypes.JiraTransition
	if err = transition.Init(&jira.client); err != nil {
		return nil, err
	}
	if err = transition.GetTransitionIdByName(&jira.IssueKey, step); err != nil {
		return nil, err
	}

	return jira.client.Send(
		http.MethodPost,
		fmt.Sprintf(path, jira.IssueKey),
		transition.JsonMarshall())

}

/*
   {
     "transition": {
       "id": "5"
     },
     "update": {
       "comment": [
         {
           "add": {
             "body": {
               "content": [
                 {
                   "content": [
                     {
                       "text": "Bug has been fixed",
                       "type": "text"
                     }
                   ],
                   "type": "paragraph"
                 }
               ],
               "type": "doc",
               "version": 1
             }
           }
         }
       ]
     }
   }
*/
