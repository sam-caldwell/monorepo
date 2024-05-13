package JiraConfig

// GetFromServer - Call the Atlassian Jira instance (specified by client) and load the settings
func (config *JiraConfig) GetFromServer() (err error) {
	if err = config.GetGlobalSettings(); err != nil {
		return err
	}
	if err = config.GetAdvancedSettings(); err != nil {
		return err
	}
	return err
}

// GetGlobalSettings - Load the global settings from the Jira server
func (config *JiraConfig) GetGlobalSettings() error {
	const path = "/rest/api/3/configuration"
	return nil
}

// GetAdvancedSettings - Load the advanced settings from the Jira server
func (config *JiraConfig) GetAdvancedSettings() error {
	const path = ""
	return nil
}

// GetProperty - Load a specific property from the JIRA server
func (config *JiraConfig) GetProperty(property *string) (string, error) {
	const path = ""
	return "", nil
}

// SetProperty - Set a specific property-value on the JIRA server
func (config *JiraConfig) SetProperty(property *string, value *string) error {
	const path = ""
	return nil
}
