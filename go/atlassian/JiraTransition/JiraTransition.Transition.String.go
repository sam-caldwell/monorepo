package JiraTransition

import "encoding/json"

// String - Return the string form of the transition object.
func (transition *Transition) String() (string, error) {
	b, err := json.Marshal(transition)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
