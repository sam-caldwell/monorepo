package JiraIssue

import (
	"fmt"
	"net/http"
)

// Transition - Transition workflow step for a given ticket
func (jira *Issue) Transition(step *string) (output []byte, err error) {

	const path = "/rest/api/3/issue/%s/transitions"

	return jira.client.Send(
		http.MethodPost,
		fmt.Sprintf(path, jira.IssueKey),
		jira.Marshall())

}

/*
   {
     "fields": {
       "assignee": {
         "name": "bob"
       },
       "resolution": {
         "name": "Fixed"
       }
     },
     "historyMetadata": {
       "activityDescription": "Complete order processing",
       "actor": {
         "avatarUrl": "http://mysystem/avatar/tony.jpg",
         "displayName": "Tony",
         "id": "tony",
         "type": "mysystem-user",
         "url": "http://mysystem/users/tony"
       },
       "cause": {
         "id": "myevent",
         "type": "mysystem-event"
       },
       "description": "From the order testing process",
       "extraData": {
         "Iteration": "10a",
         "Step": "4"
       },
       "generator": {
         "id": "mysystem-1",
         "type": "mysystem-application"
       },
       "type": "myplugin:type"
     },
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
