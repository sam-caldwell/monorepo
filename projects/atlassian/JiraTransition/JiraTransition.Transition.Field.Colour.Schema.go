package JiraTransition

// Schema - Represent Transition.Field.Colour.Schema to provide further details about the transition
type Schema struct {
	Custom   string `json:"custom"`
	CustomID int    `json:"customId"`
	Items    string `json:"items"`
	Type     string `json:"type"`
}
