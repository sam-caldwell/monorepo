package Atlassian

const (
	// DefaultJqlQuery
	//    ToDo: Correct the defaultJqlQuery with default offset and limit
	DefaultJqlQuery = "*"
)

const (
	atlassianTokenRegex = `^[A-Za-z0-9+/\-=_]{192}$`

	ContentType = "application/json;charset=UTF-8"

	errInvalidAtlassianToken = "invalid apiKey"

	MalformedJiraDomain = "JiraDomain malformed"

	JiraDomainRegEx = `^[a-zA-Z][a-zA-Z0-9\-_]+[a-zA-Z0-9]$`

	JiraDomainNotInitialized = "JiraDomain not initialized"

	JiraUrlPattern = "https://%s.%s/%s"

	JiraApiIssue = "/rest/api/2/issue/"

	// Search(List) query defaults
	MaxResults     uint = 1000
	DefaultStartAt uint = 0
	//Note: this should be a comma-delimited list
	DefaultExpand string = ""
	//Note: this should be a comma-delimited list
	DefaultFields string = ""
)
