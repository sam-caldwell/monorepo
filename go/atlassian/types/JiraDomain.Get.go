package AtlassianTypes

import (
    "github.com/sam-caldwell/monorepo/go/atlassian"
    "github.com/sam-caldwell/monorepo/go/misc/words"
)

//Get - Return the Jira domain
func (jira *AtlassianDomain) Get() string {
    if string(*jira) == words.EmptyString {
        panic(atlassian.JiraDomainNotInitialized)
    }
    return string(*jira)
}
