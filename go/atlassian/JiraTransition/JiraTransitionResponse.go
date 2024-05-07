package JiraTransition

type JiraTransitionResponse struct {
	Fields        map[string]Field `json:"fields"`
	HasScreen     bool             `json:"hasScreen"`
	ID            string           `json:"id"`
	IsAvailable   bool             `json:"isAvailable"`
	IsConditional bool             `json:"isConditional"`
	IsGlobal      bool             `json:"isGlobal"`
	IsInitial     bool             `json:"isInitial"`
	Name          string           `json:"name"`
	To            To               `json:"to"`
}

type Field struct {
	Summary `json:"summary"`
	Colour  `json:"colour,omitempty"`
}

type Summary struct {
	AllowedValues   []string `json:"allowedValues"`
	DefaultValue    string   `json:"defaultValue"`
	HasDefaultValue bool     `json:"hasDefaultValue"`
	Key             string   `json:"key"`
	Name            string   `json:"name"`
	Operations      []string `json:"operations"`
	Required        bool     `json:"required"`
	Schema          Schema   `json:"schema"`
}

type Colour struct {
	AllowedValues   []string `json:"allowedValues"`
	DefaultValue    string   `json:"defaultValue"`
	HasDefaultValue bool     `json:"hasDefaultValue"`
	Key             string   `json:"key"`
	Name            string   `json:"name"`
	Operations      []string `json:"operations"`
	Required        bool     `json:"required"`
	Schema          Schema   `json:"schema"`
}

type Schema struct {
	Custom   string `json:"custom"`
	CustomID int    `json:"customId"`
	Items    string `json:"items"`
	Type     string `json:"type"`
}

type To struct {
	Description    string         `json:"description"`
	IconURL        string         `json:"iconUrl"`
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Self           string         `json:"self"`
	StatusCategory StatusCategory `json:"statusCategory"`
}

type StatusCategory struct {
	ColorName string `json:"colorName"`
	ID        int    `json:"id"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	Self      string `json:"self"`
}

type Transitions struct {
	Transitions []Transition `json:"transitions"`
}
