package JiraTransition

// TransitionAction - structure representing request payload for JiraTransition action
type TransitionAction struct {
	Fields struct {
		Assignee struct {
			Name string `json:"name,omitempty"`
		} `json:"assignee,omitempty"`
		Resolution struct {
			Name string `json:"name,omitempty"`
		} `json:"resolution,omitempty"`
	} `json:"fields,omitempty"`

	HistoryMetadata struct {
		ActivityDescription string `json:"activityDescription,omitempty"`
		Actor               struct {
			AvatarURL   string `json:"avatarUrl,omitempty"`
			DisplayName string `json:"displayName,omitempty"`
			ID          string `json:"id,omitempty"`
			Type        string `json:"type,omitempty"`
			URL         string `json:"url,omitempty"`
		} `json:"actor,omitempty"`
		Cause struct {
			ID   string `json:"id,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"cause,omitempty"`
		Description string `json:"description,omitempty"`
		ExtraData   struct {
			Iteration string `json:"Iteration,omitempty"`
			Step      string `json:"Step,omitempty"`
		} `json:"extraData,omitempty"`
		Generator struct {
			ID   string `json:"id,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"generator,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"historyMetadata,omitempty"`

	Transition struct {
		ID string `json:"id,omitempty"`
	} `json:"transition,omitempty"`

	Update struct {
		Comment []struct {
			Add struct {
				Body struct {
					Content []struct {
						Content []struct {
							Text string `json:"text,omitempty"`
							Type string `json:"type,omitempty"`
						} `json:"content,omitempty"`
						Type string `json:"type,omitempty"`
					} `json:"content,omitempty"`
					Type    string `json:"type,omitempty"`
					Version int    `json:"version,omitempty"`
				} `json:"body,omitempty"`
			} `json:"add,omitempty"`
		} `json:"comment,omitempty"`
	} `json:"update,omitempty"`
}
