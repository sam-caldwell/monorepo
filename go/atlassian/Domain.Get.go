package Atlassian

import (
    "github.com/sam-caldwell/monorepo/go/misc/words"
)

//Get - Return the Jira domain
func (jira *Domain) Get() string {
    if string(*jira) == words.EmptyString {
        panic(JiraDomainNotInitialized)
    }
    return string(*jira)
}
