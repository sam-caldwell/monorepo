package JiraIssue

// Transition - Transition workflow step for a given ticket
func (jira *Issue) Transition(step *string) (output []byte, err error) {
	//
	//const path = "/rest/api/3/issue/%s/transitions"
	//
	//var transition AtlassianTypes.JiraTransition
	//if err = transition.Init(&jira.client,&jira.IssueKey); err != nil {
	//	return nil, err
	//}
	//if err = transition.GetTransitionIdByName(); err != nil {
	//	return nil, err
	//}
	//
	//return jira.client.Send(
	//	http.MethodPost,
	//	fmt.Sprintf(path, jira.IssueKey),
	//	transition.JsonMarshall())
	return nil, nil
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
