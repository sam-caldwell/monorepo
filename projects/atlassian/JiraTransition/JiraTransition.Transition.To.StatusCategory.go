package JiraTransition

// StatusCategory - reflect the status category of the given next step in the transition
type StatusCategory struct {
	ColorName string `json:"colorName"`
	ID        int    `json:"id"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	Self      string `json:"self"`
}
