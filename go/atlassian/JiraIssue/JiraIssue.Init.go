package JiraIssue

// Init - Initialize the state of the Jira Issue object
func (jira *Issue) Init(debug bool, noop bool, username, apiKey, domain, descriptor, issueOrKey, jqlString *string) error {

	var err error

	jira.client.SetDebug(debug)

	jira.client.SetNoop(noop)

	if err = jira.client.SetDomain(domain); err != nil {
		return err
	}

	if err = jira.client.SetUsername(username); err != nil {
		return err
	}

	if err = jira.client.SetApiKey(apiKey); err != nil {
		return err
	}

	if descriptor != nil {
		if err = jira.Load(descriptor); err != nil {
			return err
		}
	}

	if (issueOrKey != nil) && (*issueOrKey != "") {
		jira.IssueKey = *issueOrKey
	}

	if (jqlString != nil) && (*jqlString != "") {
		if err = jira.jql.FromString(jqlString); err != nil {
			return err
		}
	}

	return nil

}
