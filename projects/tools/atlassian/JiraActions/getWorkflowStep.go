package JiraActions

import "flag"

// getWorkflowStep - Get the command-line workflow step (--step)
func getWorkflowStep() *string {
	return flag.String("step", "", "Jira workflow step")
}
