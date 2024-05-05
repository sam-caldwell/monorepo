package Atlassian

import (
	"fmt"
	"regexp"
)

// Set - Sanitize and store the Jira Domain
func (jira *Domain) Set(s *string) (err error) {
	re := regexp.MustCompile(JiraDomainRegEx)
	if !re.MatchString(*s) {
		return fmt.Errorf(MalformedJiraDomain)
	}
	*jira = Domain(*s)
	return nil
}
