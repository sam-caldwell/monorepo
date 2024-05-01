package Atlassian

const (
    // DefaultJqlQuery
    //    ToDo: Correct the defaultJqlQuery with default offset and limit
    DefaultJqlQuery = "*"
)

const (
    MalformedJiraDomain      = "JiraDomain malformed"
    JiraDomainRegEx          = `[a-zA-Z][a-zA-Z0-9\-_]+[a-zA-Z0-9]`
    JiraDomainNotInitialized = "JiraDomain not initialized"
    JiraUrlPattern           = "https://%s.atlassian.net/rest/api/2/issue/"
)
