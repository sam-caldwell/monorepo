package JiraConfig

import (
	"fmt"
	"strings"
)

// Print - Return a pretty-print string from the JIRA configuration
func (jira *JiraConfig) String() (output string) {
	return fmt.Sprintf("settings:{\n%s,\n%s}\n",
		globalSettingsToString(&jira.GlobalSettings),
		advancedSettingsToString(&jira.AdvancedSettings))
}

func globalSettingsToString(settings *map[string]string) string {
	var result []string
	for key, value := range *settings {
		result = append(result, fmt.Sprintf("  %s:%s,", key, value))
	}
	return fmt.Sprintf("global:{\n%s}\n", strings.Join(result, "\n"))
}

func advancedSettingsToString(settings *map[string]string) string {
	var result []string
	for key, value := range *settings {
		result = append(result, fmt.Sprintf("  %s:%s,", key, value))
	}
	return fmt.Sprintf("advanced:{\n%s}\n", strings.Join(result, "\n"))

}
