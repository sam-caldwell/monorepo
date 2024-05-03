package Atlassian

import (
    "fmt"
    "regexp"
)

// Set - Sanitize and store the Jira Domain
func (jira *Domain) Set(s *string) (err error) {
    re := regexp.MustCompile(JiraDomainRegEx)
    if re.MatchString(*s) {
        *jira = Domain(*s)
    } else {
        err = fmt.Errorf(MalformedJiraDomain)
    }
    return err
}
