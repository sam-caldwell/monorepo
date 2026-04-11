package JiraTransition

//
//// GetTransitionIdByName - Given a workflow step name and assoc. issue key, return the step name
//func (jira *JiraTransition) GetTransitionIdByName(IssueKey, stepName *string) error {
//
//	const path = "/rest/api/3/issue/%s/transitions"
//
//	response, err := jira.client.Send(
//		http.MethodPost,
//		fmt.Sprintf(path, *IssueKey),
//		jira.Marshall())
//	if err != nil {
//		return err
//	}
//	if err = jira.Unmarshall(response); err != nil {
//		return err
//	}
//}
