package AtlassianTypes

import (
    "fmt"
    "github.com/sam-caldwell/monorepo/go/atlassian"
    "regexp"
)

// Set - Sanitize and store the Jira Domain
func (jira *AtlassianDomain) Set(s *string) (err error) {
    re := regexp.MustCompile(atlassian.JiraDomainRegEx)
    if re.MatchString(*s) {
        *jira = AtlassianDomain(*s)
    } else {
        err = fmt.Errorf(atlassian.MalformedJiraDomain)
    }
    return err
}
