package JiraIssue

type Issue struct {
    Fields struct {
        Project struct {
            Key string `json:"key"`
        } `json:"project"`
        Summary     string `json:"summary"`
        Description struct {
            Type    string `json:"type"`
            Version int    `json:"version"`
            Content []struct {
                Type    string `json:"type"`
                Content []struct {
                    Type string `json:"type"`
                    Text string `json:"text"`
                } `json:"content"`
            } `json:"content"`
        } `json:"description"`
        IssueType struct {
            Name string `json:"name"`
        } `json:"issuetype"`
    } `json:"fields"`
}
