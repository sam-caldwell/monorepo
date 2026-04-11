package JiraTransition

// Transition - Represent the top level of a single transition
type Transition struct {
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
