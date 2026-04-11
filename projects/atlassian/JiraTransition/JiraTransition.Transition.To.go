package JiraTransition

// To - show the next step to which the work flow will transition.
type To struct {
	Description    string         `json:"description"`
	IconURL        string         `json:"iconUrl"`
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Self           string         `json:"self"`
	StatusCategory StatusCategory `json:"statusCategory"`
}
