package Atlassian

const (
	// DefaultJqlQuery
	//    ToDo: Correct the defaultJqlQuery with default offset and limit
	DefaultJqlQuery = "*"
)

const (
	atlassianTokenRegex = `^[A-Za-z0-9+/\-=_]{192}$`

	errInvalidAtlassianToken = "invalid apiKey"

	MalformedJiraDomain = "JiraDomain malformed"

	JiraDomainRegEx = `^[a-zA-Z][a-zA-Z0-9\-_]+[a-zA-Z0-9]$`

	JiraDomainNotInitialized = "JiraDomain not initialized"

	JiraUrlPattern = "https://%s.%s/%s"

	JiraApiIssue = "/rest/api/2/issue/"
)
