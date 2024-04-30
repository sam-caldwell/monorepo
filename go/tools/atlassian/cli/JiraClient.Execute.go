package cli

// Execute - Execute the Jira Client command-line operation
func (jira *JiraClient) Execute() (err error) {
	switch jira.command {
	case Create:
		err = jira.Create()
	case Read:
		err = jira.Read()
	case Update:
		err = jira.Update()
	case Delete:
		err = jira.Delete()
	case List:
		err = jira.List()
	}
	return err
}
