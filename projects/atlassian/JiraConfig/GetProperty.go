package JiraConfig

// GetProperty - Get the given property value
func (jira *JiraConfig) GetProperty(property *string, reloadFirst bool) (value *Property) {
	if reloadFirst {
		if err := jira.GetFromServer(); err != nil {
			return nil
		}
	}
	for _, item := range jira.Properties {
		if item.Name == *property {
			itemCopy := item
			return &itemCopy
		}
	}
	return nil
}
